package main

import "fmt"

func birthday(ptrAge *int) {
	*ptrAge++
}

func main() {
	i := 34
	fmt.Println(i)
	birthday(&i)
	fmt.Println(i)
	var a string
	fmt.Scanf("%s", &a)
	fmt.Println("this is the scan", a)
}
