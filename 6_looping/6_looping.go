package main

import "fmt"

func main(){
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}
	fmt.Println("===========")
	//looping dengan argumen hanya kondisi
	var i = 0
	for i < 5 {
    	fmt.Println("Angka", i)
    i++ }
	
	fmt.Println("===========")
	//looping for tanpa argumen
	var j = 0
	for {
    fmt.Println("Angka", j)
	j++
    if j == 5 {
        break}
	}

	fmt.Println("===========")
	// looping dengan break dan continue 
	for i := 1 ; i <= 10 ; i++ {
		if i % 2 == 1 {
			continue
		}
		if i > 8 {
			break
		}
		fmt.Println("Angka",i)
	}

	fmt.Println("===========")
	//looping bikin segitiga
	for i := 5 ; i > 0 ; i-- { 
		for j := 0 ; j < i ; j++ {
			fmt.Print(j," ")
		}
		fmt.Println()
	}  

	fmt.Println("===========")
	//looping bikin segitiga
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
	
	//looping dengan fitur label (outerloop)
	outerLoop:
	for i := 0; i < 5; i++ {
    	for j := 0; j < 5; j++ {
        	if i == 3 {
            	break outerLoop
        }
        fmt.Print("matriks [", i, "][", j, "]", "\n")
    }
}
}