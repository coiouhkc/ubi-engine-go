package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

const EMPTY = " "
const MISS = "."
const HIT = "x"
const SHIP = "o"
const FOG = "?"

const SIZE_X = 10
const SIZE_Y = 10
const SIZE = SIZE_X * SIZE_Y

const CMD_INFO = "info"
const CMD_QUIT = "quit"
const CMD_SHOOT = "shoot"
const CMD_STATE = "state"

const INFO_HIT = "hit"
const INFO_MISS = "miss"
const INFO_SUNK = "sunk"

// TODO: transform to module containing state, separate command handling (CLI) from game/engine logic

func main() {

	// TODO: replace with a random generated field
	var BOARD_OWN string = strings.Replace(`
	oooo......
	..........
	ooo.oo.oo.
	..........
	ooo.oo.oo.
	..........
	.o.o.o.o..
	..........
	..........
	..........
	`, "\n", "", -1)
	// TODO: use regex replace to avoid multiple lines
	BOARD_OWN = strings.Replace(BOARD_OWN, "\r", "", -1)
	BOARD_OWN = strings.Replace(BOARD_OWN, "\t", "", -1)
	BOARD_OWN = strings.Replace(BOARD_OWN, " ", "", -1)
	BOARD_OWN = strings.Replace(BOARD_OWN, MISS, EMPTY, -1)

	var BOARD_OTHER string = strings.Repeat(FOG, SIZE)

	// TODO: generate the array based on size
	var AXIS_X = [10]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

	// TODO: generate the array based on size
	var AXIS_Y = [10]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")

		command, _ := reader.ReadString('\n')
		command = strings.Replace(command, "\n", "", -1)

		switch {
		case strings.HasPrefix(command, CMD_INFO):
			fmt.Println(`ubi-engine-go, version: 0.1`)
		case strings.HasPrefix(command, CMD_QUIT):
			fmt.Println(command)
			os.Exit(0)
		case strings.HasPrefix(command, CMD_STATE):
			fmt.Println("position ours " + BOARD_OWN)
			fmt.Println("position theirs " + BOARD_OTHER)
		case strings.HasPrefix(command, CMD_SHOOT):
			// TODO: add command validation
			commandWithOpts := strings.Split(command, " ")
			x, y := SplitShortAlgebraic(commandWithOpts[1])

			idx := sort.Search(len(AXIS_X), func(i int) bool {
				return AXIS_X[i] >= x
			})

			idy := sort.Search(len(AXIS_Y), func(i int) bool {
				return AXIS_Y[i] >= y
			})

			chars := []rune(BOARD_OWN)
			index := idx*SIZE_X + idy
			value := string(chars[index])
			if value == SHIP {
				// TODO: add detection of sunk ships
				fmt.Println(INFO_HIT)
				BOARD_OWN = ReplaceCharAt(BOARD_OWN, index, HIT)
			} else {
				fmt.Println(INFO_MISS)
				BOARD_OWN = ReplaceCharAt(BOARD_OWN, index, MISS)
				// TODO: add engine turn(s)
			}
		// TODO: implement further commands from spec
		default:
			fmt.Println("Invalid input")
		}
	}
}

func SplitShortAlgebraic(s string) (string, string) {
	reAlgebraic := regexp.MustCompile(`([a-z]+)(\d+)`)
	segs := reAlgebraic.FindAllStringSubmatch(s, -1)
	return segs[0][1], segs[0][2]
}

func ReplaceCharAt(s string, i int, v string) string {
	chars := []rune(s)
	chars[i] = ([]rune(v))[0]
	return string(chars)
}
