package main 
import "fmt"
func main() {
	var buah = []string{"apple", "grape", "banana", "melon", "grape"}
	fmt.Println(buah[0])
	var buah_new = buah[0:2]
	fmt.Println(buah_new)
	//len
	fmt.Println(len(buah_new))
	//cap -> biasanya ketika di slice mulai dari index x, maka cap akan beda dgn len. Kalau di slice di blkg gamasalah
	var cap_buah = buah[1:]
	var len_buah = buah[0:3]
	var coba_buah = buah[1:3]
	fmt.Println(cap(cap_buah) , len(len_buah), cap(coba_buah), len(coba_buah))

	//append
	var tambah_buah = append(buah,"watermelon")
	fmt.Println(tambah_buah)

	//copy 
	dst := make([]string, 3)
	src := []string{"watermelon", "pinnaple", "apple", "orange"}
	n := copy(dst, src)
	fmt.Println(dst) 
	fmt.Println(src) 
	fmt.Println(n)   

	//double slice 
	dslice_buah := buah[1:4:4]
	fmt.Println(dslice_buah)
	
}