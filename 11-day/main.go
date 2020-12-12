package main

import (
	"bufio"
	"fmt"
	"io"

	"github.com/kalexmills/aoc2020/bikeshed"
)

func main() {
	input := bikeshed.Read(11)
	defer input.Close()
	fmt.Println(solve1(parseInput(input)))
}

func solve1(grid SeatGrid) int {
	other := NewSeatGrid(grid.width)
	front := &grid // double-buffering
	back := &other
	front.CopyInto(back)
	changeCount := 1
	for changeCount > 0 {
		changeCount = 0
		for x := 0; x < front.Width(); x++ {
			for y := 0; y < front.Height(); y++ {
				if front.IsFloor(x, y) {
					continue
				}
				adjacentOccupants := 0
				front.VisitLineOfSight(x, y, func(x, y int) {
					if front.IsOccupied(x, y) {
						adjacentOccupants++
					}
				})
				if !front.IsOccupied(x, y) && adjacentOccupants == 0 {
					changeCount++
					back.ToggleSeat(x, y)
				}
				if front.IsOccupied(x, y) && adjacentOccupants >= 5 {
					changeCount++
					back.ToggleSeat(x, y)
				}
			}
		}
		fmt.Println(changeCount)
		back.CopyInto(front)
		back, front = front, back
	}
	count := 0
	for _, x := range back.cells {
		if x == Occupied {
			count++
		}
	}
	return count
}

func parseInput(reader io.Reader) SeatGrid {
	scanner := bufio.NewScanner(reader)
	var result *SeatGrid
	for scanner.Scan() {
		line := scanner.Text()
		if result == nil {
			r := NewSeatGrid(len(line))
			result = &r
		}
		for _, cell := range line {
			switch cell {
			case '.':
				result.cells = append(result.cells, Floor)
			case 'L':
				result.cells = append(result.cells, Unoccupied)
			case '#':
				result.cells = append(result.cells, Occupied)
			}
		}
	}
	return *result
}

type CellState uint8

const (
	Floor CellState = iota
	Unoccupied
	Occupied
)

type SeatGrid struct {
	width int
	cells []CellState
}

func NewSeatGrid(width int) SeatGrid {
	return SeatGrid{
		width: width,
	}
}

// VisitLineOfSight calls f for the cell containing the first seat in each of the 8 directions.
func (sg SeatGrid) VisitLineOfSight(x, y int, f func(x, y int)) {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i != 0 || j != 0 {
				k := 1
				xx := x + i
				yy := y + j
				for xx >= 0 && xx < sg.Width() &&
					yy >= 0 && yy < sg.Height() {
					if !sg.IsFloor(xx, yy) {
						f(xx, yy)
						break
					}
					k++
					xx = x + i*k
					yy = y + j*k
				}
			}
		}
	}
}

func (sg SeatGrid) VisitAdjacent(x, y int, f func(x, y int)) {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if (i != 0 || j != 0) &&
				x+i >= 0 && x+i < sg.Width() &&
				y+j >= 0 && y+j < sg.Height() {
				f(x+i, y+j)
			}
		}
	}
}

func (sg SeatGrid) Width() int {
	return sg.width
}
func (sg SeatGrid) Height() int {
	return len(sg.cells) / sg.width
}

func (sg SeatGrid) CopyInto(other *SeatGrid) {
	if len(other.cells) < len(sg.cells) {
		other.cells = make([]CellState, len(sg.cells))
	}
	other.width = sg.width
	copy(other.cells, sg.cells)
}

func (sg SeatGrid) IsFloor(x, y int) bool {
	return sg.cells[sg.width*y+x] == Floor
}

func (sg SeatGrid) IsOccupied(x, y int) bool {
	return sg.cells[sg.width*y+x] == Occupied
}

func (sg *SeatGrid) ToggleSeat(x, y int) {
	switch sg.cells[sg.width*y+x] {
	case Occupied:
		sg.cells[sg.width*y+x] = Unoccupied
	case Unoccupied:
		sg.cells[sg.width*y+x] = Occupied
	}
}
