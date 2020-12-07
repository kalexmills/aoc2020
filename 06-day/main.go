package main

import (
	"bufio"
	"fmt"
	"io"

	"github.com/kalexmills/aoc2020/bikeshed"
)

func main() {
	input := bikeshed.Read(6)

	fmt.Println(solve2(parseGroups(input)))
}

func solve1(groups []GroupAnswers) int {
	sum := 0
	for _, group := range groups {
		sum += len(group.counts)
	}
	return sum
}

func solve2(groups []GroupAnswers) int {
	sum := 0
	for _, group := range groups {
		groupCount := 0
		for _, count := range group.counts {
			if count == group.size {
				groupCount++
			}
		}
		sum += groupCount
	}
	return sum
}

type GroupAnswers struct {
	counts map[rune]int
	size   int
}

func NewGroupAnswers() GroupAnswers {
	return GroupAnswers{
		counts: make(map[rune]int),
	}
}

func parseGroups(input io.Reader) []GroupAnswers {
	var result []GroupAnswers
	curr := NewGroupAnswers()
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		if scanner.Text() == "" {
			result = append(result, curr)
			curr = NewGroupAnswers()
			continue
		}
		for _, answer := range scanner.Text() {
			curr.counts[answer]++
		}
		curr.size++
	}
	result = append(result, curr)
	return result
}
