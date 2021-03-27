package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func playerStep(player int) string {
	if player == 1 {
		return "O"
	}

	return "X"
}

func playerRound(player int) {
	pos := giveStep(player)
	updateBoard(pos, playerStep(player))
	if checkWinner(player) {
		fmt.Println("Player", player, "won!")
		playerRestart()
	}

}

func playerRestart() {

	fmt.Println("Do you want to restart? (y/n)")
	buf := bufio.NewReader(os.Stdin)
	answer, err := buf.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	if strings.ToLower(strings.TrimSpace(answer)) == "y" {
		newGame()
	} else if strings.ToLower(strings.TrimSpace(answer)) == "n" {
		os.Exit(0)
	}
}