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

type line []vector

var board = map[vector]string{
	{1, 3}: "",
	{2, 3}: "",
	{3, 3}: "",
	{1, 2}: "",
	{2, 2}: "",
	{3, 2}: "",
	{1, 1}: "",
	{2, 1}: "",
	{3, 1}: "",
}

var winningStrategy = map[int]line{
	1: {{1, 1}, {1, 2}, {1, 3}},
	2: {{2, 1}, {2, 2}, {2, 3}},
	3: {{3, 1}, {3, 2}, {3, 3}},
	4: {{1, 1}, {2, 1}, {3, 1}},
	5: {{1, 2}, {2, 2}, {3, 2}},
	6: {{1, 3}, {2, 3}, {3, 3}},
	7: {{1, 1}, {2, 2}, {3, 3}},
	8: {{1, 3}, {2, 2}, {3, 1}},
}

func giveStep(player int) vector {

	fmt.Println("Player", player, "'s move:")
	buf := bufio.NewReader(os.Stdin)
	pos, err := buf.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	p := []string(strings.Split(strings.TrimSpace(pos), ""))

	// validate input length == 2
	if len(p) != 2 {
		fmt.Println("Invalid position, try again.")
		return giveStep(player)
	}

	x, err1 := strconv.Atoi(p[0])
	y, err2 := strconv.Atoi(p[1])

	// validate input == integers
	if err1 != nil || err2 != nil {
		fmt.Println("Invalid position, try again.")
		return giveStep(player)
	}

	var position vector

	position.X = int(x)
	position.Y = int(y)

	// validate input is available
	_, ok := board[position]
	if ok {
		if board[position] != " " {

			fmt.Println("This position is taken!")
			return giveStep(player)
		}
	}
	return position
}

func updateBoard(position vector, step string) {

	board[position] = step
	printCurrentBoard()

}

func newBoard() {
	for i := range board {
		board[i] = " "
	}
	printCurrentBoard()
}

func printCurrentBoard() {

	ln := "   -------------\n"
	head := "(Y)                \n"
	foot := "     1   2   3  (X)\n"
	bodyThree := fmt.Sprintf(" 3 | %v | %v | %v |\n", board[vector{1, 3}], board[vector{2, 3}], board[vector{3, 3}])
	bodyTwo := fmt.Sprintf(" 2 | %v | %v | %v |\n", board[vector{1, 2}], board[vector{2, 2}], board[vector{3, 2}])
	bodyOne := fmt.Sprintf(" 1 | %v | %v | %v |\n", board[vector{1, 1}], board[vector{2, 1}], board[vector{3, 1}])

	result := head + ln + bodyThree + ln + bodyTwo + ln + bodyOne + ln + foot

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
