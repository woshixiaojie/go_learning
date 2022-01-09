package goroutine

import (
	"fmt"
	"testing"
	"time"
)

/*
System Thread 系统实际线程 kernel entity

Processor go的协程处理器

Goroutine 协程
*/

func TestGoroutine(t *testing.T) {
	for i := 0; i < 10; i++ {

		/*
			使用go关键字启动协程
			这里i传递进去的都是i的值拷贝
			在每个协程里面，都是唯一的值，值存在不同的地址，所以就不存在竞争关系
		*/
		go func(i int) {
			fmt.Println(i)
		}(i)

		/*
			全部都输出10
			使用了这种方法，相当于是指针，共享内存
			协程之间存在竞争关系
		*/
		go func() {
			fmt.Println(i)
		}()
	}

	/*
		为了避免协程退出的太快，这里加一个等待时间防止退出

	*/
	time.Sleep(time.Millisecond * 50)
}
