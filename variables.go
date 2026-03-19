package main

import ("fmt")

func main() {

	var i int = 25
	fmt.Println(i)

	j:=10
	fmt.Println(j)

	fmt.Printf(" %v , %T \n", i ,i) // where %v ---> return the values of the variables
	fmt.Printf(" %v , %T \n", j ,j) // where %T ---> returns the datatype of the variables

	// we can assign many variables at a time 
	var (
		tejas string = "hello tejas"
		jevita string = "hello jev"
	//	prashi := 69 ----> We can't use the := method while declaring the many variables in the var( )
		prashi int = 69 // Now it will run because it is in the correct format ( varkeyword Variable_name Datatype)
	)

	fmt.Println(tejas , jevita , prashi)
}