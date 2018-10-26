package gls

import (
	"github.com/jasonyuan/gls/alt/singlelock"
	"github.com/jasonyuan/gls/alt/syncmap"
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"
)

type glsFunc struct {
	putMethod func(string, string)
	getMethod func(string) string
	cleanMethod func()
}

var glsFuncGroup = [...]glsFunc {
	{Put, Get, Clean},
	{singlelock.Put, singlelock.Get, singlelock.Clean},
	{syncmap.Put, syncmap.Get, syncmap.Clean},
}


func BenchmarkGLS5(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		runInParallel(5, glsFuncGroup[0])
	}
}

func BenchmarkSingle5(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		runInParallel(5, glsFuncGroup[1])
	}
}

func BenchmarkSyncMap5(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		runInParallel(5, glsFuncGroup[2])
	}
}

func BenchmarkGLS50(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		runInParallel(50, glsFuncGroup[0])
	}
}

func BenchmarkSingle50(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		runInParallel(50, glsFuncGroup[1])
	}
}

func BenchmarkSyncMap50(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		runInParallel(50, glsFuncGroup[2])
	}
}

func BenchmarkGLS500(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		runInParallel(500, glsFuncGroup[0])
	}
}

func BenchmarkSingle500(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		runInParallel(500, glsFuncGroup[1])
	}
}

func BenchmarkSyncMap500(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		runInParallel(500, glsFuncGroup[2])
	}
}

func BenchmarkGLS5000(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		runInParallel(5000, glsFuncGroup[0])
	}
}

func BenchmarkSingle5000(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		runInParallel(5000, glsFuncGroup[1])
	}
}

func BenchmarkSyncMap5000(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		runInParallel(5000, glsFuncGroup[2])
	}
}

func runInParallel(routineCount int, glsFunc2 glsFunc) {
	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup
	for i := 0; i < routineCount; i ++ {
		wg.Add(1)
		go checkGLS(&wg, glsFunc2)
	}
	wg.Wait()
}

const key = "key"

func checkGLS(wg *sync.WaitGroup, glsFunc2 glsFunc) {
	defer wg.Done()
	value := strconv.Itoa(rand.Int())
	glsFunc2.putMethod(key, value)
	defer glsFunc2.cleanMethod()
	for i := 0; i < 10; i ++ {
		if glsFunc2.getMethod(key) != value {
			panic("value not correct")
		}
	}
}