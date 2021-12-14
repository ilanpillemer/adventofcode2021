package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var rules = map[string]string{}
var tmpl string
var pairs = map[string]int64{}

func main() {
	//load("sample.txt")
	load("input.txt")
	fmt.Println(tmpl)
	pairs = start(pairs, tmpl)
	fmt.Println(pairs)
	for i := 0; i < 40; i++ {
		pairs = step(pairs)

	}
	xs := counts(pairs, tmpl[len(tmpl)-1])

	sort.Slice(xs, func(i, j int) bool { return xs[i] < xs[j] })
	fmt.Println(xs[len(xs)-1] - xs[0])
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

func start(in map[string]int64, str string) map[string]int64 {
	for i := 0; i < len(str)-1; i++ {
		in[str[i:i+2]]++
	}
	return in
}

func step(in map[string]int64) map[string]int64 {
	out := map[string]int64{}
	for k, v := range in {
		if v != 0 {
			out[k[:1]+rules[k]] += v
			out[rules[k]+k[1:]] += v
		}
	}
	return out
}

func counts(in map[string]int64, last byte) []int64 {
	xs := []int64{}
	cx := map[byte]int64{}
	for k, v := range in {
		if v != 0 {
			cx[k[0]] = cx[k[0]] + v
		}
	}
	cx[last]++
	total := int64(0)
	for _, v := range cx {
		xs = append(xs, v)
		total += v
	}
	return xs
}
