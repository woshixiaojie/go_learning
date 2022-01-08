package map_ext

import "testing"

func TestMapWithFuncValue(t *testing.T) {
	m := map[int]func(op int) int{}

	m[1] = func(op int) int {
		return op
	}
	m[2] = func(op int) int {
		return op * op
	}
	m[3] = func(op int) int {
		return op * op * op
	}
	/* 通过key来访问不同的函数 */
	t.Log(m[1](2), m[2](2), m[3](2))
}

func TestMapForSet(t *testing.T) {
	/* go可以使用map来实现set相同的特性*/
	mySet := map[int]bool{}
	mySet[1] = true
	mySet[3] = true

	n := 3
	if mySet[n] {
		t.Logf("%d is exciting", n)
	} else {
		t.Logf("%d is not exciting", n)
	}
	t.Log(len(mySet))

	delete(mySet, 1)
	n = 1
	if mySet[n] {
		t.Logf("%d is exciting", n)
	} else {
		t.Logf("%d is not exciting", n)
	}

}
