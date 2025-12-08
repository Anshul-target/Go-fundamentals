package main

import "fmt"

// Explicit type

func main()  {
// var x int =10
var y =0
// Short variable declaration only included inside the functions
// Walrus operator (nickname)
s:="hello"
fmt.Println(y,s)	
var x *int=&y
z:=12
fmt.Println(x,z)
redeclare()
}

func redeclare(){
	
	a:=2
	a,b:=2,3
	fmt.Println(a,b)
}