package main

import "fmt"

func main(){
	var x string
	var y string

	fmt.Scan(&x)
	fmt.Scan(&y)

	fmt.Println(x)
	fmt.Println(y)

	fmt.Println(len(x))

	fmt.Scanln(&x)
	fmt.Scanln(&y)

	fmt.Println(x)
	fmt.Println(y)

	fmt.Println(len(x))
}