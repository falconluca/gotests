package main_test

import (
	"reflect"
	"testing"
	"unsafe"
)

//go test -v --run=TestStringHeader ./
func TestStringHeader(t *testing.T) {
	str := "hello world"
	// Go内存地址是会变化的。C不变
	kk := (*reflect.StringHeader)(unsafe.Pointer(&str)).Len
	t.Logf("len is %d", kk)
}
