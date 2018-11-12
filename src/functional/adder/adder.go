package main

import "fmt"

func adder() func(int) int {
	sum := 0    //sum为自由变量
	return func(v int) int {
		sum += v    //v为局部变量
		return sum
	}
}

type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {    //正统的函数式编程，只有函数和常量
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func main() {
	a := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, a = a(i)
		fmt.Printf("0 + 1 + ... + %d = %d\n", i, s)
	}
}
