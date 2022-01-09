package client

import (
	"ch16/series"
	"fmt"
	"testing"
)

func TestPackage(t *testing.T) {
	var (
		list []int
		err  error
	)
	if list, err = series.GetFibonacci(10); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(list)

}
