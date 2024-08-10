// This  package is used to handle depreceated options in the config file
// These are the options that may break during the application

package migration

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/creasty/defaults"
	"github.com/sirupsen/logrus"
	"golang.org/x/exp/maps"
)

func Migrate(logger *logrus.Entry, optPrefix string, deprecated any, newOptions map[string]Migrator) bool {
	depreceatedVal := reflect.ValueOf(deprecated)
	depreceatedTyp := depreceatedVal.Type()

	usesDepredOpts := false
	for i := 0; i < depreceatedTyp.NumField(); i++ {
		field := depreceatedTyp.Field(i)
		fieldTag := field.Tag.Get("yaml")
		oldName := fullname(optPrefix, fieldTag)

		migrator, ok := newOptions[fieldTag]

		if !ok {
			panic(fmt.Errorf("deprecated option %s has no matching %T", oldName, migrator))
		}
		delete(newOptions, fieldTag)

		migrator.dest.prefix = optPrefix

		val := depreceatedVal.Field(i)
		if val.Type().Kind() != reflect.Pointer {
			panic(fmt.Errorf("Deprecated option %s must be a pointer", oldName))
		}

		if field.Tag.Get("default") != "" {
			panic(fmt.Errorf("Deprecated option %s must not have a default tag", oldName))
		}

		if val.IsNil() {
			continue
		}
		usesDepredOpts = true

		val = val.Elem()

		if !migrator.dest.IsDefault() {
			logger.WithFields(logrus.Fields{
				migrator.dest.Name(): migrator.dest.Value.Interface(),
				oldName:              val.Interface(),
			}).Errorf("config options %q nd %q are both set, ignoring the deprecated one ", migrator.dest, oldName)
			continue
		}

		logger.Warnf("Config option %q is deprecated, please use %q instead", oldName, migrator.dest)
		migrator.apply(oldName, val)
	}

	if len(newOptions) != 0 {
		panic(fmt.Errorf("%q has unused migrations: %v", optPrefix, maps.Keys(newOptions)))
	}
	return usesDepredOpts

}

type applyFunc func(oldName string, oldValue reflect.Value)

type Migrator struct {
	dest  *Dest
	apply applyFunc
}

func newMigrator(dest *Dest, apply applyFunc) Migrator {
	return Migrator{dest, apply}
}

func Move(dest *Dest) Migrator {
	return newMigrator(dest, func(oldName string, oldValue reflect.Value) {
		dest.Value.Set(oldValue)
	})
}

func Apply[T any](dest *Dest, apply func(oldValue T)) Migrator {
	return newMigrator(dest, func(oldName string, oldValue reflect.Value) {
		valItf := oldValue.Interface()
		valTyped, ok := valItf.(T)
		if !ok {
			panic(fmt.Errorf("expected type %T, got %T", valTyped, valItf))
		}
		apply(valTyped)
	})
}

type Dest struct {
	prefix string
	name   string

	Value   reflect.Value
	Default any
}

func To[T any](newName string, newContainerStruct *T) *Dest {
	stVal := reflect.ValueOf(newContainerStruct).Elem()

	if stVal.Type().Kind() != reflect.Pointer {
		panic(fmt.Errorf("newContainerStruct for %s is a double pointer: %T", newName, newContainerStruct))
	}

	fieldIdx, newVal := func() (int, reflect.Value) {
		parts := strings.Split(newName, ".")
		tag := parts[len(parts)-1]

		for i := 0; i < stVal.NumField(); i++ {
			field := stVal.Type().Field(i)
			if field.Tag.Get("yaml") == tag {
				return i, stVal.Field(i)
			}
		}
		panic(fmt.Errorf("migrated option  %q not found in %T", newName, newContainerStruct))
	}()

	newDefaultVal := func() reflect.Value {
		defaultVals := new(T)
		defaults.MustSet(defaultVals)

		return reflect.ValueOf(defaultVals).Elem().Field(fieldIdx)
	}()

	return &Dest{
		prefix:  "",
		name:    newName,
		Value:   newVal,
		Default: newDefaultVal.Interface(),
	}

}

func (d *Dest) Name() string {
	return fullname(d.prefix, d.name)
}

func (d *Dest) IsDefault() bool {
	return reflect.DeepEqual(d.Value.Interface(), d.Default)
}

func (d *Dest) String() string {
	return d.Name()
}
func fullname(prefix, name string) string {
	if len(prefix) == 0 {
		return name
	}

	return fmt.Sprintf("%s.%s", prefix, name)
}
