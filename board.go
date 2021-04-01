package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
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

	// validate input length == 2 and != 0
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

	// validate if position is available
	if isPositionAvailable(position) == false {
		fmt.Println("Invalid position, try again:")
		return giveStep(player)
	}
	return position

}

func giveRandomStep() vector {

	var position vector

	source1 := rand.NewSource(time.Now().UnixNano())
	source2 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(source1)
	r2 := rand.New(source2)
	position.X = r1.Intn(3) + 1
	position.Y = r2.Intn(3) + 1

	fmt.Println("Computer is taking", position)

	if isPositionAvailable(position) == false {
		fmt.Println("Invalid position, try again.")
		return giveRandomStep()
	}
	return position
}

func isPositionAvailable(position vector) bool {

	_, ok := board[position]
	if ok {
		if board[position] != " " {
			return false
		}
		return true
	}
	return false
}

func updateBoard(position vector, step string, targetboard map[vector]string) {

	targetBoard := targetboard
	targetBoard[position] = step
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
	head := "     1   2   3  (Column)\n"
	foot := "(Row)           \n"
	bodyThree := fmt.Sprintf(" 1 | %v | %v | %v |\n", board[vector{1, 1}], board[vector{1, 2}], board[vector{1, 3}])
	bodyTwo := fmt.Sprintf(" 2 | %v | %v | %v |\n", board[vector{2, 1}], board[vector{2, 2}], board[vector{2, 3}])
	bodyOne := fmt.Sprintf(" 3 | %v | %v | %v |\n", board[vector{3, 1}], board[vector{3, 2}], board[vector{3, 3}])

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
