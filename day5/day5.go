package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	var vowel = regexp.MustCompile(`[aeiou]`)

	f, err := os.Open("day5.input.txt")

	fmt.Println(len(vowel.FindAllString("adambei", -1)))

	if err != nil {
		os.Exit(1)
	}

	input := bufio.NewScanner(f)
	nicecount := 0
	for input.Scan() {
		numNoOverlapRepeats(input.Text())
	}

	fmt.Println(nicecount)
}

func numNoOverlapRepeats(str string) (num int) {
	repeats := make(map[string]int)
	for i := 0; i < len(str); i++ {
		if i+1 == len(str) {
			continue
		}

		if str[i] == str[i+1] {
			repeats[str[i:i+2]]++
			i++
		}
	}
	fmt.Printf("%s %s\n", str, repeats)
	return
}

func checkRepeats(str string) (repeats bool) {
	var prev rune
	for _, curr := range str {
		if curr == prev {
			repeats = true
			return
		}
		prev = curr
	}
	return
}
