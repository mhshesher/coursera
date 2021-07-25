package main

import "fmt"

func main(){
	var a int8 = 4
	var b int16
	// b=a 				//not possible. because int8 and int16 are no the same
	b=int16(a)			//so here we converted int8 to int16
	fmt.Printf("%d\n",b)
}