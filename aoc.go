package main

import (
	"os"
	"reflect"
	"regexp"
	"sort"
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

func getSleepTotalsAndMinutes(rows []string) (map[int]int, map[int]map[int]int) {
	sort.Strings(rows)
	re := regexp.MustCompile(`\[(?P<date>\d{4}-\d{2}-\d{2}) (?P<hour>\d{2}):(?P<min>\d{2})\] (?P<text>.+)`)

	guardRe := regexp.MustCompile(`Guard #(?P<id>\d+) begins shift`)

	sleepTotals := make(map[int]int)
	sleepMinutes := map[int]map[int]int{}
	currentGuard := -1
	sleepMin := 0
	for _, row := range rows {
		match := re.FindStringSubmatch(row)
		min, _ := strconv.Atoi(match[3])
		logStr := match[4]
		if strings.Contains(logStr, "Guard") {
			guardMatch := guardRe.FindStringSubmatch(logStr)
			currentGuard, _ = strconv.Atoi(guardMatch[1])
			sleepMin = 0
		} else if logStr == "falls asleep" {
			sleepMin = min
		} else {
			if sleepTotals[currentGuard] == 0 {
				sleepMinutes[currentGuard] = map[int]int{}
			}
			sleepTotals[currentGuard] += min - sleepMin
			for i := sleepMin; i < min; i++ {
				sleepMinutes[currentGuard][i]++
			}
		}
	}
	return sleepTotals, sleepMinutes
}

func f4(input string) {
	rows := strings.Split(strings.Trim(input, "\n"), "\n")

	sleepTotals, sleepMinutes := getSleepTotalsAndMinutes(rows)

	max := 0
	maxID := -1
	for id, total := range sleepTotals {
		if total > max {
			max = total
			maxID = id
		}
	}
	max = 0
	maxMin := -1
	for min, total := range sleepMinutes[maxID] {
		if total > max {
			max = total
			maxMin = min
		}
	}
	println(maxID * maxMin)
}

func f4_2(input string) {
	rows := strings.Split(strings.Trim(input, "\n"), "\n")

	_, sleepMinutes := getSleepTotalsAndMinutes(rows)

	max := 0
	maxMin := -1
	maxID := -1
	for id, minuteMap := range sleepMinutes {
		for min, total := range minuteMap {
			if total > max {
				max = total
				maxMin = min
				maxID = id
			}
		}
	}
	println(maxID * maxMin)
}

func main() {
	funcMap := map[string]interface{}{
		"1":   f1,
		"1_2": f1_2,
		"2":   f2,
		"2_2": f2_2,
		"3":   f3,
		"3_2": f3_2,
		"4":   f4,
		"4_2": f4_2,
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
