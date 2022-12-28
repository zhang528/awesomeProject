package main

import (
	"bytes"
	"fmt"
	"strings"
)

const (
	KB = 1024
	MB = KB * KB
	GB = MB * KB
	TB = GB * KB
	PB = TB * KB
	EB = PB * KB
	ZB = EB * KB
	YB = ZB * KB
)

func main() {
	str := fmt.Sprintf("%f", 1.1)
	println(str + "_")
}

func comma(s string) string {
	buffer := bytes.NewBufferString(s[:0])
	var prefix string = ""
	var intStr = s
	var floatStr string
	if strings.Contains(s, "+") || strings.Contains(s, "-") {
		prefix = s[:1]
		intStr = s[1:]
	}
	if strings.Contains(intStr, ".") {
		index := strings.Split(intStr, ".")
		intStr = index[0]
		floatStr = "." + index[1]
	}
	if len(intStr) <= 3 {
		return s
	}
	buffer.WriteString(prefix)
	for i, val := range intStr {
		if i != 0 && i%3 == 0 {
			buffer.WriteString(",")
		}
		buffer.WriteRune(val)
	}
	buffer.WriteString(floatStr)
	return buffer.String()
}

type intStr map[string]int

func equalsString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	str := make(intStr)
	for _, val := range s1 {
		str[string(val)]++
	}
	for _, val := range s2 {
		str[string(val)]--
	}
	for _, val := range str {
		if val != 0 {
			return false
		}
	}
	return true
}
