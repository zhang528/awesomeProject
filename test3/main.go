package main

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

func main() {
	a := []string{"s", "a", "a", "s", "d", "z", "a", "z", "v", "w", "w", "a", "a"}
	RemoveAdjacentString(&a)
	b := "asdf    sadf  sd f a   sfd    a"
	removeEmpty(&b)
	fmt.Println(b)
	fmt.Println(a)
	println(lengthOfLongestSubstring("abcabcbb"))
}

// 去除相邻的string
func RemoveAdjacentString(s *[]string) {
	var str = *s
	if len(str) <= 1 {
		return
	}
	temp := str[:1]
	for i := 1; i < len(str); i++ {
		if strings.Compare(temp[len(temp)-1], str[i]) != 0 {
			temp = append(temp, str[i])
		}
	}
	*s = temp
}
func removeEmpty(a *string) {
	S := *a
	str := string(S[0])
	end := 0
	for _, s := range S {
		last := rune(str[end])
		if unicode.IsSpace(last) && unicode.IsSpace(s) {
			continue
		}
		str += string(s)
		end++
	}
	*a = str[1:]
}
func lengthOfLongestSubstring(s string) int {
	var mmap map[byte]int
	mmap = make(map[byte]int, len(s))
	var left = 0
	var max = math.MinInt
	for i := 0; i < len(s); i++ {
		_, ok := mmap[s[i]]
		if !ok {
			mmap[s[i]] = 1
		} else {
			mmap[s[i]]++
		}
		for mmap[s[i]] > 1 {
			mmap[s[left]]--
			left++
		}
		if i-left+1 > max {
			max = i - left + 1
		}
	}
	return max
}

type ListNode struct {
	Val  int
	Next *ListNode
}
