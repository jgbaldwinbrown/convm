package convm

import (
	"testing"
)

func TestConvMulti(t *testing.T) {
	in := []string{"Hello", "23", "45", "35.5"}
	var s string
	var i int64
	var i2 int
	var f float64
	out := []any{&s, &i, &i2, &f}
	ConvMulti(in, []Conv{ConvStr, ConvInt64, ConvInt, ConvFloat64}, out)
	if s != "Hello" {
		t.Errorf("s %v != \"Hello\"", s)
	}
	if i != 23 {
		t.Errorf("i %v != 23", i)
	}
	if i2 != 45 {
		t.Errorf("i2 %v != 45", i2)
	}
	if f != 35.5 {
		t.Errorf("f %v != 35.5", f)
	}
}
