package main

import "fmt"
import "strings"

func main(){
	var s string

	fmt.Scanln(&s)

	var s2 string = strings.ToLower(s)

	var tmp int = len(s2)
	if(string(s2[0])=="i" && string(s2[tmp-1])=="n" && strings.Contains(s2,"a")){
		fmt.Println("Found!")
	}else{
		fmt.Println("Not Found!")
	}
}