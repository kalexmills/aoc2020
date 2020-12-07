package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/kalexmills/aoc2020/bikeshed"
)

func main() {
	input := bikeshed.Read(5)
	defer input.Close()
	fmt.Println(solve2(parseInput(input)))
}

func solve2(passes []BoardingPass) (int, error) {
	// missingAbove/missingBelow are sets of IDs that have no neighboring ID above/below
	missingAbove := make(map[int]struct{})
	missingBelow := make(map[int]struct{})
	ids := make(map[int]struct{}) // set of all IDs
	for _, pass := range passes {
		ids[pass.SeatID()] = struct{}{}
	}
	for id := range ids {
		if _, ok := ids[id+1]; !ok {
			missingAbove[id] = struct{}{}
		}
		if _, ok := ids[id-1]; !ok {
			missingBelow[id] = struct{}{}
		}
	}
	// find any empty seat sandwiched between two filled ones.
	for id := range missingAbove {
		if _, ok := missingBelow[id+2]; ok {
			return id + 1, nil
		}
	}
	return -1, errors.New("plane was full")
}

func solve1(passes []BoardingPass) int {
	max := 0
	for _, pass := range passes {
		if max < pass.SeatID() {
			max = pass.SeatID()
		}
	}
	return max
}

type BoardingPass struct {
	Row    int
	Column int
}

func (bp BoardingPass) SeatID() int {
	return bp.Row*8 + bp.Column
}

func parseInput(reader io.Reader) []BoardingPass {
	var result []BoardingPass
	scan := bufio.NewScanner(reader)
	for scan.Scan() {
		if len(scan.Text()) != 10 {
			log.Printf("line did not have 10 characters: %s", scan.Text())
		}
		// lol, no slicing needed since makeBinary ignores uninteresting characters XD
		result = append(result, BoardingPass{
			Row:    int(makeBinary(scan.Text(), 'F', 'B')),
			Column: int(makeBinary(scan.Text(), 'L', 'R')),
		})
	}
	return result
}

func makeBinary(str string, zero, one byte) uint64 {
	result := uint64(0)
	stop := min(len(str)-1, 64)
	for i := 0; i <= stop; i++ {
		if str[i] == zero {
			result <<= 1
		} else if str[i] == one {
			result <<= 1
			result++
		}
	}
	return result
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
