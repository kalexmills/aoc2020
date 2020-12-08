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

type Opcode byte

const (
	NOP Opcode = iota
	ACC
	JMP
)

func parseOpcode(str string) (Opcode, error) {
	switch str {
	case "nop":
		return NOP, nil
	case "acc":
		return ACC, nil
	case "jmp":
		return JMP, nil
	}
	return NOP, fmt.Errorf("could not parse opcode: %s", str)
}

type Instruction struct {
	opcode Opcode
	arg    int
}

func parseInput(input io.Reader) []Instruction {
	scanner := bufio.NewScanner(input)
	var result []Instruction
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		if len(tokens) != 2 {
			log.Fatalf("expected 2 tokens in line %s", scanner.Text())
		}
		offset, err := strconv.ParseInt(tokens[1], 10, 32)
		if err != nil {
			log.Fatalf("could not parse field 1 as integer, field = %s: %v", tokens[1], err)
		}
		opcode, err := parseOpcode(tokens[0])
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, Instruction{
			opcode: opcode,
			arg:    int(offset),
		})
	}
	return result
}

type StatusCode byte

const (
	Complete StatusCode = iota
	InfiniteLoop
)

func solve1(program []Instruction) (int, StatusCode) {
	visited := make(map[int]struct{}) // visited line numbers
	accum := 0
	pc := 0
	for {
		if pc >= len(program) {
			return accum, Complete
		}
		// check if instruction was visited already
		if _, ok := visited[pc]; ok {
			return accum, InfiniteLoop
		}
		visited[pc] = struct{}{}
		// execute instruction
		switch program[pc].opcode {
		case ACC:
			accum += program[pc].arg
			pc++
		case JMP:
			pc += program[pc].arg
		case NOP:
			pc++
		}
	}
}

func solve2(program []Instruction) int {
	for idx, inst := range program {
		switch inst.opcode {
		case NOP:
			program[idx] = Instruction{JMP, inst.arg}
			if result, status := solve1(program); status == Complete {
				return result
			}
			program[idx] = inst
		case JMP:
			program[idx] = Instruction{NOP, inst.arg}
			if result, status := solve1(program); status == Complete {
				return result
			}
			program[idx] = inst
		}
	}
	log.Fatalln("the spec is a lie!")
	return 0
}

func main() {
	input := bikeshed.Read(8)
	defer input.Close()
	fmt.Println(solve2(parseInput(input)))
}
