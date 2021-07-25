package main

import "fmt"

func main(){
	var x int = 46

	if(x>=80){
		fmt.Println("A+")
	}else if(x>=70){
		fmt.Println("A")
	}else if(x>=60){
		fmt.Println("A-")
	}else if(x>=50){
		fmt.Println("B")
	}else{
		fmt.Println("F")
	}
}