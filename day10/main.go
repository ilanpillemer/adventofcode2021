package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	matches = map[byte]byte{
		'[': ']',
		'(': ')',
		'<': '>',
		'{': '}',
	}

	score = map[byte]int{
		']': 57,
		')': 3,
		'>': 25137,
		'}': 1197,
	}
	score2 = map[byte]int{
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}
)

func main() {
	fmt.Println("Day 10")
	//f, _ := os.Open("sample.txt")
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	total := 0
	part2 := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		if ch, ok, rem := recurse(line, []byte{}); ok {
			total = total + score[ch]
		} else {
			lscore := int(0)
			for i := len(rem) - 1; i > -1; i-- {
				lscore = lscore * 5
				lscore = lscore + score2[rem[i]]
			}
			part2 = append(part2, lscore)
		}
	}
	fmt.Println("** part 1 **")
	fmt.Println(total)
	fmt.Println("** part 2 **")
	sort.Ints(part2)
	fmt.Println(part2[len(part2)/2])
}

func recurse(str string, seen []byte) (byte, bool, []byte) {
	if len(str) == 0 {
		return 0, false, seen
	}
	if _, ok := matches[str[0]]; ok {
		seen = append(seen, str[0])
		return recurse(str[1:], seen)
	} else {
		if str[0] == matches[seen[len(seen)-1]] {
			return recurse(str[1:], seen[:len(seen)-1])
		} else {
			return str[0], true, nil
		}
	}
}
