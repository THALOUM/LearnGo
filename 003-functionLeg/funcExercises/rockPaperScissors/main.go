package main

import (
	"fmt"
	"math/rand"
)

const (
	rock     string = "    _______\n---'   ____)\n      (_____)\n      (_____)\n      (____)\n---.__(___)\n"
	paper    string = "    _______\n---'   ____)____\n          ______)\n          _______)\n         _______)\n---.__________)\n"
	scissors string = "    _______\n---'   ____)____\n          ______)\n       __________)\n      (____)\n---.__(___)\n"
)

func main() {
	game := map[int]string{
		0: rock,
		1: paper,
		2: scissors,
	}

	var userInput int
	fmt.Println("What do you choose? Type 0 (Rock), 1 (Paper) or 2 (Scissors): ")
	fmt.Scanln(&userInput)
	fmt.Println(game[userInput])

	computerChoice := rand.Intn(3)
	fmt.Println("Computer choice: ")
	fmt.Println(game[computerChoice])

	if userInput-computerChoice == 2 || computerChoice-userInput == 2 {
		if userInput > computerChoice {
			fmt.Println("You Lose!")
		} else {
			fmt.Println("You Win!")
		}
	} else if userInput > computerChoice {
		fmt.Println("You Win!")
	} else if userInput == computerChoice {
		fmt.Println("It's a Draw!")
	} else {
		fmt.Println("You Lose!")
	}
}
