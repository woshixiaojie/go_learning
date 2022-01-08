package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func ReturnMultiValue() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

/* 创建一个统计函数耗时的方法
参数函数，返回值也是函数
*/
func timeSpent(inner func(op int) int) func(op int) int {
	return func(op int) int {
		start := time.Now()
		ret := inner(op)
		fmt.Println("time spent :", time.Since(start))
		return ret
	}
}

/* 创建一个停止10s的函数 */
func slowFn(op int) int {
	time.Sleep(time.Second * 2)
	return 2
}

func TestFn(t *testing.T) {
	a, _ := ReturnMultiValue()
	t.Log(a)
	/* 将函数的计时功能传递给新的值 */
	newTimeSpendSlowFn := timeSpent(slowFn)
	t.Log(newTimeSpendSlowFn(2))
}

/* 创建以一个可变参数长度的函数 */
func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	sum := Sum(1, 2, 3, 4)
	sum1 := Sum(4, 5, 6, 7)
	t.Log(sum, sum1)
}

/* 创建一个用于defer调用的函数 */
func Clear(){
	fmt.Println("Clear Resources")
}

func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("Start")
	/* 就算调用了panic， defer同样也会执行 */
	panic("err")
	/* 若是其它语句将无法被执行到 */
	fmt.Println("end")
}