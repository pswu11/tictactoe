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

type line struct {
	dotOne   vector
	dotTwo   vector
	dotThree vector
}

var board = make([]string, 9)

var boardReference = map[vector]int{
	{1, 3}: 0,
	{2, 3}: 1,
	{3, 3}: 2,
	{1, 2}: 3,
	{2, 2}: 4,
	{3, 2}: 5,
	{1, 1}: 6,
	{2, 1}: 7,
	{3, 1}: 8,
}

func giveStep(player int) vector {

	fmt.Println("Player", player, "'s move:")
	buf := bufio.NewReader(os.Stdin)
	pos, err := buf.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	p := []string(strings.Split(strings.TrimSpace(pos), ""))

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

	// validate input == available
	_, ok := boardReference[position]
	if ok {
		if board[boardReference[position]] != " " {

			fmt.Println("This position is taken!")
			return giveStep(player)
		}
	}
	return position
}

func updateBoard(position vector, step string) {

	board[boardReference[position]] = step
	printCurrentBoard()

}

func newBoard() {
	for i := range board {
		if board[i] == "" {
			board[i] = " "
		}
	}
	printCurrentBoard()
}

func printCurrentBoard() {

	ln := "   -------------\n"
	head := "(Y)                \n"
	foot := "     1   2   3  (X)\n"
	bodyThree := fmt.Sprintf(" 3 | %v | %v | %v |\n", board[0], board[1], board[2])
	bodyTwo := fmt.Sprintf(" 2 | %v | %v | %v |\n", board[3], board[4], board[5])
	bodyOne := fmt.Sprintf(" 1 | %v | %v | %v |\n", board[6], board[7], board[8])

	result := head + ln + bodyThree + ln + bodyTwo + ln + bodyOne + ln + foot

	println(result)

}

func playerStep(player int) string {
	if player == 1 {
		return "O"
	}

	return "X"
}

func playerRound(player int) {
	pos := giveStep(player)
	updateBoard(pos, playerStep(player))

}
