package main

import (
	"time"
	"fmt"
)

func main()  {
	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				fmt.Printf("Hello from %d\n", i)
			}
		}(i)
	}
	time.Sleep(time.Minute)
}