package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func gameMode() {

	fmt.Println("Choose the game mode: \n(1) Two Players \n(2) Play against computer")
	buf := bufio.NewReader(os.Stdin)
	mode, err := buf.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	if strings.ToLower(strings.TrimSpace(mode)) == "1" {
		newGameTwoPlayers()
	} else if strings.ToLower(strings.TrimSpace(mode)) == "2" {
		newGamewithComputer()
	} else {
		gameMode()
	}
}

func newGameTwoPlayers() {
	newBoard()
	fmt.Println("Please choose your step by typing the coordinate (row and column), e.g. 13")
	for i := 0; i < 9; i++ {
		if i%2 == 0 {
			playerRound(1, i)
		} else {
			playerRound(2, i)
		}
	}
}

func newGamewithComputer() {
	newBoard()
	fmt.Println("You are player 1")
	fmt.Println("Please choose your step by typing the coordinate (row and column), e.g. 13")
	for i := 0; i < 9; i++ {
		if i%2 == 0 {
			playerRound(1, i)
		} else {
			computerPlayerRound(2, i)
		}
	}

}
