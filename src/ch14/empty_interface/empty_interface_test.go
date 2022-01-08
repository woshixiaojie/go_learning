package empty_interface

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}){
/*	if i, ok := p.(int); ok{
		fmt.Println("integer", i)
		return
	}
	if s, ok := p.(string); ok{
		fmt.Println("string", s)
	}
	fmt.Printf("%T\n",p)*/

	/* 使用switch可以简化 */
	switch p.(type){
	case int:
		fmt.Println("integer", p)
	case string:
		fmt.Println("string", p)
	default:
		fmt.Println("string", p)

	}
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(10)
	DoSomething("10")
}