package main

import "fmt"

func main() {
	// Basically sama aja sih pointer sama variabel. bedanya tuh pas di array ya kali
	var pt *int // membuat pointer
	num1 := 34  // ini variabel

	pt = &num1 // & itu menandakan address nya. jadi num1 itu addres, 34 itu value

	fmt.Println("Addresnya adalah", pt)
	fmt.Println("Addresnya adalah", pt)
	fmt.Println("Valuenya adalah", *pt)
}
