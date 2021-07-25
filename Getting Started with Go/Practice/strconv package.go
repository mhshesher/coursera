package main

import "fmt"
import "strconv"

func main(){
	s1:="10"
	s2:=23
	s3,s4:=strconv.Atoi(s1)		//Atoi() returns two values, first will be asigned to r3 and second will be asigned to r4, just like python.
	fmt.Println(s3+12)
	fmt.Println(s4)

	fmt.Println(strconv.Itoa(s2)+"12")
}