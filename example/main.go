package main

import (
	"fmt"
	"sync"

	"github.com/Avalanche-io/counter"
)

func main() {

	c := counter.New()
	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				c.Up()
			}
		}()
	}
	wg.Wait()

	fmt.Printf("c = %d\n", c.Get())

}
