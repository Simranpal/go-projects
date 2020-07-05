package main

import "fmt"

const (
	first = iota
	second
	third = iota
)

const (
	fourth = iota
	fifth
)

func main() {

	fmt.Println(first, second, third, fourth, fifth)
}
