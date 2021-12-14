package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

var rules = map[string]string{}
var tmpl string
var pairs = map[string]int{}

func main() {
	//	load("sample.txt")
	load("input.txt")
	fmt.Println(tmpl)
	fmt.Println(rules)
	for i := 0; i < 10; i++ {
		tmpl = step(tmpl)
		log.Println(len(tmpl))
		quantities := counts(tmpl)
		sort.Ints(quantities)
		log.Println(quantities)
	}
	log.Println(len(tmpl))
	quantities := counts(tmpl)
	sort.Ints(quantities)
	log.Println(quantities)
	p1 := quantities[len(quantities)-1] - quantities[0]
	fmt.Println("**** part 1 ****")
	fmt.Println(p1)
}

func load(fname string) {
	f, _ := os.Open(fname)
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	tmpl = scanner.Text()
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "->") {
			xs := strings.Split(line, "->")
			from, to := strings.TrimSpace(xs[0]), strings.TrimSpace(xs[1])
			rules[from] = to
		}
	}
}

func step(str string) string {
	var sb strings.Builder
	for i := 0; i < len(str)-1; i++ {
		//	fmt.Println(str[i : i+2])
		sb.WriteString(str[i:i+1] + rules[str[i:i+2]])
	}
	sb.WriteString(str[len(str)-1:])
	return sb.String()
}

func counts(str string) []int {
	elems := map[rune]int{}
	counts := []int{}
	for _, ch := range str {
		elems[ch]++
	}
	for _, count := range elems {
		counts = append(counts, count)
	}
	return counts
}
