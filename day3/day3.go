package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	input, err := ioutil.ReadFile("day3.input.txt")

	if err != nil {
		os.Exit(1)
	}

	var m = make(map[string]int)
	var x, y, rx, ry int

	m[fmt.Sprintf("x=%d,y=%d", x, y)] = 1

	for i := 0; i < len(input); i += 2 {
		direction := input[i]
		robo := input[i+1]
		switch direction {
		case 0x3c: // left
			x -= 1
			pos := fmt.Sprintf("x=%d,y=%d", x, y)
			m[pos] = m[pos] + 1
		case 0x3e: // right
			x += 1
			pos := fmt.Sprintf("x=%d,y=%d", x, y)
			m[pos] = m[pos] + 1
		case 0x5e: // up
			y += 1
			pos := fmt.Sprintf("x=%d,y=%d", x, y)
			m[pos] = m[pos] + 1
		case 0x76: // down
			y -= 1
			pos := fmt.Sprintf("x=%d,y=%d", x, y)
			m[pos] = m[pos] + 1
		}

		switch robo {
		case 0x3c: // left
			rx -= 1
			pos := fmt.Sprintf("x=%d,y=%d", rx, ry)
			m[pos] = m[pos] + 1
		case 0x3e: // right
			rx += 1
			pos := fmt.Sprintf("x=%d,y=%d", rx, ry)
			m[pos] = m[pos] + 1
		case 0x5e: // up
			ry += 1
			pos := fmt.Sprintf("x=%d,y=%d", rx, ry)
			m[pos] = m[pos] + 1
		case 0x76: // down
			ry -= 1
			pos := fmt.Sprintf("x=%d,y=%d", rx, ry)
			m[pos] = m[pos] + 1
		}
	}
	for key, value := range m {
		fmt.Println("Key:", key, "Value:", value)
	}

	fmt.Println(len(m))

}
