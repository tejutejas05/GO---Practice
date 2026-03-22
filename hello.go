package main

import( "fmt" )
import "unicode/utf8"

func main() {
	fmt.Println("Hello From Lanora")


// to find the lenght of the string in GO (we dont use the Len function)
// if we use the len() --> it will return ASICC key value 
// for that we can import the built in package " unicode/utf8 " and we have to use " utf8.RuneCountInString() "

	var tej string = "ho"
	fmt.Println(utf8.RuneCountInString(tej))

	var myRune rune = 't'
	fmt.Println(myRune)



}