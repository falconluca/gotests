package main_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"unsafe"
)

//go test -v --run=TestSlice .
func TestSlice(t *testing.T) {
	slice := make([]int, 0, 5)
	slice = []int{0, 1, 2, 3, 4}
	array := [5]int{0, 1, 2, 3, 4}
	testArray(array)
	testSlice(slice)
	testSlicePtr(&slice)
	t.Logf("new slice is %v\n", slice)
}

func testSlicePtr(slice *[]int) {
	*slice = append(*slice, 5) // 内存从新分配
}

// 切片是值传递
func testSlice(slice []int) {
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	fmt.Printf("%#v\n", *sh)
}

// 数组是值传递
func testArray(slice [5]int) {
	slice[4] = 6
}

//go test -v --run=TestDestructuringSlice .
func TestDestructuringSlice(t *testing.T) {
	tt := assert.New(t)

	result := make([]int, 2)
	nums := []int{1, 2, 3}
	result = append(result, nums...)

	tt.Equal(5, len(result))
}

func TestNilSlice(t *testing.T) {
	tt := assert.New(t)

	var s1 []int
	tt.Nil(s1)
	tt.Equal(0, len(s1))
}
