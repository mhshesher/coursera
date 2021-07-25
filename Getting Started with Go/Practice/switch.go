package main

import "fmt"

func main(){

	var x int = 10

	switch x{
	case 11:
		fmt.Println("1")
	case 12:
		fmt.Println("2")
	default:
		fmt.Println("others")
	}
}