package main

import "fmt"

type Person struct {
	Fname    string
	Lname    string
	Gender   string
	Age      int
	Number   int
	Company  string
	Position string
}

func (p Person) identity() {
	fmt.Println("\n\nIdentity")
	fmt.Printf("Name: %s %s \n", p.Fname, p.Lname)
	fmt.Printf("Age: %d \n", p.Age)
	fmt.Printf("Gender: %s\n", p.Gender)
}

func (p Person) emloyee() {
	fmt.Println("\n\nEmployee")
	fmt.Printf("Name: %s %s \n", p.Fname, p.Lname)
	fmt.Printf("Contact Number: %d\n", p.Number)
	fmt.Printf("Working in %s as %s.\n", p.Company, p.Position)
}

func (p Person) introduce() {
	fmt.Print("\n\nIntroduction\n")
	fmt.Printf("Hello, my name is %s %s. I am working at %s as %s. I'm %d years old.", p.Fname, p.Lname, p.Company, p.Position, p.Age)
}

func main() {
	person := Person{
		Fname:    "Arth",
		Lname:    "Patel",
		Gender:   "Male",
		Age:      24,
		Number:   8899889988,
		Company:  "Google, India",
		Position: "Software Developer",
	}
	person.identity()

	person.emloyee()

	person.introduce()

	fmt.Printf("\n\nPerson Age: %d", person.Age)
}
