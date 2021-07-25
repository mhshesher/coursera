package main

import "fmt"

func main(){
	var a int16 = 10
	var b float64 = 11.15
	var n float64 = float64(a)
	
	// var c float64 = a+b 		//not possible of operation with different datatype
	// var d float64 = a-b		//not possible of operation with different datatype
	// var e float64 = a*b 		//not possible of operation with different datatype
	// var f float64 = a/b 		//not possible of operation with different datatype

	var c float64 = n+b
	var d float64 = n-b
	var e float64 = n*b
	var f float64 = n/b
	var g int16 = a%3
	var h int16 = a<<2
	var i int16 = a>>2

	fmt.Printf("%f\n", c)
	fmt.Printf("%f\n", d)
	fmt.Printf("%f\n", e)
	fmt.Printf("%f\n", f)
	fmt.Printf("%d\n", g)
	fmt.Printf("%d\n", h)
	fmt.Printf("%d\n", i)
}