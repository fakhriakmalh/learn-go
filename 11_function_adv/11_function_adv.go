package main

import "fmt"
import "strings"

func main() {
	var avg = calculate(4,3,6,7,5,6,8,5,6,9)
	var msg = fmt.Sprintf("Rata-rata : %.2f", avg) //menampilkan nilainya dalam bentuk string
	fmt.Println(msg)

	var number2 = []int{2,4,7,8,6,6,7,5,4}
	var avg2 = calculate(number2 ...)
	var msg2 = fmt.Sprintf("Rata-rata : %.2f", avg2)
	fmt.Println(msg2)

	//declare cardiac function 1
	urhobby("arnan", "sleeping", "eating")

	//declare cardiac function 2
	var hobbies = []string{"running", "loving"}
    urhobby("wick", hobbies...)

	fmt.Println("====================")

	//fungsi closure
	var getMinMax = func(n []int) (int, int) {
        var min, max int
        for i, e := range n {
            switch {
            case i == 0:
                max, min = e, e
            case e > max:
                max = e
            case e < min:
                min = e
            }
        }
        return min, max
    }

	var min ,max = getMinMax(number2)
	fmt.Printf("data : %v\nmin  : %v\nmax  : %v\n", number2, min, max) 

	fmt.Println("====================")

	var numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}

	//fungsi closure 2
    var newNumbers = func(min int) []int {
        var r []int
        for _, e := range numbers {
            if e < min {
                continue
            }
            r = append(r, e)
        }
        return r
    }(3)

    fmt.Println("original number :", numbers)
    fmt.Println("filtered number :", newNumbers)

	fmt.Println("====================")

	//print fungsi closure 3
	var mak1 , mak2 = findMax(numbers,3)
	var theNumber = mak2()
	fmt.Printf("jumlah yang ada di array %d \n",mak1)
	fmt.Printf("array yg tersisa = %v\n",theNumber)
}

//contoh fungsi dgn param cardiac adalah fungsi yang bisa define variabel yang isinya banyak
func calculate(numbers ...int) float64{
	var total int = 0
	for _, number := range numbers {
		total += number
	}

	var avg = float64(total) / float64(len(numbers))
	return avg
}

//contoh fungsi dgn parameter biasa dan cardiac 
func urhobby(name string , hobby ...string) {
	var hobbies_string = strings.Join(hobby, ", ")
	fmt.Printf("Hello, my name is: %s\n", name)
    fmt.Printf("My hobbies are: %s\n", hobbies_string)
}

//closure 3 , fungsi sebagai return 

func findMax(numbers []int, max int) (int, func() []int) {
	var res []int
	for _,e := range numbers {
		if e <= max {
			res = append(res,e)
		}
	}
	return len(res) , func() []int {
		return res
	}
}

