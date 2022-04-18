package main_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//go test -v --run=TestAccessAndAssignmentInNilMap .
func TestAccessAndAssignmentInNilMap(t *testing.T) {
	tt := assert.New(t)

	var m map[string]map[string]map[string]*struct{}

	// access entry in nil map
	_, ok := m[""]
	tt.False(ok, `access nil map(m[""]) won't produce panic`)
	_, ok = m[""][""]
	tt.False(ok, `access nil map(m[""][""]) won't produce panic`)
	_, ok = m[""][""][""]
	tt.False(ok, `access nil map(m[""][""][""]) won't produce panic`)

	// assignment to entry in nil map
	tt.Panics(func() {
		m[""][""][""] = &struct{}{}
	}, "assignment to entry in nil map should produce panic")
}

//go test -v --run=TestRangeMap .
func TestRangeMap(t *testing.T) {
	// map range 无序
	m := map[string]string{"key1": "value1", "key2": "value2", "key3": "value3"}
	for key, value := range m {
		t.Logf("key = %s value = %s", key, value)
	}
	// Output:
	// map_test.go:33: key = key2 value = value2
	// map_test.go:33: key = key3 value = value3
	// map_test.go:33: key = key1 value = value1
}
