package operator_test

import "testing"

const(
	Readable = 1 << iota
	Writeable
	Executable
)

func TestCompareArray(t *testing.T){
	a := [...]int{1,2,3,4}
	b := [...]int{1,2,4,3}
	//c := [...]int{1,2,3,4,5}
	d := [...]int{1,2,3,4}
	t.Log(a == b) // go只允许相同长度的数组之间判断
	//t.Log(a == c) // invalid operation: a == c (mismatched types [4]int and [5]int)
	t.Log(a == d)
}

func TestBitClear(t *testing.T){
	a := 7 //0111
	//a = a&^Readable // 0111&^0001 = 0110 = 6 (按位清零，相当于是清除掉这一位)
	a = a&^Writeable // 0111&^0010 = 0101 = 5
	a = a&^Executable // 0111&^0100 = 0011 = 3
	t.Log(a, Readable, Writeable, Executable)
	t.Log(a&Readable == Readable, a&Writeable == Writeable,
		a&Executable == Executable)
}