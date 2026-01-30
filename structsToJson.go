package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// User defines the structure for user data.
// Fields must be exported (capitalized) to be processed by the json package.
type User struct {
	FirstName string `json:"first_name"` // Customizes the JSON key name
	LastName  string `json:"last_name"`  // Customizes the JSON key name
	Age       int    `json:"age,omitempty"` // omitempty means field is skipped if zero-valued
	IsActive  bool   `json:"is_active"`
}

func main() {
	// 1. Marshaling (Go struct to JSON)
	user := User{
		FirstName: "Jane",
		LastName:  "Doe",
		Age:       30,
		IsActive:  true,
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Marshaled JSON:")
	fmt.Println(string(jsonData)) // Output: {"first_name":"Jane","last_name":"Doe","age":30,"is_active":true}

	// 2. Unmarshaling (JSON to Go struct)
	jsonDataFromAPI := []byte(`{"first_name": "John", "last_name": "Smith", "is_active": false}`)
	var newUser User // Target variable

	err = json.Unmarshal(jsonDataFromAPI, &newUser) // Pass a pointer to the target struct
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nUnmarshaled Struct:")
	fmt.Printf("Name: %s %s\n", newUser.FirstName, newUser.LastName)
	fmt.Printf("Age: %d\n", newUser.Age) // Age will be the zero value (0) if omitted in the JSON string
	fmt.Printf("Active: %v\n", newUser.IsActive)
}
