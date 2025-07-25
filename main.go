package main

import (
	"fmt"
	"time"
)

func main() {
	name := "Go Developers"
	fmt.Println("Azure for", name)
	time.Sleep(500)
	listing1()

	// i := 0
	// go func() {
	// 	i++
	// 	fmt.Println("***")
	// }()
	// go func() {
	// 	i++
	// 	fmt.Println("*****")
	// }()
}

// func main() {
// 	for i := 1; i < 7; i++ {
// 		go func(n int) {
// 			result := 0
// 			for j := 1; j <= n; j++ {
// 				result += j
// 			}
// 			fmt.Println(n, "-", result)
// 		}(i)
// 	}
// 	fmt.Scanln() // ждем ввода пользователя
// 	fmt.Println("The End")
// }

func listing1() {

	fmt.Println("*")
	i := 0

	go func() {
		i++
		fmt.Println("***")
	}()

	go func() {
		i++
		fmt.Println("*****")
	}()
}
