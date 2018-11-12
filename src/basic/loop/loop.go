package main

import (
	"fmt"
	"io"
	"bufio"
	"strings"
)

func printFileContents(reader io.Reader)  {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever()  {
	for {
		fmt.Println("abc")
	}
}

func main() {
	s := `abc"d
	"kkk`
	printFileContents(strings.NewReader(s))
}
