package convm

import (
	"strconv"
	"fmt"
)

type Nil struct{}

func Conv(in string, ret any) error {
	var err error
	switch retc := ret.(type) {
		case *string:
			*retc, err = Str(in)
			return err
		case *int64:
			*retc, err = Int64(in)
			return err
		case *int:
			*retc, err = Int(in)
			return err
		case *float64:
			*retc, err = Float64(in)
			return err
		case *bool:
			*retc, err = Bool(in)
			return err
		case *complex128:
			*retc, err = Complex(in)
			return err
		case Nil:
			return nil
		default:
			return fmt.Errorf("Conv: return value %v %#v not a compatible pointer", retc, retc)
	}
}

func Multi(in []string, out ...any) (int, error) {
	if len(in) != len(out) {
		return 0, fmt.Errorf("ConvMulti: len(in) %v != len(out) %v\n", len(in), len(out))
	}
	for i:=0; i<len(in); i++ {
		err := Conv(in[i], out[i])
		if err != nil {
			return i, fmt.Errorf("ConvMulti: %w", err)
		}
	}
	return len(in), nil
}

func Str(in string) (string, error) {
	return in, nil
}

func Int64(in string) (int64, error) {
	out, err := strconv.ParseInt(in, 0, 64)
	if err != nil {
		return 0, fmt.Errorf("ConvInt64: failed to parse input %v", in)
	}
	return out, nil
}

func Int(in string) (int, error) {
	i, err := strconv.ParseInt(in, 0, strconv.IntSize)
	if err != nil {
		return 0, fmt.Errorf("ConvInt: failed to parse input %v", in)
	}
	out := int(i)
	return out, nil
}

func Float64(in string) (float64, error) {
	out, err := strconv.ParseFloat(in, 64)
	if err != nil {
		return 0, fmt.Errorf("ConvFloat64: failed to parse input %v", in)
	}
	return out, nil
}

func Bool(in string) (bool, error) {
	out, err := strconv.ParseBool(in)
	if err != nil {
		return false, fmt.Errorf("ConvBool: failed to parse input %v", in)
	}
	return out, nil
}

func Complex(in string) (complex128, error) {
	out, err := strconv.ParseComplex(in, 128)
	if err != nil {
		return 0, fmt.Errorf("ConvComplex: failed to parse input %v", in)
	}
	return out, nil
}

type PtrParser func(string) error
type Parser[T any] func(string) (T, error)

func ParsePtr[T any](ptr *T, f Parser[T]) PtrParser {
	return func(s string) error {
		var err error
		*ptr, err = f(s)
		return err
	}
}

func Generic(in []string, parsers ...PtrParser) (int, error) {
	if len(in) != len(parsers) {
		return 0, fmt.Errorf("ConvMulti: len(in) %v != len(parsers) %v\n", len(in), len(parsers))
	}
	for i, str := range in {
		err := parsers[i](str)
		if err != nil {
			return i, err
		}
	}
	return len(in), nil
}
