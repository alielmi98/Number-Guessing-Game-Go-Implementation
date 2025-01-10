package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")

	var chances int
	fmt.Println("Please select the difficulty level:")
	fmt.Println("1. Easy (10 chances)")
	fmt.Println("2. Medium (5 chances)")
	fmt.Println("3. Hard (3 chances)")

	var choice int
	fmt.Print("Enter your choice: ")
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Invalid input!")
		os.Exit(1)
	}

	switch choice {
	case 1:
		chances = 10
	case 2:
		chances = 5
	case 3:
		chances = 3
	default:
		fmt.Println("Invalid choice!")
		os.Exit(1)
	}

	fmt.Printf("Great! You have selected the difficulty level with %d chances.\n", chances)
	fmt.Println("Let's start the game!")

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	targetNumber := r.Intn(100) + 1

	for attempts := 1; attempts <= chances; attempts++ {
		fmt.Print("Enter your guess: ")

		var guess int
		_, err := fmt.Scan(&guess)
		if err != nil {
			fmt.Println("Invalid input! Please enter a number.")
			attempts--
			continue
		}

		if guess < targetNumber {
			fmt.Printf("Incorrect! The number is greater than %d.\n", guess)
		} else if guess > targetNumber {
			fmt.Printf("Incorrect! The number is less than %d.\n", guess)
		} else {
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts.\n", attempts)
			return
		}
	}

	fmt.Printf("Sorry! You've run out of chances. The correct number was %d.\n", targetNumber)
}
