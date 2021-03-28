package main

import "fmt"

func main() {
	newGame()
}

func newGame() {
	newBoard()
	fmt.Println("Please take your position by typing the XY coordinate  (e.g. 13)")
	for i := 0; i < 9; i++ {
		if i%2 == 0 {
			playerRound(1, i)
		} else {
			playerRound(2, i)
		}
	}
}
