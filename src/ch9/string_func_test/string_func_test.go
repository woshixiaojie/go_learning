package string_func_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	for idx, val := range s {
		t.Log(idx, val)
	}
	parts := strings.Split(s, ",")
	for idx, val := range parts {
		t.Log(idx, val)
	}
	t.Log(strings.Join(parts, "-"))
}

func TestStringConvert(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("Str" + s)
	val, err := strconv.Atoi("10")
	if err == nil {
		t.Log(10 + val)
	}

}