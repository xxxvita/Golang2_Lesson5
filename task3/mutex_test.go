package mutex_test

// go test -bench=. mutex_test.go
import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

var (
	i       int
	mutex   sync.Mutex
	mutexRW sync.RWMutex
)

func AccessFromMutex(percentRead int) int {
	mutex.Lock()
	defer mutex.Unlock()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rnd := r.Intn(percentRead)

	if rnd <= percentRead {
		return i
	} else {
		i++
	}

	return i
}

func AccessFromRWMutex(percentRead int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rnd := r.Intn(percentRead)

	if rnd <= percentRead {
		mutexRW.RLock()
		defer mutexRW.RUnlock()
	} else {
		mutexRW.Lock()
		defer mutexRW.Unlock()

		i++
	}

	return i
}

func BenchmarkMutex(b *testing.B) {

	b.Run(fmt.Sprintf("%%_ReadWrite_90_10"), func(b *testing.B) {
		// Установка 100*GOMAXPROCS горутин
		b.SetParallelism(100)

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				AccessFromMutex(90)
			}
		})
	})
	b.Run(fmt.Sprintf("%%_ReadWrite_50_50"), func(b *testing.B) {
		// Установка 100*GOMAXPROCS горутин
		b.SetParallelism(100)

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				AccessFromMutex(50)
			}
		})
	})
	b.Run(fmt.Sprintf("%%_ReadWrite_10_90"), func(b *testing.B) {
		// Установка 100*GOMAXPROCS горутин
		b.SetParallelism(100)

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				AccessFromMutex(10)
			}
		})
	})
}

func BenchmarkRWMutex(b *testing.B) {
	b.Run(fmt.Sprintf("%%_ReadWrite_90_10"), func(b *testing.B) {
		// Установка 100*GOMAXPROCS горутин
		b.SetParallelism(100)

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				AccessFromRWMutex(90)
			}
		})
	})
	b.Run(fmt.Sprintf("%%_ReadWrite_50_50"), func(b *testing.B) {
		// Установка 100*GOMAXPROCS горутин
		b.SetParallelism(100)

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				AccessFromRWMutex(50)
			}
		})
	})
	b.Run(fmt.Sprintf("%%_ReadWrite_10_90"), func(b *testing.B) {
		// Установка 100*GOMAXPROCS горутин
		b.SetParallelism(100)

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				AccessFromRWMutex(10)
			}
		})
	})
}
