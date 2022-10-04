package convm

import (
	"testing"
)

func TestConvMulti(t *testing.T) {
	in := []string{"Hello", "23", "35.5"}
	var s string
	var i int64
	var f float64
	out := []any{&s, &i, &f}
	ConvMulti(in, []Conv{ConvStr, ConvInt64, ConvFloat64}, out)
	if s != "Hello" {
		t.Errorf("s %v != \"Hello\"", s)
	}
	if i != 23 {
		t.Errorf("i %v != 23", i)
	}
	if f != 35.5 {
		t.Errorf("f %v != 35.5", f)
	}
}
