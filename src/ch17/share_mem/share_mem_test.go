/*
共享内存并发机制
*/
package share_mem

import (
	"sync"
	"testing"
	"time"
)

/* 不安全的线程 */
func TestCounter(t *testing.T) {

	counter := 0

	/* 循环5000次，相当于启动了5000个协程 */
	for i := 0; i < 5000; i++ {

		go func() {
			/* 直接这样写会导致线程之间处于竞争 */
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	/* 输出并不是5000，因为共享内存的机制 */
	t.Logf("counter = %d", counter)
}

/*
安全的线程
通过互斥锁来保证线程安全
*/
func TestCounterTreadSafe(t *testing.T) {

	/* 定义一个锁 */
	var mut sync.Mutex

	counter := 0

	/* 循环5000次，相当于启动了5000个协程 */
	for i := 0; i < 5000; i++ {
		go func() {

			/* 使用defer保证锁忘记被释放，导致程序被挂起 */
			defer func() {
				mut.Unlock()
			}()

			/* 上锁 */
			mut.Lock()

			counter++
		}()
	}

	/*
	加上等待是因为，test最外围的协程处理的速度，快于循环的协程速度。
	不然会导致主程序退出了，打印的数字也不符合预期
	*/
	time.Sleep(1 * time.Second)

	/* 这时counter的值是符合预期 */
	t.Logf("counter = %d", counter)

}


/*
通过增加等待组，来实现当全部执行完毕后，继续往下执行
*/
func TestCounterWaitGroup(t *testing.T) {

	var wg sync.WaitGroup
	var mut sync.Mutex
	counter := 0

	/* 循环5000次，相当于启动了5000个协程 */
	for i := 0; i < 5000; i++ {

		/* 每启动一个协程，新增一个等待 */
		wg.Add(1)
		go func() {

			/* 使用defer保证锁忘记被释放，导致程序被挂起 */
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++

			/* 执行完毕即down */
			wg.Done()
		}()
	}

	/* 等待直到，计数为0 */
	wg.Wait()
	t.Logf("counter = %d", counter)
}