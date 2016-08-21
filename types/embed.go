package types

import "fmt"

// Person type
type Person struct {
	Name string
}

// SayName say person name
func (person Person) SayName() {
	fmt.Println("My name is ", person.Name)
}

// Student type
type Student struct {
	Person
	Name string
	ID   string
}

// SampleEmbed sample embed
func SampleEmbed() {
	student := Student{
		Person: Person{
			Name: "Eko",
		},
		Name: "Kurniawan",
		ID:   "10106031",
	}

	fmt.Println(student)
	fmt.Println(student.Name)
	fmt.Println(student.Person.Name)

	student.SayName()
}
