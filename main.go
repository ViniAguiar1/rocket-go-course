package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Guessing Game!")
	fmt.Println("An aleatory number will be generated between 1 and 100, try to guess it!")
	fmt.Println("You have 10 attempts to guess the number.")

	scanner := bufio.NewScanner(os.Stdin)
	chutes := [10]int64{}
	x := rand.Int64N(101)

	for i := range chutes {
		fmt.Println("Enter your guess:")
		scanner.Scan()
		chute := scanner.Text()
		chute = strings.TrimSpace(chute)

		chuteInt, err := strconv.ParseInt(chute, 10, 64)
		if err != nil {
			fmt.Println("Invalid input, please enter a number.")
			return
		}

		switch {
		case chuteInt < x:
			fmt.Println("Your guess is too low.", chuteInt)
		case chuteInt > x:
			fmt.Println("Your guess is too high.", chuteInt)
		case chuteInt == x:
			fmt.Println("Congratulations! You guessed the number!", chuteInt)
			return
		}

		chutes[i] = chuteInt
	}

	fmt.Printf(
		"Sorry, you didn't guess the number, which was %d\n"+
			"You had 10 attempts\n"+
			"These were your attempts: %v\n",
		x, chutes,
	)

}
