package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	parseInput()
}

var wire = make(map[string]string)

func parseInput() {
	f, err := os.Open("day7.input.txt")

	if err != nil {
		os.Exit(1)
	}

	input := bufio.NewScanner(f)
	for input.Scan() {
		split := strings.Split(input.Text(), " -> ")
		wire[split[1]] = split[0]
	}

	fmt.Println(solve("a", wire["a"]))
}

func solve(reg, op string) (result uint16) {
	fmt.Printf("running solve for %s with op %s\n", reg, op)
	split := strings.Split(op, " ")
	switch len(split) {
	case 1:
		if val, err := strconv.Atoi(split[0]); err == nil {
			result = uint16(val)
		} else {
			result = solve(split[0], wire[split[0]])
		}
	case 2:
		result = ^solve(split[1], wire[split[1]])
	case 3:
		switch split[1] {
		case "OR":
			if val, err := strconv.Atoi(split[0]); err != nil {
				result = solve(split[0], wire[split[0]]) | solve(split[2], wire[split[2]])
			} else {
				result = uint16(val) | solve(split[2], wire[split[2]])
			}
		case "AND":
			if val, err := strconv.Atoi(split[0]); err != nil {
				result = solve(split[0], wire[split[0]]) & solve(split[2], wire[split[2]])
			} else {
				result = uint16(val) & solve(split[2], wire[split[2]])
			}
		case "LSHIFT":
			val, _ := strconv.Atoi(split[2])
			result = solve(split[0], wire[split[0]]) << uint16(val)
		case "RSHIFT":
			val, _ := strconv.Atoi(split[2])
			result = solve(split[0], wire[split[0]]) >> uint16(val)
		}
	}
	wire[reg] = fmt.Sprintf("%d", result)
	return
}
