package main_test

import (
	"fmt"
	"testing"
)

func TestNewMake(t *testing.T) {
	// 指针的零值是nil，字符串的零值是空字符
	//zeroValue()

	// Go 语言中引用类型的变量，在使用时不仅要声明，而且还要为它分配内存空间。
	// new 和 make 都是用于内存分配的。
	// new() 接收类型T返回类型的指针*T
	// make() 返回类型本身，仅用于 slice，map和chan的内存分配
	//initIntPtr()
	//initMap()

	// range map 时，不可依赖于 map 的顺序。
	// 如果需要有序 map 的话，可借助 slice 实现。
	ss := make(map[string]*struct{})
	s, _ := ss["name"]
	fmt.Println(s)
}

func initMap() {
	//var m map[string]int
	//m["name"] = 1000 // panic
	//fmt.Println(m)

	var m map[string]int
	m = make(map[string]int, 10)
	m["name"] = 1000
	fmt.Println(m)
}

func initIntPtr() {
	//var a *int
	//*a = 100 // panic
	//fmt.Println(*a)

	var a *int
	a = new(int)
	*a = 100
	fmt.Println(*a)

	//var a int
	//fmt.Println(&a)
	//var p *int
	//p = &a // 这里对p进行赋值了，因此不会 panic
	//*p = 20
	//fmt.Println(a)
}

func zeroValue() {
	var p *string
	fmt.Println(p)
	var str string
	fmt.Println(str)
}
