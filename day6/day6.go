package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type LightGrid [1000][1000]int

type GridAction struct {
	Action, X1, Y1, X2, Y2 int
}

const (
	ON = iota
	OFF
	TOGGLE
)

func main() {
	f, err := os.Open("day6.input.txt")

	if err != nil {
		os.Exit(1)
	}

	input := bufio.NewScanner(f)
	var lg LightGrid
	for input.Scan() {
		ga := processLine(input.Text())
		fmt.Println(ga)
		applyGridAction(ga, &lg)
	}
	fmt.Println(countLidLights(&lg))
}

func countLidLights(lg *LightGrid) (count int) {
	for i := 0; i <= 999; i++ {
		for j := 0; j <= 999; j++ {
			if lg[i][j] > 0 {
				count += lg[i][j]
			}
		}
	}
	return
}

func applyGridAction(ga *GridAction, lg *LightGrid) {
	for x := ga.X1; x <= ga.X2; x++ {
		for y := ga.Y1; y <= ga.Y2; y++ {
			switch ga.Action {
			case ON:
				lg[x][y]++
			case OFF:
				if lg[x][y] > 0 {
					lg[x][y]--
				}
			case TOGGLE:
				lg[x][y] += 2
			}
		}
	}
}

func processLine(str string) *GridAction {
	ga := GridAction{}
	if strings.Contains(str, "turn on") {
		ga.Action = ON
		split := strings.Split(str, " ")
		point := strings.Split(split[2], ",")
		ga.X1, _ = strconv.Atoi(point[0])
		ga.Y1, _ = strconv.Atoi(point[1])
		point = strings.Split(split[4], ",")
		ga.X2, _ = strconv.Atoi(point[0])
		ga.Y2, _ = strconv.Atoi(point[1])
	} else if strings.Contains(str, "turn off") {
		ga.Action = OFF
		split := strings.Split(str, " ")
		point := strings.Split(split[2], ",")
		ga.X1, _ = strconv.Atoi(point[0])
		ga.Y1, _ = strconv.Atoi(point[1])
		point = strings.Split(split[4], ",")
		ga.X2, _ = strconv.Atoi(point[0])
		ga.Y2, _ = strconv.Atoi(point[1])
	} else { // toggle
		ga.Action = TOGGLE
		split := strings.Split(str, " ")
		point := strings.Split(split[1], ",")
		ga.X1, _ = strconv.Atoi(point[0])
		ga.Y1, _ = strconv.Atoi(point[1])
		point = strings.Split(split[3], ",")
		ga.X2, _ = strconv.Atoi(point[0])
		ga.Y2, _ = strconv.Atoi(point[1])
	}
	return &ga
}
