package main

import (
	"fmt"
	"io"
	"bufio"
	"strings"
	"functional/fib"
)

func fibonacci()  intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a + b
		return a
	}
}

type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)

	//TODO: incorrect if p is to small
	return strings.NewReader(s).Read(p)
}

//类型都能实现接口
func printFileContents(reader io.Reader)  {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fib.Fibonacci()
	printFileContents(f)
}
