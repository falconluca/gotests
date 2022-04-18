package main_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func TestReadFromCloseChannel(t *testing.T) {
	ch := make(chan string, 10)

	go func(writeCh chan<- string) {
		for i := 0; i < 10; i++ {
			writeCh <- fmt.Sprintf("%d", i)
		}
		close(ch)
		log.Println("sub goroutine write done, then close")
	}(ch)

	time.Sleep(3 * time.Second)

	for r := range ch {
		log.Println(r)
	}
}

func TestWriteFromCloseChannel(t *testing.T) {
	tt := assert.New(t)

	ch := make(chan string, 1)
	close(ch)

	go func(writeCh chan<- string) {
		defer func() {
			if err := recover(); err != nil {
				// send on closed channel
				tt.NotNil(err)
			}
		}()

		writeCh <- "1"
		log.Println("sub goroutine write done, then close")
	}(ch)

	time.Sleep(3 * time.Second)
}

// ------------------------------------------------------------

type ChSt struct {
	Key,
	Value string
}

var (
	writeCh = make(chan *ChSt, 0)
	readCh  = make(chan chan *ChSt, 0)

	resourceMap = map[string]string{} // 非线程安全
)

//go test -v --run=TestChannel ./
func TestChannel(t *testing.T) {
	// TODO
	go testChGoroutine(readCh, writeCh)
	writeSt(&ChSt{Key: "1", Value: "1"})
	writeSt(&ChSt{Key: "2", Value: "2"})
	fmt.Printf("value of key=1 is %s", readSt("1").Value)
}

func readSt(key string) *ChSt {
	rCh := make(chan *ChSt, 0)
	readCh <- rCh
	rCh <- &ChSt{Key: key}
	return <-rCh
}

func writeSt(chSt *ChSt) {
	writeCh <- chSt
}

func testChGoroutine(readCh chan chan *ChSt, writeCh chan *ChSt) {
	for {
		select {
		case rCh := <-readCh:
			st := <-rCh
			st.Value = resourceMap[st.Key]
			rCh <- st
		case st := <-writeCh:
			resourceMap[st.Key] = st.Value
		}
		continue
	}
}

// ------------------------------------------------------------

// 流量控制
//go test -v --run=TestConcurrentControl ./
func TestConcurrentControl(t *testing.T) {
	// 单次最多只允许 5 个 Goroutine 并发执行
	ch := make(chan struct{}, 5)

	taskFunc := func(i int) {
		defer func() {
			<-ch
		}()
		fmt.Printf("index = %d\n", i)
	}

	for i := 0; i < 100; i++ {
		ch <- struct{}{}
		go taskFunc(i)
	}
}

//go test -v --run=TestReadChannelByRange ./
func TestReadChannelByRange(t *testing.T) {
	ch := make(chan string)
	go rangeChannel(ch)

	go func() {
		ch <- "1"
	}()
	go func() {
		ch <- "2"
	}()

	time.Sleep(time.Second)
	close(ch)
	time.Sleep(3 * time.Second)
}

func rangeChannel(ch chan string) {
LOOP:
	for {
		select {
		case str, ok := <-ch:
			if ok {
				fmt.Printf("str is %s\n", str)
				continue
			}
			fmt.Printf("inner ch is closed\n")
			break LOOP
		}
	}

	for str := range ch {
		fmt.Printf("str is %s\n", str)
	}
	fmt.Printf("ch is closed\n")

	if _, ok := <-ch; !ok {
		fmt.Printf("channel is closed\n")
	}
}
