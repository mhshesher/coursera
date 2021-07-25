package main

import "fmt"

func main(){
	var x int = 0

	switch{
	case x>0:
		fmt.Println("positive")
	case x<0:
		fmt.Println("negative")
	default:
		fmt.Println("zero")
	}
}