package convm

import (
	"reflect"
	"testing"
)

func TestMulti(t *testing.T) {
	in := []string{"Hello", "23", "45", "earlobe", "35.5", "true", "33.3+29.5i"}
	var s string
	var i int64
	var i2 int
	var f float64
	var b bool
	var c complex128
	out := []any{&s, &i, &i2, Nil{}, &f, &b, &c}
	n, err := Multi(in, out...)
	if err != nil {
		panic(err)
	}
	if n != 7 {
		t.Errorf("n %v != 7", n)
	}
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
	if b != true {
		t.Errorf("b %v != true", b)
	}
	if c != 33.3 + 29.5i {
		t.Errorf("c %v != 33.3 + 29.5i", c)
	}
}

func TestMulti2(t *testing.T) {
	in := []string{"Hello", "23", "45", "35.5"}
	var s string
	var i int64
	var i2 int
	var f float64
	n, err := Multi(in, &s, &i, &i2, &f)
	if err != nil {
		panic(err)
	}
	if n != 4 {
		t.Errorf("n %v != 4", n)
	}
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

func TestMatching(t *testing.T) {
	in := []string{"3.5", "33.0", "100"}
	var buf []float64
	err := Matching(in, &buf)
	if err != nil {
		panic(err)
	}
	expect := []float64{3.5, 33.0, 100}
	if !reflect.DeepEqual(buf, expect) {
		t.Errorf("buf %v != expect %v", buf, expect)
	}
}

func TestGeneric(t *testing.T) {
	in := []string{"Hello", "23", "45", "35.5"}
	var s string
	var i int64
	var i2 int
	var f float64
	n, err := Generic(in, ParsePtr(&s, Str), ParsePtr(&i, Int64), ParsePtr(&i2, Int), ParsePtr(&f, Float64))
	if err != nil {
		panic(err)
	}
	if n != 4 {
		t.Errorf("n %v != 4", n)
	}
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

