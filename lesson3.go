package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func askInt(message string) int {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("%s\n>>", message)
		scanner.Scan()
		s := scanner.Text()
		n , err := strconv.Atoi(s)

		if err == nil{
			return n
		}
		fmt.Println("Error: '%s' is not a number\n", s)
	}
}

func main() {

	var answer int = 35

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

	// a := askInt("Choose a number")
	// b := askInt("Choose another number")
	// c := a + b
	// fmt.Printf("%d + %d = %d\n", a, b, c)

	// n := askInt("Pick a number")
	//
	// fmt.Printf("You chose :  %d", n)
	// var a int
	// var b int
	//
	// fmt.Printf("Hi, Please give me a number. \n")
	// fmt.Scanf("%d\n", &a)
	// fmt.Printf("Hi, Please give me another number. \n")
	// fmt.Scanf("%d\n", &b)
	//
	// c := a + b
	//
	// fmt.Printf("%d + %d = %d", a, b, c)
}