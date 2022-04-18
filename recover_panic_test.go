package main_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

// 开启 goroutine 后，必须使用 defer/recover 对 goroutine 进行 panic 监控。
// 如果子 goroutine 发生 panic 之后，没有进行 recover，那么就会导致整个进程直接挂掉！
//go test -v --run=TestRecoverPanic .
func TestRecoverPanic(t *testing.T) {
	tt := assert.New(t)

	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("[recover]: %+v\n", err)
		}
		tt.NotNil(err)
	}()
	panic("produce panic")
}

//go test -v --run=TestRecoverPanicInsideGoroutine .
func TestRecoverPanicInsideGoroutine(t *testing.T) {
	tt := assert.New(t)

	go func() {
		defer func() {
			err := recover()
			if err != nil {
				fmt.Printf("[recover] err: %v\n", err)
			}
			tt.NotNil(err)
		}()

		f, _ := os.OpenFile("./go.mod", os.O_RDONLY, os.ModePerm)
		defer f.Close()

		panic("Whoooooops!")
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("done")
}
