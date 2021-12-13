package main

import "fmt"

func main(){
	//tipe data non desimal 
	var nond_1 uint8 = 34
	var nond_2 = -20
	fmt.Printf("positif = %d\n", nond_1)
	fmt.Printf("negatif = %d\n", nond_2)

	//tipe data boolean
	var exist bool = true
	fmt.Printf("exist? %t \n", exist)	

	//tipe data desimal
	var decimalNumber = 2.62
	fmt.Printf("bilangan desimal: %f\n", decimalNumber)

	//message
	var message = `Nama saya "John Wick".
	Salam kenal.
	Mari belajar "Golang".`
	fmt.Println(message)

	//constanta -> variabel yg nilainya tdk bisa diubah 
	const kon1 = "ilmuwan"
	const kon2 = "ilmu komputer"
	fmt.Printf("%s %s\n",kon1,kon2)
}