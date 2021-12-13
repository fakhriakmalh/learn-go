package main
import "fmt"
func main() {
    //deklarasi variabel nama awal
    var firstName string = "john"
    //deklarasi variabel nama akhir
    var lastName string
    lastName = "wick"

    fmt.Printf("halo %s %s!\n", firstName, lastName)

    //deklarasi variabel without data type , jadi ga perlu ada var dan tipe data nya 
    nama_1 := "data"
    nama_2 := "engineer"
    fmt.Printf("your job role is %s %s!\n", nama_1, nama_2)
    
    //deklarasi multivariabel 
    var fourth, fifth, sixth string = "empat", "lima", "enam"
    fmt.Printf("%s %s %s \n",fourth , fifth, sixth)
}