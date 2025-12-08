package main

import "fmt"

func main() {
	// Proper function
	var a int = sum(10, 20)
	fmt.Println(a)
	// Multiple return values
	a,b:=multiply(10,10)
	if b==nil {
		fmt.Println(a)
	} else{
		fmt.Println("Cannot multiply")
	}	
	// Logging the message
	logMessage()
	result:=sum1(1,2,3,4)
	fmt.Println(result)
	// Ignoring the return value
	result1,_:=divide(1,2)
	fmt.Println(result1)

	// Anonymous functions
	printSum:=func(a,b int) (int,error){
		return a+b,nil
	}
	fmt.Println(printSum(10,20))
// IIFE
printMultiple,errros:=func(a,b int)(int,error) {
return a*b,nil
}(10,20)
fmt.Println(errros)
fmt.Println(printMultiple)

}





func sum(a int, b int) int {
	return a + b
}
func multiply(a ,b int) (int,error){
	if b==0{
		return 0,fmt.Errorf("Multiply by zero")
	}
	return a*b,nil
}
func logMessage(){
	fmt.Println("Hello world")
}
// variadic function
func sum1(nums ...int)int{
	
	total:=0
	for _,n:=range nums{
		total+=n
	}
	return total
}
// Division
func divide(a,b int)(int,error){
	if b==0 {
		return 0,fmt.Errorf("Error format")
	}else{
		return a/b,nil
	}

}

