package main

import "fmt"



func main() {
	
	//Go String!
	var char rune = '정'
	fmt.Printf("%T\n", char)	
	fmt.Printf("%c\n", char)
	fmt.Println(char)

	str1 := "가나다라마"
	str2 := "abcde"
	fmt.Println("==============================================")
	fmt.Println(len(str1))
	fmt.Println(len(str2))

	str := "Hello World"
	runes := []rune{72, 101, 108, 108, 111, 32, 87,111,114,108,100}
	fmt.Println("==============================================")
	fmt.Println(str)
	fmt.Println(string(runes))

	newstr := "안녕 나는 정다운입니다. haha~"
	newrunes := []rune(newstr)

	fmt.Println(len(newstr))
	fmt.Println(len(newrunes))

}