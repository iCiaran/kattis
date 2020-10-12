package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

func (p *point) String() string {
	return fmt.Sprintf("[%d, %d]", p.x, p.y)
}

var Empty struct{}

func NewPoint(x, y int) *point {
	return &point{x, y}
}

func (p *point) sameCoord(q *point) bool {
	return p.x == q.x && p.y == q.y
}

func check(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	check(scanner.Err())

	split := strings.Fields(scanner.Text())
	height, err := strconv.Atoi(split[0])
	check(err)
	width, err := strconv.Atoi(split[1])
	check(err)

	grid := make([][]int, height)

	for y := 0; y < height; y++ {
		grid[y] = make([]int, width)

		scanner.Scan()
		check(scanner.Err())

		for x := 0; x < width; x++ {
			v, err := strconv.Atoi(string(scanner.Text()[x]))
			check(err)
			grid[y][x] = v
		}
	}

	scanner.Scan()
	check(scanner.Err())

	nLines, err := strconv.Atoi(scanner.Text())
	check(err)

	visited := make(map[point]struct{}, 0)
	regions := make([]map[point]struct{}, 0)

	directions := &[]*point{NewPoint(-1, 0), NewPoint(1, 0), NewPoint(0, -1), NewPoint(0, 1)}

	for i := 0; i < nLines; i++ {
		scanner.Scan()
		check(scanner.Err())
		split = strings.Fields(scanner.Text())

		x1, err := strconv.Atoi(split[0])
		check(err)
		y1, err := strconv.Atoi(split[1])
		check(err)

		a := NewPoint(y1-1, x1-1)

		x2, err := strconv.Atoi(split[2])
		check(err)
		y2, err := strconv.Atoi(split[3])
		check(err)

		b := NewPoint(y2-1, x2-1)

		regionA := getRegion(a, visited, regions)
		if regionA == -1 {
			regionA = searchRegion(a, visited, &regions, grid, *directions)
		}

		regionB := getRegion(b, visited, regions)
		if regionB == -1 {
			regionB = searchRegion(b, visited, &regions, grid, *directions)
		}

		if regionA == regionB {
			if grid[a.y][a.x] == 1 {
				fmt.Println("decimal")
			} else {
				fmt.Println("binary")
			}
		} else {
			fmt.Println("neither")
		}
	}
}

func getRegion(p *point, v map[point]struct{}, r []map[point]struct{}) int {
	if _, ok := v[*p]; !ok {
		return -1
	}

	for i := 0; i < len(r); i++ {
		if _, ok := r[i][*p]; ok {
			return i
		}
	}

	return -1
}

func searchRegion(p *point, v map[point]struct{}, r *[]map[point]struct{}, g [][]int, d []*point) int {
	*r = append(*r, make(map[point]struct{}))
	i := len(*r) - 1
	t := g[p.y][p.x]
	queue := make(chan *point, 1000000)
	queue <- p
	for len(queue) > 0 {
		curr := <-queue
		v[*curr] = Empty
		(*r)[i][*curr] = Empty
		for _, dir := range d {
			nx := curr.x + dir.x
			if nx < 0 || nx >= len(g[0]) {
				continue
			}
			ny := curr.y + dir.y
			if ny < 0 || ny >= len(g) {
				continue
			}
			if g[ny][nx] != t {
				continue
			}
			nextPoint := NewPoint(nx, ny)
			if _, ok := v[*nextPoint]; !ok {
				queue <- nextPoint
			}
		}
	}
	return i
}
