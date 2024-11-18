package main

import "fmt"

// Define the Person struct
type Person struct {
    name string
}

// Method to print the name of the Person
func (p Person) printName() {
    fmt.Println("Name:", p.name)
}

func main() {
    // Step 2: Create a Person variable and assign "John Doe" to the name member
    person := Person{name: "John Doe"}

    // Step 3: Call the printName method
    person.printName()
}