package main

import "fmt"

func main() {
	// Creating the struct
type Person struct {
name string
age  int
}

// intializing a struct
p := Person{
name: "Anshul",
age:  34,
}
fmt.Println(p.age,p.name)
// Using new keywords
p1 := new(Person)
fmt.Println(p1.age,p1.name)

// Zero value initialization
var p2 Person
fmt.Println(p2.age,p2.name)

// Initializing with the pointer
p3 := &Person{
name: "Anshul",
age:  12,
}
fmt.Println((*p3).age,(*p3).name)
fmt.Println(p3.age,p3.name)

// Anonymous fields
type Address struct{
	city string
	state string
}
type Employee struct{
	name string
	address Address 
}

// Accessing this 
e:=Employee{name: "Anshul Yadav",address: Address{city:"Ghazipur",state:"Up"},
}

fmt.Println(e.address.city)
// Anonymous (Embedded) fields
type Engine struct{
	HorsePower int
}
type Car struct{
	name string
	Engine
}
car:=Car{
	name:"Toyata",
	Engine: Engine{HorsePower: 240},
}
fmt.Println(car.name,car.HorsePower)


// Methods on the structs


// Initializing the struct
p5:=Person1{name:"Anshul "}
p5.Greet1()


type doc struct{
	Name string
}
a:=doc{Name:"Hi there"}
b:=doc{Name:"How are you"}
fmt.Println(a,b,a==b)
}
func (p Person1) Greet1(){
fmt.Println("Hello",p.name)
}
type Person1 struct{
	name string
}

