package main

import (
	"fmt"
	"sort"

	"github.com/kalexmills/aoc2020/bikeshed"
)

func main() {
	input := bikeshed.Read(10)
	defer input.Close()
	ints := bikeshed.ParseIntList(input)
	fmt.Println(solve1(ints))
}

/*func solve2(ints []int) int {
	sort.Ints(ints)
	check := append(append([]int{0}, ints...), ints[len(ints)-1]+3)
	result := 1
	for i := 0; i < len(check)-1; i++ {
		mul := 0
		for j := max(0, i-3); j < i; j++ {
			if check[i]-check[j] <= 3 {
				fmt.Println(check[j], check[i], result)
				mul++
			}
		}
		if mul > 1 {
			result *= 2
		}
	}
	return result
}*/

func solve1(ints []int) int {
	sort.Ints(ints)
	check := append(append([]int{0}, ints...), ints[len(ints)-1]+3)
	oneDiff, threeDiff := 0, 0
	for i := 0; i < len(check)-1; i++ {
		if check[i+1]-check[i] == 1 {
			oneDiff++
		}
		if check[i+1]-check[i] == 3 {
			threeDiff++
		}
	}
	return oneDiff * threeDiff
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
