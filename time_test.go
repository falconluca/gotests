package main_test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestTimeTicker(t *testing.T) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancelFunc()

	ticker := time.NewTicker(2 * time.Second)
LOOP:
	for {
		select {
		case t := <-ticker.C:
			fmt.Printf("time: %s\n", t)
		case <-ctx.Done():
			break LOOP
		}
	}
	fmt.Println("done")
}

func TestDuration(t *testing.T) {
	fmt.Println(time.Duration(86400))
	fmt.Println(time.Duration(86400) * time.Second)
}
