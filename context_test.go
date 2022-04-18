package main_test

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

//go test -v --run=TestContext1 ./
func TestContext1(t *testing.T) {
	wg := sync.WaitGroup{}

	taskFunc := func(ctx context.Context, id int) {
		defer wg.Done()

		ticker := time.NewTicker(time.Second)
	LOOP:
		for {
			select {
			case <-ticker.C:
				fmt.Printf("id = %d, the ticks are delivered\n", id)
				continue
			case <-ctx.Done():
				break LOOP
			}
		}
		fmt.Printf("goroutine(id = %d) done\n", id)
	}

	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(2)
	go taskFunc(ctx, 1)
	go taskFunc(ctx, 2)

	time.Sleep(3 * time.Second)

	// 防止Goroutine泄露
	cancel()

	wg.Wait()
	fmt.Println("done")
}

//go test -v --run=TestContextWithTimeout ./
func TestContextWithTimeout(t *testing.T) {
	wg := sync.WaitGroup{}

	taskFunc := func(ctx context.Context, id int) {
		defer wg.Done()

		ticker := time.NewTicker(time.Second)
	LOOP:
		for {
			select {
			case <-ticker.C:
				fmt.Printf("id = %d, the ticks are delivered\n", id)
				continue
			case <-ctx.Done():
				fmt.Printf("goroutine(id = %d) done, err: %+v\n", id, ctx.Err())
				break LOOP
			}
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel() // TODO ctx 设置 timeout 之后，还需要调用 cancel() 吗？

	wg.Add(2)
	go taskFunc(ctx, 1)
	go taskFunc(ctx, 2)

	wg.Wait()
	fmt.Println("done")
}
