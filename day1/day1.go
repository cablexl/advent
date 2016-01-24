package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	input, err := ioutil.ReadFile("day1.input.txt")

	if err != nil {
		os.Exit(1)
	}

	output := 0
	for idx, x := range input {
		if x == 0x28 {
			output += 1
		} else if x == 0x29 {
			output -= 1
		}
		if output < 0 {
			fmt.Println(idx + 1)
			break
		}
	}
	fmt.Println(output)
}
