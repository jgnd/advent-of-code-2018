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
	numbers := strings.Split(input, "\n")
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

	numbers := strings.Split(input, "\n")
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

func hasDoublesOrTriples(id string) (bool, bool) {
	charCounts := map[byte]int{}
	for i := range id {
		charCounts[id[i]]++
	}
	hasDouble := false
	hasTriple := false
	for _, v := range charCounts {
		if v == 2 {
			hasDouble = true
		}
		if v == 3 {
			hasTriple = true
		}
	}
	return hasDouble, hasTriple
}

func f2(input string) {
	ids := strings.Split(input, "\n")
	twos := 0
	threes := 0
	for _, el := range ids {
		hasTwo, hasThree := hasDoublesOrTriples(el)
		if hasTwo {
			twos++
		}
		if hasThree {
			threes++
		}
	}
	println(twos * threes)
}

func f2_2(input string) {
	ids := strings.Split(strings.Trim(input, "\n"), "\n")
	for i, id1 := range ids {
		for _, id2 := range ids[i:] {
			diffCount := 0
			matchingChars := ""
			for k := range id1 {
				if id1[k] != id2[k] {
					diffCount++
				} else {
					matchingChars += string(id1[k])
				}
			}
			if diffCount == 1 {
				println(matchingChars)
			}
		}
	}
}

func main() {
	funcMap := map[string]interface{}{
		"1":   f1,
		"1_2": f1_2,
		"2":   f2,
		"2_2": f2_2,
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
