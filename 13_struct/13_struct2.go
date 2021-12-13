package main

import "fmt"

type person struct {
    name string
    age  int
}

type student struct {
    person
    age   int
    grade int
}

type profil_tukang struct {
    name string
    age  int
}

type tukang struct {
    profil_tukang
    grade int }


func main() {
    var s1 = student{}
    s1.name = "wick"
    s1.age = 21        // age of student
    s1.person.age = 22 // age of person -> variabel nya sama
    s1.grade = 80

    fmt.Println(s1.name)
    fmt.Println(s1.age)
    fmt.Println(s1.person.age) //embedded struct
    fmt.Println(s1.grade)

    var pr1 = profil_tukang{name: "wick", age: 21}
    var tk1 = tukang{profil_tukang: pr1, grade: 2}
    fmt.Println("name  :", pr1.name)
    fmt.Println("age   :", pr1.age)
    fmt.Println("grade :", tk1.grade)


}


