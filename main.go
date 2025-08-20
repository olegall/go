package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	name := "Go Developers"
	fmt.Println("Azure for", name)
	//listing1()

	// read 10 fibonacci numbers from channel returned by `fib` function

	//listing2()
	// i := 0
	// go func() {
	// 	i++
	// 	fmt.Println("***")
	// }()
	// go func() {
	// 	i++
	// 	fmt.Println("*****")
	// }()

	main_habr()

	fmt.Println("*** 9.3 ***")
	//s := []int{1, 2, 3}
	/*for _, i := range s { // Итерации по каждому элементу
		go func() {
			fmt.Print(i) // Обращение к переменной цикла
		}()
	}*/
	/*for _, i := range s {
		val := i // Создание переменной, локальной для каждой итерации
		go func() {
			fmt.Print(val)
		}()
	}*/
	/*for _, i := range s {
		go func(val int) {
			fmt.Print(val)
		}(i) // Вызов этой функции с текущим значением переменной i
	}*/

	time.Sleep(1000)
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
		fmt.Println("**")
	}()

	go func() {
		i++
		fmt.Println("***")
	}()

	time.Sleep(100)
	fmt.Println(i) // 0, 1, 2
}

func listing2() {
	var i int64

	go func() {
		atomic.AddInt64(&i, 1)
	}()

	go func() {
		atomic.AddInt64(&i, 1)
	}()

	//time.Sleep(100)
	fmt.Println("atomic", i) // 0, 1, 2
}
