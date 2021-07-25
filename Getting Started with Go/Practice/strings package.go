package main

import "fmt"
import "strings"

func main(){
	s1:="abcd"
	s2:="acbd"
	s3:="abcd"
	s4:="asasasasas"
	s5:="ASHajshajhasJHSAJKJSNajhsjHAS"

	fmt.Println(strings.Compare(s1,s2))
	fmt.Println(strings.Compare(s2,s1))
	fmt.Println(strings.Compare(s1,s3))

	fmt.Println(strings.Contains(s1,"as"))
	fmt.Println(strings.Contains(s1,"ab"))

	fmt.Println(strings.HasPrefix(s1,"bc"))
	fmt.Println(strings.HasPrefix(s3,"ab"))

	fmt.Println(strings.Index(s4,"sa"))

	fmt.Println(strings.Replace(s4,"sa","in",2))

	fmt.Println(strings.ToLower(s5))

	fmt.Println(strings.ToUpper(s5))
}