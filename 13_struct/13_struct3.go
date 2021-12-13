package main
import "fmt"

type person struct {
    name string
    age  int
}

func main() {

	//kombinasi slice dan struct
	var allStudents = []person{
		{name: "Wick", age: 23},
		{name: "Ethan", age: 23},
		{name: "Bourne", age: 22},
	}

	for _, student := range allStudents {
		fmt.Println(student.name, "age is", student.age)
	}

	var allStudents2 = []struct {
		person
		grade int
	}{
		{person: person{"wick", 21}, grade: 2},
		{person: person{"ethan", 22}, grade: 3},
		{person: person{"bond", 21}, grade: 3},
	}
	
	for _, student := range allStudents2 {
		fmt.Println(student.person.age)
		fmt.Println(student)
		fmt.Println(student.grade)
		fmt.Println("=============")
	}
}

//ada method nested struct yg blum dicoba

