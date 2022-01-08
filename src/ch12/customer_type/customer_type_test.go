/* 自定义类型 */
package customer_type

import (
	"fmt"
	"testing"
	"time"
)

/* 由于参数比较长，所以我们创建一个别名*/
type IntConv func(op int) int

func timeSpent(inner IntConv) IntConv {
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
	/* 将函数的计时功能传递给新的值 */
	newTimeSpendSlowFn := timeSpent(slowFn)
	t.Log(newTimeSpendSlowFn(2))
}
