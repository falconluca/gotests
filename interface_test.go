package main_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

//go test -v --run=TestNullInterface .
func TestNullInterface(t *testing.T) {
	tt := assert.New(t)

	type ITest interface{}
	var it ITest
	tt.Nil(it)

	f := func() {}
	it = f
	tt.NotNil(it)
}

// ------------------------------------------------

type (
	TestI interface {
		Api()
	}

	ApiS struct {
		Field string
	}
)

func (apiS *ApiS) Api() {
	fmt.Println("executing api")
}

//go test -v --run=TestNullInterface2 ./
func TestNullInterface2(t *testing.T) {
	var apis *ApiS
	var testI TestI = apis

	// Go只有当interface的类型和值都为nil时，interface才为nil
	switch {
	case testI == nil:
		t.Log("it is nil")
	default:
		t.Log("it is NOT nil")
	}

	if interfaceIsNil(testI) {
		t.Log("it is nil, checking by interfaceIsNil")
		testI.Api()
		testI.(*ApiS).Api()
		//_ = testI.(*ApiS).Field // show panic
	}
	testI.Api()
}

func interfaceIsNil(val interface{}) (flag bool) {
	if !reflect.ValueOf(val).IsValid() {
		return true
	}
	if isReferenceType(reflect.ValueOf(val)) && reflect.ValueOf(val).IsNil() {
		return true
	}
	return
}

func isReferenceType(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.Ptr, reflect.Chan, reflect.Interface, reflect.Func, reflect.Map, reflect.Slice:
		return true
	default:
		return false
	}
}

func TestTypeSwitch(t *testing.T) {
	tt := assert.New(t)

	var str interface{} = "Greetings"
	_, ok := str.(int)
	tt.False(ok)

	_, ok = str.(string)
	tt.True(ok)
}

func TestInterfaceToString(t *testing.T) {
	tt := assert.New(t)

	InterfaceToString := func(v interface{}) string {
		var ret string
		//fmt.Println(v.(type))
		switch v.(type) {
		case []byte:
			ret = string(v.([]byte))
		case string:
			ret = v.(string)
		case []interface{}:
			ret = ""
		case error:
			return ""
		default:
			ret = ""
		}
		return ret
	}

	tt.Equal("", InterfaceToString(1000))

	tt.Equal("", InterfaceToString(false))

	var msg interface{} = "我明白了哈"
	tt.Equal("我明白了哈", InterfaceToString(msg))

	tt.Equal("我明白了", InterfaceToString([]byte("我明白了")))
}
