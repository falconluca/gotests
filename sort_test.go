package main_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestBuiltinSort(t *testing.T) {
	tt := assert.New(t)

	strs := []string{"c", "a", "b"}
	tt.False(sort.StringsAreSorted(strs))
	sort.Strings(strs)
	tt.True(sort.StringsAreSorted(strs))

	ints := []int{12, 3, 4, 1}
	tt.False(sort.IntsAreSorted(ints))
	sort.Ints(ints)
	tt.True(sort.IntsAreSorted(ints))
}

// ----------------------------------------

type SortByLen []string

func (s SortByLen) Len() int {
	return len(s)
}

func (s SortByLen) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortByLen) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func TestInterfaceSort(t *testing.T) {
	names := []string{"jamesLee", "luca", "allen"}
	sort.Sort(SortByLen(names))
	fmt.Println(names)
}
