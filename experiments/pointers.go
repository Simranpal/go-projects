package main

import "fmt"

func main() {
	var firstName *string = new(string)
	
	*firstName = "Simran"
	fmt.Println(*firstName)
	
	myName := "Sim"
	ptr := &myName
	fmt.Println(ptr, *ptr)

	myName = "Pal"
	fmt.Println(ptr, *ptr)
}
