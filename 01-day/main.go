package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"strconv"

	"github.com/kalexmills/aoc2020/bikeshed"
)

func main() {
	input, err := bikeshed.Read(1)
	defer input.Close()
	if err != nil {
		log.Fatalf("could not read day 1 input: %v", err)
	}
	result, err := solve2(parseInput(input))
	if err != nil {
		log.Fatalf("could not solve input: %v", err)
	}
	fmt.Printf("answer is %d\n", result)
}

// 299  1041 1654
// 1721  979  366  299  675

type pair struct {
	i, j int
}

func solve2(numbers []int) (int, error) {
	diffMap := make(map[int]pair)
	for i, v1 := range numbers {
		for j, v2 := range numbers {
			diffMap[2020-(v1+v2)] = pair{i, j}
		}
	}
	for _, v3 := range numbers {
		if pair, ok := diffMap[v3]; ok {
			return numbers[pair.i] * numbers[pair.j] * v3, nil
		}
	}
	return 0, errors.New("no answer found")
}

func solve1(numbers []int) (int, error) {
	idxMap := make(map[int]int)
	for idx, i := range numbers {
		if pairIdx, ok := idxMap[2020-i]; ok {
			return i * numbers[pairIdx], nil
		}
		idxMap[i] = idx
	}
	return 0, errors.New("no answer found")
}

func parseInput(input io.Reader) []int {
	scanner := bufio.NewScanner(input)
	var result []int
	for scanner.Scan() {
		i, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			panic("Santa is a lie!")
		}
		result = append(result, int(i))
	}
	return result
}
