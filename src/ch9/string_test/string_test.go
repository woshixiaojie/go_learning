package string_test

import "testing"

/*string是一个数据类型，不是指针或者引用类型
  string的byte数组可以存放任意数据
*/
func TestStringAccess(t *testing.T) {
	var s string
	t.Log(s)
	s = "hello"
	t.Log(s, len(s))
	//s[1] = "s" // 字符串是一个只读的 byte slice，不能被修改
	s = "\xE4\xB8\xA5"
	t.Log(s, len(s))
	t.Logf("%x", s) // e4b8a5
	s = "\xE4\xB8\xA5\xBB"
	t.Log(s, len(s))
}

func TestStringUnicodeAndUTF8(t *testing.T) {
	s := "中"
	t.Log(len(s))

	c := []rune(s)
	t.Log(c, len(c))

	t.Logf("中`s unicode %x", c)
	t.Logf("中`s utf-8 %x", s)
}

func TestStringToRune(t *testing.T) {
	s := "中华人名共和国"
	for _, v := range s {
		t.Logf("%[1]c, %[1]x", v)
	}
}
