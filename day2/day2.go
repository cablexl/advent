package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Dimensions struct {
	L, W, H int
}

func main() {
	f, err := os.Open("day2.input.txt")

	if err != nil {
		os.Exit(1)
	}

	input := bufio.NewScanner(f)
	var tp, tr int
	for input.Scan() {
		dim := splitDimension(input.Text())
		total, slack, bow, ribbon := calcPaper(dim)

		fmt.Printf("length=%d,width=%d,height=%d,total=%d,slack=%d\n", dim.L, dim.W, dim.H, total, slack)
		tp = tp + total + slack
		tr = tr + bow + ribbon

	}
	fmt.Println(tp)
	fmt.Println(tr)
}

func splitDimension(input string) Dimensions {
	split := strings.Split(input, "x")
	l, _ := strconv.Atoi(split[0])
	w, _ := strconv.Atoi(split[1])
	h, _ := strconv.Atoi(split[2])
	return Dimensions{l, w, h}
}

func calcPaper(dim Dimensions) (total, slack, bow, ribbon int) {
	arr := []int{dim.L, dim.W, dim.H}

	sort.Ints(arr)
	lw := dim.L * dim.W
	wh := dim.W * dim.H
	hl := dim.H * dim.L
	arr2 := []int{lw, wh, hl}
	sort.Ints(arr2)
	fmt.Printf("lw=%d,wh=%d,hl=%d\n", lw, wh, hl)

	total = 2*lw + 2*wh + 2*hl

	bow = arr[0] * arr[1] * arr[2]
	ribbon = arr[0] + arr[0] + arr[1] + arr[1]

	slack = arr2[0]

	return
}
