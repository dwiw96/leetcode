package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// a := []int{4, 3, 1, 5, 9}
	// b := 8
	// fmt.Println("twoSum = ", twoSum(a, b))

	/*c := "III"
	d := "LVIII"
	e := "MCMXCIV"
	f := "XCIV"

	fmt.Printf("romanToInt = \n%d \n%d \n%d  \n", RomanToInt(c), RomanToInt(d), RomanToInt(e))
	fmt.Printf("%d\n", RomanToInt(f))*/

	var strs = []string{"abc", "abcd", "abcd"}
	res := LongestCommonPrefix(strs)
	fmt.Println("result: ", res)
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, j := range nums {
		if x, y := m[target-j]; y == true {
			return []int{x, i}
		}
		m[j] = i
	}
	return []int{0, 0}
}

func RomanToInt(s string) int {
	var m = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	var res int

	length := len(s)
	last_M := m[string(s[length-1])]
	var map_2 int

	if length == 0 {
		return 0
	}
	if length == 1 {
		return last_M
	}

	for i := length - 1; i >= 0; i-- {
		map_1 := m[string(s[i])]
		if map_1 < map_2 {
			res -= map_1
		} else {
			res += map_1
		}
		map_2 = map_1
	}

	return res
}

func LongestCommonPrefix(strs []string) string {
	var res strings.Builder
	sort.Strings(strs)
	first := strs[0]
	last := strs[len(strs)-1]

	if len(strs) == 1 {
		return string(strs[0])
	}

	if len(strs) > 0 {
		for i := 0; i < len(first); i++ {
			if first[i] == last[i] {
				res.WriteString(string(first[i]))
				fmt.Println("res: ", res)
			} else {
				break
			}
		}
	}

	return res.String()
}

func IsValid(s string) bool {
	m := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	var stack []rune

	if len(s)%2 == 1 {
		return false
	}

	for _, char := range s {
		if _, isOk := m[char]; isOk {
			stack = append(stack, char)
		} else if len(stack) == 0 {
			return false
		} else if m[stack[len(stack)-1]] != char {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return true
}
