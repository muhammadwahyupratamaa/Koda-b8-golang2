package main

import "fmt"

type We struct{
	are Are
}
type Are struct{
	the The	
}
type The struct{
	best string 
}

type Hello struct{
	world string
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

fmt.Println(we.are.the.best)
fmt.Println(hello.world)
}