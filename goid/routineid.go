package goid

import (
	"reflect"
	"unsafe"
)

func getG() unsafe.Pointer
func getGInterface() interface{}

var goIDOffset uintptr

func init() {
	g := getGInterface()
	if f, ok := reflect.TypeOf(g).FieldByName("goid"); ok {
		goIDOffset = f.Offset
	} else {
		panic("can not find g.goid field")
	}
}

func GetGoid() int64 {
	g := getG()
	p := (*int64)(unsafe.Pointer(uintptr(g) + goIDOffset))
	return *p
}
