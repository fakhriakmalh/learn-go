package main
import "fmt" 
func main(){
	var chicken map[string]int
	chicken = map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("januari", chicken["januari"]) // januari 50
	fmt.Println("mei",     chicken["mei"])     // mei 0

	fmt.Println("===========")

	// cara horizontal
	var meat1 = map[string]int{"maret": 70, "april": 55}

	// cara vertical
	var meat2 = map[string]int{
		"juni": 60,
		"juli": 80 }
	fmt.Println("maret", meat1["maret"])
	fmt.Println("april", meat1["april"])
	fmt.Println("juni", meat2["juni"])
	fmt.Println("juli", meat2["juli"])

	fmt.Println("===========")

	//iterasi item map dengan for & range 
	for key, val := range meat1 { 
		fmt.Println(key, "\t:", val)
	}

	fmt.Println("===========")

	//hapus item di map 
	delete(meat1,"april")
	fmt.Println(len(meat1))
	fmt.Println(meat1)

	fmt.Println("===========")
	//cek item di map

	var chicken2 = map[string]int{"januari": 50, "februari": 40,"mei":35}
	var value, isExist = chicken2["mei"]

	if isExist {
		fmt.Println(value)
		fmt.Println(isExist)
	} else {
		fmt.Println("item is not exists")
	}

	fmt.Println("===========")
	//kombinasi slice dan map 

	var chickens_mapslice = []map[string]string{
		{"name": "chicken blue",   "gender": "male"},
		{"name": "chicken red",    "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
	}
	
	for _, ch := range chickens_mapslice {
		fmt.Println(ch["gender"], ch["name"])
	}

}