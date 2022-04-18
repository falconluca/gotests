package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestStringsReplace(t *testing.T) {
	tt := assert.New(t)

	val := strings.Replace("Bearer Bearer Bearer TOKEN_VALUE", "Bearer ", "", 1)
	tt.Equal("Bearer Bearer TOKEN_VALUE", val)

	val = strings.Replace("Bearer Bearer Bearer TOKEN_VALUE", "Bearer ", "", -1)
	tt.Equal("TOKEN_VALUE", val)

	val = strings.Replace("Bearer Bearer Bearer TOKEN_VALUE", "Bearer ", "", 2)
	tt.Equal("Bearer TOKEN_VALUE", val)
}
