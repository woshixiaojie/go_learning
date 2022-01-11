/* channel的广播和关闭 */
package channel_close

import (
	"fmt"
	"sync"
	"testing"
)

/*
创建一个数据的生产者
参数 装int类型的channel，以及一个同步的等待组类型指针
*/
func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		/*
			痛点，对于多个主协程，由于receiver不能准确的知道producer的实际关闭时间
			以及producer不知道多个receiver的实际需要的退出时间

			所以执行完毕则关闭通道
		*/
		close(ch)

		/* 继续往协程里放消息会出发panic */
		//ch <- 11

		wg.Done()
	}()
}

/*
数据的接收或者消费者
参数 装int类型的channel，以及一个同步的等待组类型指针
*/
func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {

		var (
			data int
			ok   bool
		)
		for i := 0; i < 10; i++ {
			/*
				关闭的通道将返回两个参数。
				通道里的值、bool
			*/
			if data, ok = <-ch; ok {
				fmt.Println(data)
			} else {
				break
			}
		}

		/* 如果我们选择继续通道里面的内容，将会返回 定义通道类型的0值 */
		fmt.Println(<-ch)

		/*
			如果你判断了返回的bool类型，则通道关闭之后将会立即返回false
			这即使channel的广播机制
			所以此处不会被打印
		*/
		if data, ok = <-ch; ok {
			fmt.Println(data)
		}

		/* 会将计数器减1 */
		wg.Done()
	}()
}

func TestChannelClose(t *testing.T) {
	/* 定义一个用于管理协程的变量 */
	var wg sync.WaitGroup
	/* 创建一个用于装int类型的channel */
	ch := make(chan int)
	/* 先+1，然后-1,保持同步 */
	wg.Add(1)
	dataProducer(ch, &wg)

	/*
		go func() {
			// 这时尽管启动两个协程，也不会导致运行失败
			// 只是由于协程的调度机制，打印会不一样
			for i := 0; i < 2; i++ {
				wg.Add(1)
				dataReceiver(ch, &wg)
			}
		}()
	*/
	wg.Add(1)
	dataReceiver(ch, &wg)

	/* 等待阻塞知道变为计数器0 */
	wg.Wait()
}
