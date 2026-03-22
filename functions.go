package main

import "fmt"

func main(){
	var value string = "Hi from tejas"
	tejas(value)
}

//Decalring the new functions and calling it in the main function 
func tejas(value string){ //With argument 
	fmt.Println(value)
}