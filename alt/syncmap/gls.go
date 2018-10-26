package syncmap

import (
	"github.com/jasonyuan/gls/goid"
	"sync"
)

var m = sync.Map{}

func getMap() map[string]string {
	gid := goid.GetGoid()
	if context, _ := m.Load(gid); context != nil {
		return context.(map[string]string)
	}
	context := make(map[string]string)
	m.Store(gid, context)
	return context
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
	m.Delete(goid.GetGoid())
}
