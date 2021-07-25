package main

import "fmt"

func main(){
	const x int = 10
	var y int = 12
	fmt.Printf("%d %d\n", x,y)

	// x=15 	// we can not do that as it is constant
	y=23
	fmt.Printf("%d %d\n", x+y,y)
}