package main

import "fmt"
// PublicVar is a public variable and can be accessed from other packages.
var PublicVar = "I am public"

// privateVar is a private variable and can only be accessed within mypackage.
var privateVar = "I am private"

// PublicFunction is a public function.
func PublicFunction() string {
    // Both PublicVar and privateVar can be accessed here
    return PublicVar + " " + privateVar 
}

// privateFunction is a private function.
func privateFunction() string {
    return privateVar
}

func main(){
	fmt.Println("nani")
	fmt.Println(PublicFunction())
}
