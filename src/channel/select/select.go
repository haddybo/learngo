package main

import (
	"fmt"
	"time"
	"math/rand"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func worker(id int, c chan int) {
	for n := range c{
		time.Sleep(time.Second)
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func main() {
	var c1, c2 = generator(), generator()
	var worder = createWorker(0)

	var values []int
	tm := time.After(10 * time.Second)
	tick := time.Tick(time.Second)
	for {

		time.Sleep(time.Second)

		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worder
			activeValue = values[0]
		}
		select {
		case n := <- c1:
			values = append(values, n)
		case n:= <- c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <- time.After(800 * time.Millisecond):    //如果800ms内还没有生成数据，则会打印timeout,原因是生成的channel是个临时变量，第一次
		//访问的时候，初始化一下，第二次访问的时候都直接用了。而且select也是如果没有分支被选中则会不断的去选择
			fmt.Println("timeout")
		case <- tm:
			fmt.Println("bye")
			return
		case <- tick:
			fmt.Printf("queue len = %d \n", len(values))
		}
	}


}
