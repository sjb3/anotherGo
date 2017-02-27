package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func askInt(message string) int {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("%s\n>>", message)
		scanner.Scan()
		s := scanner.Text()
		n, err := strconv.Atoi(s)

		if err == nil {
			return n
		}
		fmt.Println("Error: '%s' is not a number\n", s)
	}
}

func askYesHigherLower(message string)int{
	scanner := bufio.NewScanner(os.Stdin)
	for {
		ftm.Printf("%s\n>>", message)
		scanner.Scan()
		s := scanner.Text()

		if s == "yes" {
			return 0
		}
		if s == "higher" {
			return 1
		}
		if s == "lower" {
			return -1
		}
		fmt.Printf("Please only enter, 'yes', 'higher', or 'lower' \n")
	}
}





func printTo100() {
	var i int = 0
	for {
		i = i + 1
		fmt.Println(i)
		if i == 100 {
			return
		}
	}
}

func keeper() {

	var answer int = rand.Intn(100)

	fmt.Printf("I'm thinking of a number ")

	for {
		g := askInt(" Take a Guess")
		if g == answer {
			fmt.Println(" You've got it!")
			return
		}
		if g < answer {
			fmt.Println(" Nope: higher")
		} else {
			fmt.Println(" Nope: lower")
		}
	}
}

func seeker(){
	fmt.Println("Think of a number")
	askYesHigherLower("Sat 'yes' when you think of a number){
	
	}

}
