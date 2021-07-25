package main

import fmt

var x int = 4

func g(){
	fmt.Printf("%d\n", x)
}

func f(){
	fmt.Printf("%d\n", x)
}

func main(){
	g()
	f()
}
