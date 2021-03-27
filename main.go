package main

func main() {
	newGame()
}

func newGame() {
	newBoard()
	for i := 0; i < 9; i++ {
		if i%2 == 0 {
			playerRound(1, i)
		} else {
			playerRound(2, i)
		}
	}
}
