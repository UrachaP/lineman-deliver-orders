package main

import (
	"container/heap"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Point struct {
	row, col int
}

type Node struct {
	point Point
	dist  int
}

type Stack []*Node

func (s Stack) Len() int {
	return len(s)
}

func (s Stack) Less(i, j int) bool {
	return s[i].dist < s[j].dist
}

func (s Stack) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s *Stack) Push(x interface{}) {
	item := x.(*Node)
	*s = append(*s, item)
}

func (s *Stack) Pop() interface{} {
	old := *s
	n := len(old)
	item := old[n-1]
	*s = old[:n-1]
	return item
}

func dijkstra(grid [][]int, rows int, cols int, start, end Point) int {
	dist := make([][]int, rows)
	for i := range dist {
		dist[i] = make([]int, cols)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt32
		}
	}

	dist[start.row][start.col] = grid[start.row][start.col]
	visited := make(map[Point]bool)
	stack := make(Stack, 0)
	heap.Push(&stack, &Node{start, grid[start.row][start.col]})

	for len(stack) > 0 {
		node := heap.Pop(&stack).(*Node)
		currentPoint := node.point
		visited[currentPoint] = true
		//test
		// check go to end point
		if currentPoint == end {
			return dist[currentPoint.row][currentPoint.col]
		}

		// loop for check out of table
		for _, next := range []Point{{currentPoint.row, currentPoint.col + 1}, {currentPoint.row + 1, currentPoint.col}, {currentPoint.row, currentPoint.col - 1}, {currentPoint.row - 1, currentPoint.col}} {
			if next.row >= 0 && next.row < rows && next.col >= 0 && next.col < cols && visited[next] == false {
				// check minimum distance in 4 ways
				newDist := dist[currentPoint.row][currentPoint.col] + grid[next.row][next.col]
				if newDist < dist[next.row][next.col] {
					dist[next.row][next.col] = newDist
					heap.Push(&stack, &Node{next, newDist})
				}
			}
		}
	}

	return -1 // not found
}

func main() {
	for i := 1; i <= 5; i++ {
		startTime := time.Now()
		input, err := os.ReadFile(fmt.Sprintf("in-%d.txt", i))
		if err != nil {
			panic(err)
		}

		rows := strings.Split(string(input), "\n")

		rowsLength := len(rows) - 1
		colsLength := len(strings.Fields(rows[0]))

		grid := make([][]int, rowsLength)

		for i := 0; i < rowsLength; i++ {
			grid[i] = make([]int, colsLength)
			values := strings.Fields(rows[i])
			for j := 0; j < colsLength; j++ {
				grid[i][j], err = strconv.Atoi(values[j])
				if err != nil {
					log.Println(err)
					return
				}
			}
		}

		start, end := Point{0, 0}, Point{rowsLength - 1, colsLength - 1}
		distance := dijkstra(grid, rowsLength, colsLength, start, end)
		fmt.Printf("Exam file name \"in-%d.txt\n", i)
		fmt.Printf("Shortest distance: %d\n", distance)

		endTime := time.Since(startTime)
		fmt.Printf("Execution time: %f Seconds\n\n", endTime.Seconds())
	}

}
