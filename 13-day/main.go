package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"

	"github.com/kalexmills/aoc2020/bikeshed"
)

func main() {
	input := bikeshed.Read(13)
	defer input.Close()
	// 	input := strings.NewReader(`939
	// 7,13,x,x,59,x,31,19`)
	fmt.Println(solve2(parseInput(input)))
}

func solve2(_ int, departures []Departure) int {
	// TODO: brute force will not solve this; number theory is needed
	i := 100000000000003
outer:
	for {
		i += departures[0].id
		for _, d := range departures {
			waitTime := (d.id - (i % d.id)) % d.id
			if waitTime != d.idx {
				continue outer
			}
		}
		return i
	}
}

func solve1(timestamp int, times []Departure) int {
	min := math.MaxInt32
	id := -1
	for _, d := range times {
		time := d.id
		waitTime := time - (timestamp % time)
		if min > waitTime {
			min = waitTime
			id = time
		}
	}
	return min * id
}

type Departure struct {
	idx int
	id  int
}

func parseInput(reader io.Reader) (int, []Departure) {
	scanner := bufio.NewScanner(reader)
	if !scanner.Scan() {
		panic("EOF")
	}
	timestamp, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	if !scanner.Scan() {
		panic("EOF")
	}
	timeToks := strings.Split(scanner.Text(), ",")
	var times []Departure
	for idx, token := range timeToks {
		time, err := strconv.Atoi(token)
		if err != nil {
			continue
		}
		//fmt.Printf("%d, %d\n", idx, time)
		times = append(times, Departure{idx, time})
	}
	return timestamp, times
}
