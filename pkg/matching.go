package convm

import (
	"fmt"
)

func ParseMatching[T any](in []string, buf []T, parse Parser[T]) (out []T, err error) {
	buf = buf[:0]
	for _, str := range in {
		val, err := parse(str)
		if err != nil {
			return buf, err
		}
		buf = append(buf, val)
	}
	return buf, nil
}

func Matching(in []string, buf any) error {
	var err error
	switch bufc := buf.(type) {
		case *[]int:
			*bufc, err = ParseMatching(in, *bufc, Int)
			return err
		case *[]int64:
			*bufc, err = ParseMatching(in, *bufc, Int64)
			return err
		case *[]float64:
			*bufc, err = ParseMatching(in, *bufc, Float64)
			return err
		case *[]bool:
			*bufc, err = ParseMatching(in, *bufc, Bool)
			return err
		case *[]string:
			*bufc, err = ParseMatching(in, *bufc, Str)
			return err
		case *[]complex128:
			*bufc, err = ParseMatching(in, *bufc, Complex)
			return err
		default:
			return fmt.Errorf("buf %v %#v not parseable", buf, buf)
	}
	return nil
}
