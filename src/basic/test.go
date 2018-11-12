package main

import "fmt"

func main() {
	name := "haddybo"
	banjue := name[3:]
	arr := []int{1, 2, 3, 4, 5}
	arr2 := arr[2:]
	arr = append(arr, 6)
	arr2 = append(arr2, 3)
	fmt.Println(arr2)
	fmt.Println(banjue)
}
