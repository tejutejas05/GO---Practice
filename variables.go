package main

import ("fmt")

func getting_numbers() (int, int, int) {
	return 23, 34 ,45
}


func main() {

	//var i int = 25
	//fmt.Println(i)

	//j:=10
	//fmt.Println(j)

	//fmt.Printf(" %v , %T \n", i ,i) // where %v ---> return the values of the variables
	//fmt.Printf(" %v , %T \n", j ,j) // where %T ---> returns the datatype of the variables

	// we can assign many variables at a time 
	//var (
	//	tejas string = "hello tejas"
	//	jevita string = "hello jev"
	//	prashi := 69 ----> We can't use the := method while declaring the many variables in the var( )
	//	prashi int = 69 // Now it will run because it is in the correct format ( varkeyword Variable_name Datatype)
	

	//var i int 
	//println(i) // it will print the default value of the int which is 0

	//fmt.Println(tejas , jevita , prashi)

	//Declaring the Multiple variables in a Same Type
	//var a,b,c int = 1,2,3
	//fmt.Printf(" %v , %T \n", a ,a)
	//fmt.Printf(" %v , %T \n", b ,b)
	//fmt.Printf(" %v , %T \n", c ,c)

	//Declaring the Multiple variables in a Diferent Type
	//var d,e,f = 1,"tejas",3.14
	//fmt.Printf(" %v , %T \n", d ,d)
	//fmt.Printf(" %v , %T \n", e ,e)
	//fmt.Printf(" %v , %T \n", f ,f)


	//Multiple Return values from a Function
	a,b,c := getting_numbers()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)


}