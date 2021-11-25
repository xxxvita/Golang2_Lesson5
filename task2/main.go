package main

import (
	"fmt"
	"sync"
)

func main() {
	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			mutex.Lock()
			defer func() {
				wg.Done()
				mutex.Unlock()
			}()
		}()
	}

	wg.Wait()
	fmt.Println("Завершшились потоки через освобождение мьютекса в defer")
}
