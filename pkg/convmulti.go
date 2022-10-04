package convm

import (
	"strconv"
	"fmt"
)

type Conv func(string, any) error

func ConvMulti(in []string, convs []Conv, out []any) error {
	if len(in) != len(convs) {
		return fmt.Errorf("ConvMulti: len(in) %v != len(convs) %v\n", len(in), len(convs))
	}
	if len(in) != len(out) {
		return fmt.Errorf("ConvMulti: len(in) %v != len(out) %v\n", len(in), len(out))
	}
	for i:=0; i<len(in); i++ {
		err := convs[i](in[i], out[i])
		if err != nil {
			return fmt.Errorf("ConvMulti: %w", err)
		}
	}
	return nil
}

func ConvStr(in string, out any) error {
	outp, ok := out.(*string)
	if !ok {
		return fmt.Errorf("ConvStr: output %v not of type *string", out)
	}
	*outp = in
	return nil
}

func ConvInt64(in string, out any) error {
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

func ConvFloat64(in string, out any) error {
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
