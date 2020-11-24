package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, Q int64
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	fmt.Sscanf(scanner.Text(), "%d %d", &N, &Q)
	lastReset := 0
	wealth := make(map[int]int, N)

	for ; Q > 0; Q-- {
		scanner.Scan()
		b := scanner.Bytes()
		if b[0] == 'S' {
			n, m := toIntSet(b[4:])
			wealth[n] = m
		} else if b[0] == 'P' {
			i := toInt(b[6:])
			if v, ok := wealth[i]; ok {
				fmt.Println(v)
			} else {
				fmt.Println(lastReset)
			}
		} else {
			wealth = make(map[int]int, N)
			lastReset = toInt(b[8:])
		}
	}
}

func toIntSet(buf []byte) (n, m int) {
	first := true
	for _, v := range buf {
		if v == ' ' {
			first = false
		} else {
			if first {
				n = n*10 + int(v-'0')
			} else {
				m = m*10 + int(v-'0')
			}
		}
	}
	return
}

func toInt(buf []byte) (n int) {
	for _, v := range buf {
		n = n*10 + int(v-'0')
	}
	return
}
