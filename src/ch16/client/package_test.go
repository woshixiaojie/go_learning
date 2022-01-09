package client

import (
	/*
		包的导入依赖gopath，你的路径是/user/.../gohomework/go_learning
		所以从src下的路径开始导入
	*/
	"ch16/series"
	"testing"
)

func TestPackage(t *testing.T) {
	var (
		list []int
		err  error
	)
	if list, err = series.GetFibonacci(10); err != nil {
		t.Log(err)
		return
	}
	t.Log(list)
	t.Log(series.Square(3))

}
