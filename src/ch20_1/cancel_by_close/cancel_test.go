/* 人物的取消 */
package cancel_by_close

import (
	"fmt"
	"testing"
	"time"
)

/*
判断通道是否收到消息
思考，为什么channel要是struct类型
*/
func isCancelled(cancelChan chan struct{}) bool {
	/* 使用select多路机制 */
	select {
	case <-cancelChan:
		return true
		/* 记住select不是顺序判断的，所以不是当没有消失时，就执行到default */
	default:
		return false

	}
}

/*  */
func cancel_1(cancelChan chan struct{}) {
	/* struct{}代表空结构, {}代表实例化这个空结构 */
	cancelChan <- struct{}{}
}

func cancel_2(cancelChan chan struct{})  {
	close(cancelChan)
}


func TestCancel(t *testing.T) {
	/*
		声明一个带有0个buffer机制的channel,所以此channel也是阻塞状态的
	*/
	cancelChan := make(chan struct{}, 0)

	/* 如果声明具有一个容量的channel，则放入channel之后，接受者将会立即感知到 */
	//cancelChan := make(chan struct{}, 1)

	/* 开启5个协程 */
	for i := 0; i < 5; i++ {
		/*
			规定函数的传参，才能传值
			把i传进去，实际上是为了判断那一次被关闭了
		*/
		go func(i int, ch chan struct{}) {
			for {
				/* 利用csp机制，增加buffer */
				if isCancelled(cancelChan) {
					/* 当装完了之后，返回了false */
					break
				}
				//time.Sleep(time.Millisecond * 3)
			}
			/* 放在里面 */
			fmt.Println(i, "Cancelled")
		}(i, cancelChan)

		//不然后放在外面，不然每次这里都会走到
		//fmt.Println(i, "Cancelled")
	}

	cancel_2(cancelChan)


	/* 避免主进程退出 */
	time.Sleep(time.Second * 1)
}
