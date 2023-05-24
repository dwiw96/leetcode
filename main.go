package main

import (
	"fmt"
	"leetcode/easy"
)

func main() {
	slice := []int{1, 1, 2}
	res := easy.RemoveDuplicates(slice)
	fmt.Println("result: ", res)
}
