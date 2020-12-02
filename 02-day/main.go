package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/kalexmills/aoc2020/bikeshed"
)

func main() {
	input := bikeshed.Read(2)
	defer input.Close()
	lines, err := parseLines(input)
	if err != nil {
		log.Fatalf("couldn't parse input: %v", err)
	}
	fmt.Println(solve(lines, func(p Policy, password string) bool {
		return p.IsActuallyValid(password)
	}))
}

func solve(lines []Line, valid func(Policy, string) bool) int {
	validCount := 0
	for _, line := range lines {
		if valid(line.Policy, line.Password) {
			validCount++
		}
	}
	return validCount
}

type Policy struct {
	Char                           rune
	MinOccurrences, MaxOccurrences int
}

func (p Policy) IsActuallyValid(password string) bool {
	idx1 := p.MinOccurrences - 1
	idx2 := p.MaxOccurrences - 1
	idx1Contains := idx1 < len(password) && rune(password[idx1]) == p.Char
	idx2Contains := idx2 < len(password) && rune(password[idx2]) == p.Char
	return (idx1Contains || idx2Contains) && !(idx1Contains && idx2Contains)
}

func (p Policy) IsValid(password string) bool {
	counts := 0
	for _, c := range password {
		if c == p.Char {
			counts++
		}
	}
	return p.MinOccurrences <= counts && counts <= p.MaxOccurrences
}

type Line struct {
	Policy   Policy
	Password string
}

func parseLines(reader io.Reader) ([]Line, error) {
	var result []Line
	scan := bufio.NewScanner(reader)
	for scan.Scan() {
		tokens := strings.Split(scan.Text(), ": ")
		if len(tokens) != 2 {
			return nil, fmt.Errorf("bad format for line '%s'", scan.Text())
		}
		policy, err := parsePolicy(tokens[0])
		if err != nil {
			return nil, err
		}
		result = append(result, Line{
			Policy:   policy,
			Password: tokens[1],
		})
	}
	return result, nil
}

func parsePolicy(str string) (Policy, error) {
	tokens := strings.Split(str, " ")
	if len(tokens) != 2 {
		return Policy{}, fmt.Errorf("bad format for policy: '%s'", str)
	}
	occurrences := strings.Split(tokens[0], "-")
	if len(tokens) != 2 {
		return Policy{}, fmt.Errorf("bad occurrences format: '%s'", tokens[0])
	}
	min, err := strconv.ParseInt(occurrences[0], 10, 32)
	if err != nil {
		return Policy{}, err
	}
	max, err := strconv.ParseInt(occurrences[1], 10, 32)
	if err != nil {
		return Policy{}, err
	}
	if min > max || max < 0 || min < 0 {
		return Policy{}, fmt.Errorf("occurrences unexpected: min = %d, max = %d", min, max)
	}
	return Policy{
		Char:           rune(tokens[1][0]),
		MinOccurrences: int(min),
		MaxOccurrences: int(max),
	}, nil
}
