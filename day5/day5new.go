package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("day5.input.txt")
	if err != nil {
		os.Exit(1)
	}

	input := bufio.NewScanner(f)
	for input.Scan() {
		repeats(input.Text())
	}
}

func repeats(str string) {
	reps := make(map[string]int)

	for idx, _ := range str {
		reps[str[idx:idx+1]]++

	}

	for i := 0; i < len(str)-3; i++ {
	}
	fmt.Printf("%s %s\n", str, reps)
}
