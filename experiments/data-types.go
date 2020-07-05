package main

import "fmt"

func main() {
	var i int 
	i = 42
	fmt.Println(i)
	
	var f float32 = 3.14
	fmt.Println(f)
	
	//use implicit initialization syntax
	firstName := "Simran"
	fmt.Println(firstName)
	
	b :=  true
	fmt.Println(b)
	
	c := complex(3, 4)
	fmt.Println(c)
}
