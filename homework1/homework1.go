package main

import (
	"fmt"
	"math/rand"
	"time"
)

func guessNumber(secretNumber int) int {
	var yourNumber int
	var result int
	for {
		fmt.Print("Please enter your guess: ")
		_, err := fmt.Scanf("%d", &yourNumber)
		fmt.Scanln() //读取缓冲区的回车
		if err != nil {
			fmt.Println("Invalid input, please try again.")
			continue
		}
		if yourNumber == secretNumber {
			result = 0
			break
		}
		if yourNumber < secretNumber {
			result = 1
		}
		if yourNumber > secretNumber {
			result = 2
		}
		break
	}
	return result
}

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	var count int
	for {
		receiveNumber := guessNumber(secretNumber)
		count++
		switch receiveNumber {
		case 0:
			fmt.Printf("You are right! It took you %d guesses.\n", count)
			return
		case 1:
			fmt.Println("Guess it's small")
		case 2:
			fmt.Println("Guess it's big")
		}
	}
}
