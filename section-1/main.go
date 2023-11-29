package main

import (
	"fmt"
	"strings"
)

func main () {
	m1 := "Hello World!"
	m2 := "String: "

	m1 = "Another value"

	fmt.Println(m2 + m1, strings.Split(m1, " "), strings.Contains(m1, m2), len(m1))
}

// func main() {
// 	// different vars in same block with same type
// 	// var (
// 	// 	m1 = 2
// 	// 	m2 = 3
// 	// )
// 	// var instance
// 	// var m1 int
// 	// m1 = 2
// 	// var conversion
// 	// var m1 int32
// 	// var m2 int64
// 	// fmt.Println(int64(m1) + m2)
	
// 	m1 := 2
// 	m2 := 3

// 	// fmt.Println("Hello World!")
// 	fmt.Println(m1 + m2)
// }