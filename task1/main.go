package main

import (
	"flag"
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	nFlag := flag.Int("n", 1000, "Количество потоков по умолчанию")
	flag.Parse()

	wg := sync.WaitGroup{}
	var nGorutines int64 = 0

	wg.Add(*nFlag)
	for i := 0; i < *nFlag; i++ {
		go func() {
			defer wg.Done()

			atomic.AddInt64(&nGorutines, 1)
		}()
	}

	wg.Wait()
	fmt.Printf("Все потоки завершены (было обработано задач %d)\n", nGorutines)
}
