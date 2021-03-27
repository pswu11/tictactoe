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

var pi = make([]string, 9)

var refer = map[vector]int{
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
	x, err1 := strconv.Atoi(p[0])
	y, err2 := strconv.Atoi(p[1])

	if err1 != nil || err2 != nil {
		fmt.Println("Invalid position, try again.")
		return giveStep(player)
	}

	var position vector

	position.X = int(x)
	position.Y = int(y)

	_, ok := refer[position]
	if ok {
		if pi[refer[position]] != " " {

			fmt.Println("This position is taken!")
			return giveStep(player)
		}
	}
	return position
}
