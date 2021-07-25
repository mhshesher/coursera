package main

import "fmt"

func main(){

	// var a[10] int
	// a[0]=10

	// fmt.Println(a[0])
	// fmt.Println(a[1])

	// var x[5] int = [5]int{1,2,3,4,5}
	// fmt.Println(x[3])

	// y:=[5]int{1,2,3,4,3}
	// fmt.Println(y[2])

	ar:=[5]int{2,4,6,7,3}

	// for i, v:=range ar{				// here i is the index and v is the value
	// 	fmt.Printf("%d %d\n",i,v)
	// }

	for i:=range ar{					//it prints indices only
		fmt.Println(i)
	}
}