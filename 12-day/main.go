package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"

	"github.com/kalexmills/aoc2020/bikeshed"
)

func main() {
	input := bikeshed.Read(12)
	fmt.Println(solve2(parseInput(input)))
}

func solve1(moves []Move) int {
	facing := 0
	x, y := 0, 0
	for _, move := range moves {
		switch move.op {
		case 'N':
			fallthrough
		case 'S':
			fallthrough
		case 'E':
			fallthrough
		case 'W':
			x, y = translate(x, y, move)
		case 'L':
			facing = (facing + move.arg) % 360
		case 'R':
			facing = (360 + facing - move.arg) % 360
		case 'F':
			switch facing {
			case 0:
				x, y = translate(x, y, Move{'E', move.arg})
			case 90:
				x, y = translate(x, y, Move{'N', move.arg})
			case 180:
				x, y = translate(x, y, Move{'W', move.arg})
			case 270:
				x, y = translate(x, y, Move{'S', move.arg})
			default:
				log.Fatalf("%d was not a valid facing", facing)
			}
		}
	}
	return abs(x) + abs(y)
}

func solve2(moves []Move) int {
	rotate := func(x, y int, ang int) (int, int) {
		switch ang {
		case 0:
			return x, y
		case 90:
			return -y, x
		case 180:
			return -x, -y
		case 270:
			return y, -x
		default:
			log.Fatalf("can only do orthogonal 2D rotations, received angle of: %d", ang)
		}
		return -1, -1
	}
	wx, wy := 10, 1
	x, y := 0, 0
	for _, move := range moves {
		switch move.op {
		case 'N':
			fallthrough
		case 'S':
			fallthrough
		case 'E':
			fallthrough
		case 'W':
			wx, wy = translate(wx, wy, move)
		case 'L':
			wx, wy = rotate(wx, wy, move.arg)
		case 'R':
			wx, wy = rotate(wx, wy, (360-move.arg)%360)
		case 'F':
			x, y = x+(move.arg*wx), y+(move.arg*wy)
		}
	}
	return abs(x) + abs(y)
}

func translate(x, y int, move Move) (int, int) {
	switch move.op {
	case 'N':
		y += move.arg
	case 'S':
		y -= move.arg
	case 'E':
		x += move.arg
	case 'W':
		x -= move.arg
	}
	return x, y
}

type Move struct {
	op  byte
	arg int
}

func parseInput(reader io.Reader) []Move {
	scanner := bufio.NewScanner(reader)
	var result []Move
	for scanner.Scan() {
		str := scanner.Text()
		arg, err := strconv.Atoi(str[1:])
		if err != nil {
			log.Fatalf("could not parse line containing %s", str)
		}
		result = append(result, Move{
			op:  str[0],
			arg: arg,
		})
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
