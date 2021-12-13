package main

import "fmt"

func main(){
	var nilai = 50
	if nilai == 100 {
		fmt.Println("lulus dengan sempurna")
	} else if nilai > 49 {
		fmt.Println("lulus")
	} else if nilai == 49 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Println("tidak lulus")
	}

	var score = 10000
	if divide_hundred := score / 100 ; divide_hundred == 100 {
		fmt.Println("terbaik")
	} else if divide_hundred >= 70 {
		fmt.Println("good enough")
	} else {
		fmt.Println("not quite")
	}

	var poin = 8
	switch poin {
	case 8 :
		fmt.Println("perfect")
	case 7,6,5 : 
		fmt.Println("awesome")
	default : 
		fmt.Println("not bad")
	}

	var point = 8
	switch { 
		case point == 8 :
			fmt.Println("awesome")
		case (point < 8) && (point > 3):
			fmt.Println("not bad")
		default :
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}
}