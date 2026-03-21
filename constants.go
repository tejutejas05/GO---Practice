package main

import ( "fmt" )

const PI = 3.14

func main() {

	const tej = "Hello From Tej!!"
	fmt.Println("Hiii?" , tej)

	fmt.Println("what is the constant value of the PI" , PI)


	//Integer Constant: 

	//85         /* decimal */
	//0213       /* octal */
	//0x4b       /* hexadecimal */
	//30         /* int */
	//30u        /* unsigned int */
	//30l        /* long */
	//30ul       /* unsigned long */
	//212         /* Legal */
	//215u        /* Legal */
	//0xFeeL      /* Legal */
	//078         /* Illegal: 8 is not an octal digit */
	//032UU       /* Illegal: cannot repeat a suffix */

	//Complex Constant
	//(0.0, 0.0) (-123.456E+30, 987.654E-29)

	//Floating type Constants
//	3.14159       /* Legal */
//	314159E-5L    /* Legal */
//	510E          /* Illegal: incomplete exponent */
//	210f          /* Illegal: no decimal or exponent */
//	.e55          /* Illegal: missing integer or fraction */




}