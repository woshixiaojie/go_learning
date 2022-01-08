/* 接口定义的原因，两个模块可能需要使用相同的功能
但是具体实现不一样。此操作是为了方便扩展。
且接口的实现不依赖接口的定义，采用duck type方式
*/
package interface_test

import (
	"testing"
)

/* 首先需要定义一个接口 interface */
type Program interface {
	WriteHelloWorld() string
}

/* 然后创建一个类型 */
type GoProgrammer struct {
}

/* 然后类型作为参数，绑定方法/实例 */
func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello,World\")"
}

func TestClient(t *testing.T) {
	/* 创建一个接口类型变量 */
	var g Program
	/* 使用接口需要传指针 */
	g = new(GoProgrammer)
	t.Log(g.WriteHelloWorld())
}
