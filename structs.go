package main

import "fmt"

type Person struct {
    Name string
    Age  int
}
type Address struct {
    Street, City, State string
	Pincode int
}

func main(){
	person1 := Person{Name: "Alice", Age: 30}
	fmt.Println(person1.Name, person1.Age)
}