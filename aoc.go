package main

import (
	"os"
	"reflect"
	"strconv"
	"strings"
)

func inSlice(slice []int64, element int64) bool {
	for i := range slice {
		if slice[i] == element {
			return true
		}
	}
	return false
}

func f1(input string) {
	var result int64
	numbers := strings.Split(input, ",")
	for _, el := range numbers {
		s := strings.Trim(el, " ")
		n, _ := strconv.ParseInt(s, 10, 64)
		result += n
	}
	println(result)
}

func f1_2(input string) {
	var result int64
	seenFrequencies := []int64{0}

	numbers := strings.Split(input, ",")
	for j := 0; j < len(numbers); j++ {
		for _, el := range numbers {
			s := strings.Trim(el, " ")
			n, _ := strconv.ParseInt(s, 10, 64)
			result += n
			if inSlice(seenFrequencies, result) {
				println(result)
				return
			}
			seenFrequencies = append(seenFrequencies, result)
		}
	}
}

func main() {
	funcMap := map[string]interface{}{
		"1":   f1,
		"1_2": f1_2,
	}
	funcName := os.Args[1]
	args := os.Args[2:]

	input := make([]reflect.Value, len(args))
	for i, param := range args {
		input[i] = reflect.ValueOf(param)
	}
	f := reflect.ValueOf(funcMap[funcName])
	f.Call(input)
}
