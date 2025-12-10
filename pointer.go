package main

import "fmt"

type Person struct{
	Name string
	age int
}

// Pointer receiver
func (p *Person) UpdateAge(){
	p.age++
}
func main(){
	var a int=10
	var p *int=&a
	fmt.Println(a,p)
i:=0
j:=&i
fmt.Println(i,j,)
// Derefrencing
fmt.Println(*j)
*j=10
fmt.Println(i)


// Pass by reference
fmt.Println("Value of a ",a)
convert(&a)
fmt.Println("Value of a ",a)
}
func convert(address *int ){
*address=12

// Pointer to the structs



 var p1 *Person=&Person{}
(*p1).age =1
(*p1).Name ="Anshul Yadav"

fmt.Println("Age before calling",p1.age)
p1.UpdateAge()
fmt.Println("Age after calling",p1.age)
}



