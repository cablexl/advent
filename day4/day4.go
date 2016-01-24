package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func main() {
	input := []byte("ckczppom")
	var count int
	for ; ; count++ {
		val := strconv.Itoa(count)
		sum := md5.Sum(append(input, []byte(val)...))
		ss := sum[:3]
		if ss[0] == 0x00 && ss[1] == 0x00 && ss[2] == 0x00 {
			fmt.Printf("%x %d %d\n", sum, count, sum[:3])
			break
		}
	}
}
