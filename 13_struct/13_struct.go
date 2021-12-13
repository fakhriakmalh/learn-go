package main

import "fmt"

	
type student struct {
	name string
	grade int
}

type mahasiswa struct {
	name string
	grade int
}

type person struct {
    name string
    age  int
}

func main() {
	var s1 student
	s1.name = "arnan"
	s1.grade = 13
	fmt.Println("name : ", s1.name)
	fmt.Println("grade : ", s1.grade)

	var m1 = mahasiswa{}
	m1.name="alif"
	m1.grade = 12
	fmt.Println("name : ", m1.name)

	var m2 = mahasiswa{"ahmad",10}
	var m3 = mahasiswa{name:"dini",grade:11}

	fmt.Printf("student 2 name is %s and he is on grade %d\n", m2.name , m2.grade)
	fmt.Printf("student 3 name is %s and he is on grade %d\n", m3.name , m3.grade)

	fmt.Println("=================")

	//variabel objek pointer
	var m4 *mahasiswa = &m3
	fmt.Println("student 3, name :", m3.name)
	fmt.Println("student 4, name :", m4.name)

	m4.name = "ethan"
	fmt.Println("student 3, name :", m3.name)
	fmt.Println("student 4, name :", m4.name)

	fmt.Println("=================")

	//anonymous struct
	var s2 = struct {
        person
        grade int
    }{}
    s2.person = person{"wick", 21}
    s2.grade = 2

    fmt.Println("name  :", s2.person.name)
    fmt.Println("age   :", s2.person.age)
    fmt.Println("grade :", s2.grade)

}