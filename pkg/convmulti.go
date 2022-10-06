package convm

import (
	"strconv"
	"fmt"
)

type Conv func(string, any) error

func Multi(in []string, convs []Conv, out []any) (int, error) {
	if len(in) != len(convs) {
		return 0, fmt.Errorf("ConvMulti: len(in) %v != len(convs) %v\n", len(in), len(convs))
	}
	if len(in) != len(out) {
		return 0, fmt.Errorf("ConvMulti: len(in) %v != len(out) %v\n", len(in), len(out))
	}
	for i:=0; i<len(in); i++ {
		err := convs[i](in[i], out[i])
		if err != nil {
			return i, fmt.Errorf("ConvMulti: %w", err)
		}
	}
	return len(in), nil
}

func Str(in string, out any) error {
	outp, ok := out.(*string)
	if !ok {
		return fmt.Errorf("ConvStr: output %v not of type *string", out)
	}
	*outp = in
	return nil
}

func Int64(in string, out any) error {
	outp, ok := out.(*int64)
	if !ok {
		return fmt.Errorf("ConvInt64: output %v not of type *int64", out)
	}
	var err error
	*outp, err = strconv.ParseInt(in, 0, 64)
	if err != nil {
		return fmt.Errorf("ConvInt64: failed to parse input %v", in)
	}
	return nil
}

func Int(in string, out any) error {
	outp, ok := out.(*int)
	if !ok {
		return fmt.Errorf("ConvInt: output %v not of type *int", out)
	}
	i, err := strconv.ParseInt(in, 0, strconv.IntSize)
	if err != nil {
		return fmt.Errorf("ConvInt: failed to parse input %v", in)
	}
	*outp = int(i)
	return nil
}

func Float64(in string, out any) error {
	outp, ok := out.(*float64)
	if !ok {
		return fmt.Errorf("ConvFloat64: output %v not of type *float64", out)
	}
	var err error
	*outp, err = strconv.ParseFloat(in, 64)
	if err != nil {
		return fmt.Errorf("ConvFloat64: failed to parse input %v", in)
	}
	return nil
}
