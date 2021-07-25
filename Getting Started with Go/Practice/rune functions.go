package main

import "fmt"
import "unicode"
import "strconv"

func main(){
	r1:='a'
	r2:='9'
	r3:=' '
	r4:='S'
	r5:=';'
	

	fmt.Println(unicode.IsDigit(r1))
	fmt.Println(unicode.IsDigit(r2))

	fmt.Println(unicode.IsSpace(r2))
	fmt.Println(unicode.IsSpace(r3))

	fmt.Println(unicode.IsLetter(r1))
	fmt.Println(unicode.IsLetter(r2))

	fmt.Println(unicode.IsLower(r1))
	fmt.Println(unicode.IsLower(r4))

	fmt.Println(unicode.IsUpper(r1))
	fmt.Println(unicode.IsUpper(r4))

	fmt.Println(unicode.IsPunct(r1))
	fmt.Println(unicode.IsPunct(r5))
}