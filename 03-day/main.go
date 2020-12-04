package main

import (
	"bufio"
	"fmt"
	"io"

	"github.com/kalexmills/aoc2020/bikeshed"
)

func main() {
	input := bikeshed.Read(3)
	defer input.Close()
	trees := ParseTreeGrid(input)
	fmt.Println(solve2(trees, []Point2D{{3, 1}, {1, 1}, {5, 1}, {7, 1}, {1, 2}}))
}

func solve2(trees TreeGrid, slopes []Point2D) int {
	result := 1
	for _, slope := range slopes {
		result *= solve1(trees, slope)
	}
	return result
}

func solve1(trees TreeGrid, slope Point2D) int {
	count := 0
	x := slope.x
	for y := slope.y; y < trees.height; y += slope.y {
		if trees.HasTree(x, y) {
			count++
		}
		x += slope.x
	}
	return count
}

type Point2D struct {
	x, y int
}

type TreeGrid struct {
	width, height int
	trees         map[Point2D]struct{}
}

func (tg TreeGrid) HasTree(x, y int) bool {
	if !tg.InBounds(x, y) {
		return false
	}
	_, ok := tg.trees[Point2D{x % tg.width, y}]
	return ok
}

func (tg TreeGrid) InBounds(x, y int) bool {
	return x >= 0 && y >= 0 && y < tg.height
}

func ParseTreeGrid(reader io.Reader) TreeGrid {
	result := TreeGrid{
		trees: make(map[Point2D]struct{}),
	}
	scan := bufio.NewScanner(reader)
	y := 0
	for scan.Scan() {
		line := scan.Text()
		for x, ch := range line {
			if ch == '#' {
				result.trees[Point2D{x, y}] = struct{}{}
			}
			result.width = max(result.width, x+1)
		}
		result.height = max(result.height, y+1)
		y++
	}
	return result
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
