package parsers

import (
	"context"
	"errors"
	"fmt"
	"io"
)

// SeriesParser parses a series of `T`.
type SeriesParser[T any] interface {
	Next(context.Context) (T, error)

	Position() string
}

func ForEach[T any](ctx context.Context, parser SeriesParser[T], callback func(T) error) (rerr error) {
	defer func() {
		rerr = ErrWithPosition(parser, rerr)
	}()

	for {
		if err := ctx.Err(); err != nil {
			return err
		}

		res, err := parser.Next(ctx)
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}

			return err
		}

		err = callback(res)
		if err != nil {
			return err
		}
	}
}

// ErrWithPosition adds the `parser`'s position to the given `err`.
func ErrWithPosition[T any](parser SeriesParser[T], err error) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf("%s: %w", parser.Position(), err)
}

// IsNonResumableErr is a helper to check if an error returned by a parser is resumable.
func IsNonResumableErr(err error) bool {
	var nonResumableError *NonResumableError

	return errors.As(err, &nonResumableError)
}

// NonResumableError represents an error from which a parser cannot recover.
type NonResumableError struct {
	inner error
}

// NewNonResumableError creates and returns a new `NonResumableError`.
func NewNonResumableError(inner error) error {
	return &NonResumableError{inner}
}

func (e *NonResumableError) Error() string {
	return fmt.Sprintf("non resumable parse error: %s", e.inner.Error())
}

func (e *NonResumableError) Unwrap() error {
	return e.inner
}
