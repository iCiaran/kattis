package main

import (
	"bufio"
	"fmt"
	"os"
)

type pair struct {
	value   int
	lastSet int
}

func main() {
	var N, Q int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	fmt.Sscanf(scanner.Text(), "%d %d", &N, &Q)
	lastResetAmount := 0
	lastResetLine := -1
	wealthValues := make([]int, N+1)
	wealthLast := make([]int, N+1)

	for i := 1; i <= Q; i++ {
		scanner.Scan()
		b := scanner.Bytes()
		if b[0] == 'S' {
			n, m := toIntSet(b[4:])
			wealthValues[n] = m
			wealthLast[n] = i
		} else if b[0] == 'P' {
			person := toInt(b[6:])
			if wealthLast[person] > lastResetLine {
				fmt.Println(wealthValues[person])
			} else {
				fmt.Println(lastResetAmount)
			}
		} else {
			lastResetAmount = toInt(b[8:])
			lastResetLine = i
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
