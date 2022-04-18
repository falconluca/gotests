package main

import (
	"bytes"
	"fmt"
	"github.com/cch123/elasticsql"
	"reflect"
	"sync"
	"time"
)

func main() {
	//lenCapInSlice()
	//closureGoroutine()

	sql := "select * from `bank` where `address` = \"mill\" and age != 38 "
	dsl, _, _ := elasticsql.Convert(sql)
	fmt.Println(dsl)

	//name := "Luca"
	//var name struct{}
	name := new(struct{})
	val := reflect.ValueOf(name)
	fmt.Printf("kind: %s\n", val.Kind())
	fmt.Printf("is nil?: %t\n", val.IsNil())

	for _, c := range "Aa" {
		fmt.Println(string(c), c)
	}
}

func closureGoroutine() {
	ch := make(chan int)

	policy := SafePolicy()
	for i := 0; i < 100; i++ {
		go func(ch chan int) {
			res := policy()
			ch <- res
		}(ch)
	}

	cnt := 0
LOOP:
	for {
		select {
		case re := <-ch:
			fmt.Println(re)
			cnt = cnt + re
		case <-time.After(1 * time.Second):
			break LOOP
		}
	}

	fmt.Printf("main goroutine done: %v\n", cnt)
}

func SafePolicy() func() int {
	pos := 0
	var m sync.Mutex
	return func() int {
		m.Lock()
		defer m.Unlock()
		pos++
		return pos
	}
}

func NotSafePolicy() func() int {
	pos := 0
	return func() int {
		pos++
		return pos
	}
}

func lenCapInSlice() {
	slice1()
	//slice2() // FIXME https://www.topgoer.com/go%E5%9F%BA%E7%A1%80/%E5%88%87%E7%89%87Slice.html
	//slice3()

	// Go数组在数组赋值和函数传参都是值复制，会造成消耗大量内存
	//array1()

	// nil切片。函数发生异常
	//var nilSlice []int
	//fmt.Printf("slice: %v, len: %v, cap: %v\n",
	//	nilSlice, len(nilSlice), cap(nilSlice))
	//// 空切片。数据库查询
	//emptySlice := make([]int, 0)
	//fmt.Printf("slice: %v, len: %v, cap: %v\n",
	//	emptySlice, len(emptySlice), cap(emptySlice))
	//emptySlice2 := []int{}
	//fmt.Printf("slice: %v, len: %v, cap: %v\n",
	//	emptySlice2, len(emptySlice2), cap(emptySlice2))
	//// append, len, cap 对于 nil 切片和空切片效果都是一样的
}

func testArrayPoint(x *[]int) {
	fmt.Printf("func Array : %p , %v\n", x, *x)
	(*x)[1] += 100
}

func array1() {
	arr1 := [2]int{100, 200}
	fmt.Printf("arr1: %p, %v\n", &arr1, arr1)

	// 数组赋值
	var arr2 = arr1
	fmt.Printf("arr2: %p, %v\n", &arr2, arr2)

	// 函数传参
	testArr := func(arr [2]int) {
		fmt.Printf("func arr: %p, %v\n", &arr, arr)
	}
	testArr(arr1)
}

func slice3() {
	nums := make([]int, 0, 3)
	for i := 0; i < 5; i++ {
		nums = append(nums, i)
		fmt.Printf("cap %v, len %v, ptr %p\n", cap(nums), len(nums), nums)
	}
	// Output:
	// cap 3, len 1, ptr 0xc0000b4000
	// cap 3, len 2, ptr 0xc0000b4000
	// cap 3, len 3, ptr 0xc0000b4000
	// cap 6, len 4, ptr 0xc0000aa060
	// cap 6, len 5, ptr 0xc0000aa060
}

func slice2() {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/') // 4

	dir1 := path[:sepIndex:sepIndex]
	fmt.Println(string(dir1))
	dir2 := path[sepIndex+1:]
	fmt.Println(string(dir2))

	dir1 = append(dir1, "suffix"...)
	fmt.Println(string(dir1))
	fmt.Println(string(dir2))
}

func slice1() {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/') // 4

	dir1 := path[:sepIndex]
	fmt.Println(string(dir1)) // AAAA
	dir2 := path[sepIndex+1:]
	fmt.Println(string(dir2)) // /BBBBBBBBB

	dir1 = append(dir1, "suffix"...)
	fmt.Println(string(dir1)) // AAAAsuffix
	fmt.Println(string(dir2)) // uffixBBBB
}
