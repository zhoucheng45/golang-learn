package main

import "fmt"

func main() {
	arr := []int{1, 3, 5, 7}
	for i, v := range arr {
		fmt.Println(i, v)
	}

}
