package main

import (
	"os"
	"reflect"
	"regexp"
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

type rect struct {
	left   int
	right  int
	top    int
	bottom int
}

type point struct {
	x int
	y int
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func intersection(r1 rect, r2 rect) rect {
	return rect{
		left:   max(r1.left, r2.left),
		top:    max(r1.top, r2.top),
		right:  min(r1.right, r2.right),
		bottom: min(r1.bottom, r2.bottom),
	}
}

func overlapArea(r1 rect, r2 rect) int {
	i := intersection(r1, r2)
	w := max(0, i.right-i.left)
	h := max(0, i.bottom-i.top)
	return w * h
}

func extractIDAndRect(re *regexp.Regexp, idStr string) (int, rect) {
	match := re.FindStringSubmatch(idStr)
	id, _ := strconv.Atoi(match[1])
	x, _ := strconv.Atoi(match[2])
	y, _ := strconv.Atoi(match[3])
	w, _ := strconv.Atoi(match[4])
	h, _ := strconv.Atoi(match[5])
	r := rect{
		left:   x,
		top:    y,
		right:  x + w,
		bottom: y + h,
	}
	return id, r
}

func f3(input string) {
	ids := strings.Split(strings.Trim(input, "\n"), "\n")
	re := regexp.MustCompile(`#(?P<id>\d+) @ (?P<x>\d+),(?P<y>\d+): (?P<w>\d+)x(?P<h>\d+)`)

	overlappingPoints := make(map[point]bool)
	rects := []rect{}
	for _, id := range ids {
		_, r := extractIDAndRect(re, id)
		for _, s := range rects {
			i := intersection(r, s)
			for v := i.left; v < i.right; v++ {
				for u := i.top; u < i.bottom; u++ {
					overlappingPoints[point{x: v, y: u}] = true
				}
			}
		}
		rects = append(rects, r)
	}
	totalOverlap := 0
	for _, v := range overlappingPoints {
		if v == true {
			totalOverlap++
		}
	}
	println(totalOverlap)
}

func f3_2(input string) {
	ids := strings.Split(strings.Trim(input, "\n"), "\n")
	re := regexp.MustCompile(`#(?P<id>\d+) @ (?P<x>\d+),(?P<y>\d+): (?P<w>\d+)x(?P<h>\d+)`)

	rects := make(map[int]rect)
	for _, idStr := range ids {
		id, r := extractIDAndRect(re, idStr)
		rects[id] = r
	}
	for id, r := range rects {
		noOverlap := true
		for id2, s := range rects {
			if id != id2 && overlapArea(r, s) > 0 {
				noOverlap = false
				break
			}
		}
		if noOverlap {
			println(id)
			return
		}
	}
}

func main() {
	funcMap := map[string]interface{}{
		"1":   f1,
		"1_2": f1_2,
		"2":   f2,
		"2_2": f2_2,
		"3":   f3,
		"3_2": f3_2,
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
