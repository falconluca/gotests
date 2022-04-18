package main_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//go test -v --run=TestClosureDelivery .
func TestClosureDelivery(t *testing.T) {
	tt := assert.New(t)

	i := 0
	f1 := func() {
		// Go的闭包是传指针
		i++
		tt.Equal(1, i)
	}
	f1()
	tt.Equal(1, i)
}

//go test -v --run=TestCallFunctionWithParameter .
func TestCallFunctionWithParameter(t *testing.T) {
	tt := assert.New(t)

	i := 0
	f1 := func(j int) {
		// 函数传参是传值
		j++
		tt.Equal(1, j)
	}
	f1(i)
	tt.Equal(0, i)
}
