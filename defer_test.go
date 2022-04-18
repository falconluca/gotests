package main_test

import (
	"os"
	"testing"
)

// 避免在循环中直接使用 defer ，容易造成内存泄露！
// 可以使用 func(){defer...}() 来释放资源
//go test -v --run=TestDeferInLoop ./
func TestDeferInLoop(t *testing.T) {
	for i := 0; i < 10; i++ {
		// 错误的写法：导致内存泄露
		// f, _ := os.OpenFile("./go.mod", 0777, os.ModePerm)
		// defer f.Close()

		func() {
			f, _ := os.OpenFile("./go.mod", 0777, os.ModePerm)
			defer f.Close()
		}()
	}
}
