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
	input := bikeshed.Read(7)
	defer input.Close()
	fmt.Println(parseInput(input).solve2("shiny gold"))
}

func solve1(graph Graph) int {
	reversed := reverse(graph)
	visited := make(map[string]struct{})
	frontier := []string{"shiny gold"}
	for len(frontier) > 0 {
		curr := frontier[0]
		frontier = frontier[1:]
		visited[curr] = struct{}{}
		for neighbor := range reversed[curr] {
			if _, ok := visited[neighbor]; !ok {
				frontier = append(frontier, neighbor)
			}
		}
	}
	return len(visited) - 1
}

type Graph map[string]map[string]int

func (g Graph) solve2(source string) int {
	if len(g[source]) == 0 {
		return 0
	}
	sum := 0
	for target, weight := range g[source] {
		sum += weight + weight*g.solve2(target)
	}
	return sum
}

func reverse(g Graph) Graph {
	result := make(Graph)
	for source, adjacent := range g {
		for target, weight := range adjacent {
			if _, ok := result[target]; !ok {
				result[target] = make(map[string]int)
			}
			result[target][source] = weight
		}
	}
	return result
}

func parseInput(reader io.Reader) Graph {
	scanner := bufio.NewScanner(reader)
	result := make(Graph)
	for scanner.Scan() {
		source, adjacent := parseLine(scanner.Text())
		if _, ok := result[source]; ok {
			log.Fatalf("found multiple rules for %s", source)
		}
		result[source] = adjacent
	}
	return result
}

func parseLine(str string) (string, map[string]int) {
	tokens := strings.Split(str, " bags contain ")
	if len(tokens) != 2 {
		log.Fatalf("expected two tokens around ' bags contain ', found %s", str)
	}
	source := tokens[0]
	adjacent := make(map[string]int)
	list := strings.Split(tokens[1], ", ")
	for _, edge := range list {
		cleaned := bagless(edge)
		if cleaned == "no other" {
			break
		}
		tokens := strings.Split(cleaned, " ")
		if len(tokens) != 3 {
			log.Fatalf("expected 3 tokens separated by spaces, found %s", cleaned)
		}
		weight, err := strconv.ParseInt(tokens[0], 10, 32)
		if err != nil {
			panic(err)
		}
		target := tokens[1] + " " + tokens[2]
		adjacent[target] = int(weight)
	}
	return source, adjacent
}

func bagless(str string) string {
	for _, bag := range []string{" bag.", " bags.", " bags", " bag"} {
		str = strings.Replace(str, bag, "", 1)
	}
	return str
}
