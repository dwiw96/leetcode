package main

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	c := "III"
	d := "LVIII"
	e := "MCMXCIV"
	f := "XCIV"

	var res int

	res = RomanToInt(c)
	if res != 3 {
		t.Fatalf("III = %d", res)
	}

	res = RomanToInt(d)
	if res != 58 {
		t.Fatalf("LVIII = %d", res)
	}

	res = RomanToInt(e)
	if res != 1994 {
		t.Fatalf("MCMXCIV = %d", res)
	}

	res = RomanToInt(f)
	if res != 94 {
		t.Fatalf("XCIV = %d", res)
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	var strs []string
	var res string

	strs = []string{"ac", "ac", "a", "a"}
	res = LongestCommonPrefix(strs)
	if res != "a" {
		t.Errorf("Failed to get \"a\" = %s", res)
	}

	strs = []string{"acc", "aaa", "aaba"}
	res = LongestCommonPrefix(strs)
	if res != "a" {
		t.Errorf("Failed to get \"a\" = %s", res)
	}

	strs = []string{"abc", "abcd", "abcd"}
	res = LongestCommonPrefix(strs)
	if res != "abc" {
		t.Errorf("Failed to get \"abc\" = %s", res)
	}

	strs = []string{"flower", "flow", "flight"}
	res = LongestCommonPrefix(strs)
	if res != "fl" {
		t.Errorf("Failed to get \"fl\" = %s", res)
	}

	strs = []string{"dog", "racecar", "car"}
	res = LongestCommonPrefix(strs)
	if res != "" {
		t.Errorf("Failed to get \"\" = %s", res)
	}

	strs = []string{"green", "re", "replay"}
	res = LongestCommonPrefix(strs)
	if res != "" {
		t.Errorf("Failed to get \"\" = %s", res)
	}

	strs = []string{"abcd", "ab", "abc", "abc"}
	res = LongestCommonPrefix(strs)
	if res != "ab" {
		t.Errorf("Failed to get \"ab\" = %s", res)
	}

	strs = []string{"abc", "abc", "", "abc"}
	res = LongestCommonPrefix(strs)
	if res != "" {
		t.Errorf("Failed to get \"\" = %s", res)
	}
}

func TestIsValid(t *testing.T) {
	var (
		s   string
		res bool
	)

	s = "(){}[]"
	res = IsValid(s)
	if res == false {
		t.Errorf("Wrong, should be %v", true)
	}

	s = "(){})(]"
	res = IsValid(s)
	if res == true {
		t.Errorf("Wrong, should be %v", false)
	}

	s = "((([{}]))){}[]"
	res = IsValid(s)
	if res == false {
		t.Errorf("Wrong, should be %v", true)
	}

	s = ")(){}[]("
	res = IsValid(s)
	if res == true {
		t.Errorf("Wrong, should be %v", false)
	}

	s = "("
	res = IsValid(s)
	if res == true {
		t.Errorf("Wrong, should be %v", false)
	}

	s = "[[[])]"
	res = IsValid(s)
	if res == true {
		t.Errorf("Wrong, should be %v", false)
	}
}
