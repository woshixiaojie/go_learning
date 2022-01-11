/*
select 多路选择、超时设置
*/
package select_test

import (
	"fmt"
	"testing"
	"time"
)

/* 定义一个客户端的服务 */
func service() string {
	/* 由于有一个等待时间，所以会被阻塞在此 */
	time.Sleep(time.Millisecond * 200)
	return "service Done"
}

/* 创建一个具有buffer等待机制的异步服务 */
func AsyncService() chan string {
	retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("AsyncService returned result")
		retCh <- ret
		fmt.Println("service process exited")
	}()
	return retCh
}

func TestAsyncService(t *testing.T) {

	/* 我们可以定义一个超时设置，来避免主进程被阻塞 */
	/*
		select 不会按照重上到下的顺序执行
			1、当case阻塞，如果有default则走，尽管时间非常短
			2、当case阻塞，如果没有default，则继续等待

	*/
	select {

	/* 正常返回，打印日志返回结果 */
	case retCh := <-AsyncService():
		t.Log(retCh)

		/* 超时则打印错误 */
	case <-time.After(time.Millisecond * 100):
		t.Error("tome out")

		/* 尽管阻塞的时间很短，也会走到default */
	default:
		t.Log("case recover")
	}

}
