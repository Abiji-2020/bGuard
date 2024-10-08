package config

import (
	"fmt"
	"strings"
)

const (
	// BytesSourceTypeText is a BytesSourceType of type Text.
	// Inline YAML block.
	BytesSourceTypeText BytesSourceType = iota + 1
	// BytesSourceTypeHttp is a BytesSourceType of type Http.
	// HTTP(S).
	BytesSourceTypeHttp
	// BytesSourceTypeFile is a BytesSourceType of type File.
	// Local file.
	BytesSourceTypeFile
)

var ErrInvalidBytesSourceType = fmt.Errorf("not a valid BytesSourceType, try [%s]", strings.Join(_BytesSourceTypeNames, ", "))

const _BytesSourceTypeName = "texthttpfile"

var _BytesSourceTypeNames = []string{
	_BytesSourceTypeName[0:4],
	_BytesSourceTypeName[4:8],
	_BytesSourceTypeName[8:12],
}

// BytesSourceTypeNames returns a list of possible string values of BytesSourceType.
func BytesSourceTypeNames() []string {
	tmp := make([]string, len(_BytesSourceTypeNames))
	copy(tmp, _BytesSourceTypeNames)
	return tmp
}

// BytesSourceTypeValues returns a list of the values for BytesSourceType
func BytesSourceTypeValues() []BytesSourceType {
	return []BytesSourceType{
		BytesSourceTypeText,
		BytesSourceTypeHttp,
		BytesSourceTypeFile,
	}
}

var _BytesSourceTypeMap = map[BytesSourceType]string{
	BytesSourceTypeText: _BytesSourceTypeName[0:4],
	BytesSourceTypeHttp: _BytesSourceTypeName[4:8],
	BytesSourceTypeFile: _BytesSourceTypeName[8:12],
}

// String implements the Stringer interface.
func (x BytesSourceType) String() string {
	if str, ok := _BytesSourceTypeMap[x]; ok {
		return str
	}
	return fmt.Sprintf("BytesSourceType(%d)", x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x BytesSourceType) IsValid() bool {
	_, ok := _BytesSourceTypeMap[x]
	return ok
}

var _BytesSourceTypeValue = map[string]BytesSourceType{
	_BytesSourceTypeName[0:4]:  BytesSourceTypeText,
	_BytesSourceTypeName[4:8]:  BytesSourceTypeHttp,
	_BytesSourceTypeName[8:12]: BytesSourceTypeFile,
}

// ParseBytesSourceType attempts to convert a string to a BytesSourceType.
func ParseBytesSourceType(name string) (BytesSourceType, error) {
	if x, ok := _BytesSourceTypeValue[name]; ok {
		return x, nil
	}
	return BytesSourceType(0), fmt.Errorf("%s is %w", name, ErrInvalidBytesSourceType)
}

// MarshalText implements the text marshaller method.
func (x BytesSourceType) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *BytesSourceType) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseBytesSourceType(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}
