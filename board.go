package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type vector struct {
	X, Y int
}

type line []int

var board = map[int]string{
	1: "1",
	2: "2",
	3: "3",
	4: "4",
	5: "5",
	6: "6",
	7: "7",
	8: "8",
	9: "9",
}

var winningStrategy = map[int]line{
	1: {1, 2, 3},
	2: {4, 5, 6},
	3: {7, 8, 9},
	4: {1, 4, 7},
	5: {2, 5, 8},
	6: {3, 6, 9},
	7: {1, 5, 9},
	8: {3, 5, 7},
}

func giveStep(player int) int {

	fmt.Println("Player", player, "'s move:")
	buf := bufio.NewReader(os.Stdin)
	pos, err := buf.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	p := strings.TrimSpace(pos)

	// validate input length == 1
	if len(p) != 1 {
		fmt.Println("Invalid position, try again.")
		return giveStep(player)
	}

	position, err1 := strconv.Atoi(p)

	// validate input == integers
	if err1 != nil {
		fmt.Println("Invalid position, try again.")
		return giveStep(player)
	}

	// validate input is available
	_, ok := board[position]
	if ok {
		if board[position] == "X" || board[position] == "O" {

			fmt.Println("This position is taken!")
			return giveStep(player)
		}
	}
	return position
}

func updateBoard(position int, step string) {

	board[position] = step
	printCurrentBoard()

}

func newBoard() {
	for i := range board {
		board[i] = strconv.Itoa(i)
	}
	printCurrentBoard()
}

func printCurrentBoard() {

	ln := " -------------\n"
	bodyThree := fmt.Sprintf(" | %v | %v | %v |\n", board[1], board[2], board[3])
	bodyTwo := fmt.Sprintf(" | %v | %v | %v |\n", board[4], board[5], board[6])
	bodyOne := fmt.Sprintf(" | %v | %v | %v |\n", board[7], board[8], board[9])

	result := ln + bodyThree + ln + bodyTwo + ln + bodyOne + ln

	println(result)

}

func checkWinner(player int) bool {
	for _, v := range winningStrategy {
		if board[v[0]] == playerStep(player) && board[v[1]] == playerStep(player) && board[v[2]] == playerStep(player) {
			fmt.Println("Player", player, "won!")
			return true
		}
	}
	return false
}
