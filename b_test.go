package main_test

import "testing"

//go test -bench=CalN* -run=999 --test.cpuprofile=cpu.out
//go tool pprof ./cpu.out
func BenchmarkCalNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calNumber()
	}
}

func calNumber() {
	sum := 0
	for i := 0; i < 100; i++ {
		sum += i
	}
}
