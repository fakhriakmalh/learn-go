package main

import "fmt"

func main(){
	//print array
	var bajak_laut[3]string
	bajak_laut[0] = "Monkey"
	bajak_laut[1] = "D"
	bajak_laut[2] = "Luffy"

	fmt.Println( bajak_laut[0], bajak_laut[1], bajak_laut[2] )

	//inisialisasi array gaya vertikal & horizontal
	var fruits [4]string

	// cara horizontal
	fruits  = [4]string{"apple", "grape", "banana", "melon"}

	// cara vertikal
	fruits  = [4]string{
		"apple",
		"grape",
		"banana",
		"melon",
	}
	fmt.Println(fruits)

	//inisialisasi array tanpa jumlah elemen nya 
	var numbers = [...]int{2, 3, 2, 4, 3}
	fmt.Println("data array \t:", numbers)
	fmt.Println("jumlah elemen \t:", len(numbers))

	fmt.Println("===========")
	//looping dengan for 
	var buah = [] string{"apple", "grape", "banana", "melon"}
	for i := 0  ;  i < len(buah) ; i++ { 
		fmt.Printf("elemen %d : %s\n", i, fruits[i])
	}
	fmt.Println("===========")
	//looping dengan range 
	for i,namabuah := range buah {
		fmt.Printf("elemen %d : %s\n", i, namabuah)
	}
	fmt.Println("===========")
	//menggunakan underscore agar tidak ada variabel nganggur
	for _, namabuaha := range buah {
		fmt.Printf("nama buah : %s\n", namabuaha)
	}

	
	//Alokasi Elemen Array Menggunakan Keyword make
	var fruits_2 = make([]string, 2)
	fruits_2[0] = "apple"
	fruits_2[1] = "manggo"

	fmt.Println(fruits_2)

	//cek isi array / list
	fmt.Println("===========")
	var cek_buah = stringInSlice("pisang",buah)
	fmt.Println(cek_buah)


}

func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}
	