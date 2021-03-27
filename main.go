package main

func main() {
	newBoard()
	for i := 0; i < 8; i++ {
		if i%2 == 0 {
			playerRound(1)
		} else {
			playerRound(2)
		}
	}
}
