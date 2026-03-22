package main

import "fmt"
import "errors"

func main(){
	var value string = "Hi from tejas"
	tejas(value)
	// The variables which are used in the intDiv func have been declared in the main function 
	var num int = 10
	var den int = 5
	var result, rem, err int = intDiv(num,den)
//	if err!=nil{
//		fmt.Printf(err.Error())
//	}else if rem == 0{
//		fmt.Printf("the result of the integer divison is %v", result)
//	}else{
//		fmt.Printf("%v is the divison integer and %v is the remiander" , result, rem)
//	}

}

//Decalring the new functions and calling it in the main function 
func tejas(value string){ //With argument 
	fmt.Println(value)
}

//Division function (which returns the multiple value)
func intDiv(num int , den int) (int, int, error) {
	var err error 
	if den==0{
		err = errors.New("cannot divide by zero")
		return 0,0,err
	}
	var result int = num/den
	var rem int = num%den
	return result, rem
}