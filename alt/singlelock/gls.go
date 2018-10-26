package singlelock

import (
	"github.com/jasonyuan/gls/goid"
	"sync"
)
var gls struct {
	m map[int64]map[string]string
	sync.RWMutex
}

func init() {
	gls.m = make(map[int64] map[string] string)
}

//Single lock
func getMap() map[string]string {
	gid := goid.GetGoid()
	gls.Lock()
	defer gls.Unlock()

	if m, _ := gls.m[gid]; m != nil {
		return m
	}

	m := make(map[string]string)

	gls.m[gid] = m
	return m
}

func Get(key string) string {
	return getMap()[key]
}
func Put(key string, v string) {
	getMap()[key] = v
}
func Delete(key string) {
	delete(getMap(), key)
}

func Clean() {
	gls.Lock()
	defer gls.Unlock()
	delete(gls.m, goid.GetGoid())
}
