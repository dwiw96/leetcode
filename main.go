package main

import (
	"fmt"
	"leetcode/easy"
)

func main() {
	var strs = []string{"abc", "abcd", "abcd"}
	res := easy.LongestCommonPrefix(strs)
	fmt.Println("result: ", res)
}
