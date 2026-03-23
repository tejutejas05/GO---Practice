package main

import "fmt"

func main() {
	fmt.Println("Welcome to my Quiz game")
	
	// using the user input (taking the input from the user)
	fmt.Printf("Enter your name: \n")
	var name string
	fmt.Scan(&name)

	fmt.Printf("hello %v, You want to start the game?\n",name)

	fmt.Println("Enter your age")
	var age uint     
	fmt.Scan(&age)
	if age < 18 {
		fmt.Println("You are not eligible to play our game")
	} else {
		fmt.Println("You can play the game")
	}
}