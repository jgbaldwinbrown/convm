package convm

import (
	"strconv"
	"fmt"
)

type Nil struct{}

func Conv(in string, ret any) error {
	switch retc := ret.(type) {
		case *string: return Str(in, retc)
		case *int64: return Int64(in, retc)
		case *int: return Int(in, retc)
		case *float64: return Float64(in, retc)
		case *bool: return Bool(in, retc)
		case *complex128: return Complex(in, retc)
		case Nil: return nil
		default: return fmt.Errorf("Conv: return value %v not a compatible pointer", retc)
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

func Str(in string, ret *string) error {
	*ret = in
	return nil
}

func Int64(in string, ret *int64) error {
	var err error
	*ret, err = strconv.ParseInt(in, 0, 64)
	if err != nil {
		return fmt.Errorf("ConvInt64: failed to parse input %v", in)
	}
	return nil
}

func Int(in string, ret *int) error {
	i, err := strconv.ParseInt(in, 0, strconv.IntSize)
	if err != nil {
		return fmt.Errorf("ConvInt: failed to parse input %v", in)
	}
	*ret = int(i)
	return nil
}

func Float64(in string, ret *float64) error {
	var err error
	*ret, err = strconv.ParseFloat(in, 64)
	if err != nil {
		return fmt.Errorf("ConvFloat64: failed to parse input %v", in)
	}
	return nil
}

func Bool(in string, ret *bool) error {
	var err error
	*ret, err = strconv.ParseBool(in)
	if err != nil {
		return fmt.Errorf("ConvBool: failed to parse input %v", in)
	}
	return nil
}

func Complex(in string, ret *complex128) error {
	var err error
	*ret, err = strconv.ParseComplex(in, 128)
	if err != nil {
		return fmt.Errorf("ConvComplex: failed to parse input %v", in)
	}
	return nil
}
