package main

import (
	"errors"
	"fmt"
	"math"

	"github.com/kalexmills/aoc2020/bikeshed"
)

func main() {
	input := bikeshed.Read(9)
	defer input.Close()
	run := bikeshed.ParseIntList(input)
	target, err := solve1(run)
	if err != nil {
		panic(err)
	}
	fmt.Println(solve2(target, run))
}

func solve2(target int, ints []int) int {
	// puzzle input is only length 1000; brute-force suffices
	for i := 0; i < len(ints); i++ {
		for j := 0; j < len(ints); j++ {
			sum := 0
			max, min := math.MinInt32, math.MaxInt32
			for k := i; k <= j; k++ {
				sum += ints[k]
				if max < ints[k] {
					max = ints[k]
				}
				if ints[k] < min {
					min = ints[k]
				}
			}
			if sum == target {
				return max + min
			}
		}
	}
	panic("no contiguous range found!")
}

func solve1(input []int) (int, error) {
	for i := 25; i < len(input); i++ {
		if !validRun(input[i-25 : i+1]) {
			return input[i], nil
		}
	}
	return 0, errors.New("no valid number found")
}

func validRun(ints []int) bool {
	target := ints[len(ints)-1]
	sumMap := make(map[int]struct{}, len(ints))
	for i := 0; i < len(ints)-1; i++ {
		sumMap[target-ints[i]] = struct{}{}
		if _, ok := sumMap[ints[i]]; ok {
			return true
		}
	}
	return false
}
