package main

import "fmt"

// soal 1
type We struct{
	are Are
}
type Are struct{
	the The	
}
type The struct{
	best string 
}

// Soal 2

type Hello struct{
	world string
}



// type Obj struct {
// 	str [] 
// }

// type Academy struct {
// 	academy string
// }


// Soal 5
type Number struct {
	first [] int
	second [] int
}




func main() {
we := We{
	are: Are{
		the: The{
			best: "KODA",
		},
	} ,
}

hello := Hello{
	world : "Hello World",
}

num := Number{
	first: []int{
		10,15,30,
	},
	second: []int{
		10,17,17,

	},
}

fmt.Println(we.are.the.best)
fmt.Println(hello.world)
fmt.Println(num.first[1] + num.second[2])
}