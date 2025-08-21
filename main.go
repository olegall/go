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

	//fmt.Println("*** 9.9 ***")
	// s := make([]int, 1)
	// go func() { // Добавление к s нового элемента в новой горутине
	// 	s1 := append(s, 1)
	// 	fmt.Println(s1)
	// }()
	// go func() { // То же самое
	// 	s2 := append(s, 1)
	// 	fmt.Println(s2)
	// }()

	// s := make([]int, 0, 1)
	// go func() { // Добавление к s нового элемента в новой горутине
	// 	s1 := append(s, 1)
	// 	fmt.Println(s1)
	// }()
	// go func() { // То же самое
	// 	s2 := append(s, 1)
	// 	fmt.Println(s2)
	// }()

	// s := make([]int, 0, 1)
	// go func() {
	// 	sCopy := make([]int, len(s), cap(s))
	// 	copy(sCopy, s) // Создание копии для использования append на копии среза
	// 	s1 := append(sCopy, 1)
	// 	fmt.Println(s1)
	// }()
	// go func() {
	// 	sCopy := make([]int, len(s), cap(s))
	// 	copy(sCopy, s) // То же самое
	// 	s2 := append(sCopy, 1)
	// 	fmt.Println(s2)
	// }()

	fmt.Println("*** 9.11 ***")
	// wg := sync.WaitGroup{}
	// var v uint64
	// for i := 0; i < 3; i++ {
	// 	go func() { // Создается горутина
	// 		wg.Add(1)               // Увеличивается значение счетчика группы ожидания
	// 		atomic.AddUint64(&v, 1) // Атомарно увеличивается значение переменной v
	// 		wg.Done()
	// 	}()
	// }
	// wg.Wait()
	// fmt.Println(v)

	// wg := sync.WaitGroup{}
	// var v uint64
	// wg.Add(1) // Увеличивается значение счетчика группы ожидания
	// for i := 0; i < 3; i++ {
	// 	go func() { // Создается горутина
	// 		atomic.AddUint64(&v, 1) // Атомарно увеличивается значение переменной v
	// 		wg.Done()
	// 	}()
	// }
	// wg.Wait()
	// fmt.Println(v)

	// wg := sync.WaitGroup{}
	// var v uint64
	// for i := 0; i < 3; i++ {
	// 	wg.Add(1)   // Увеличивается значение счетчика группы ожидания
	// 	go func() { // Создается горутина
	// 		atomic.AddUint64(&v, 1) // Атомарно увеличивается значение переменной v
	// 		wg.Done()
	// 	}()
	// }
	// wg.Wait()
	// fmt.Println(v)

	// fmt.Println("*** 9.12 ***")
	// type Donation struct {
	// 	mu      sync.RWMutex
	// 	balance int
	// }
	// donation := &Donation{}
	// // горутины-слушатели
	// f := func(goal int) { //Создание замыкания
	// 	donation.mu.RLock()
	// 	for donation.balance < goal { // Проверка достижения цели
	// 		donation.mu.RUnlock()
	// 		donation.mu.RLock()
	// 	}
	// 	fmt.Printf("$%d goal reached\n", donation.balance)
	// 	donation.mu.RUnlock()
	// }
	// go f(10)
	// go f(15)
	// // горутина обновления
	// go func() {
	// 	for { // Продолжение увеличения баланса
	// 		time.Sleep(time.Second)
	// 		donation.mu.Lock()
	// 		donation.balance++
	// 		donation.mu.Unlock()
	// 	}
	// }()

	// type Donation struct { // пример выводит в консоль
	// 	balance int
	// 	ch      chan int // Обновление Donation таким образом, что он содержит каналы
	// }
	// donation := &Donation{ch: make(chan int)}
	// // горутины-слушатели
	// f := func(goal int) {
	// 	for balance := range donation.ch { //Получение обновлений каналов
	// 		if balance >= goal {
	// 			fmt.Printf("$%d goal reached\n", balance)
	// 			return
	// 		}
	// 	}
	// }
	// go f(10)
	// go f(15)
	// // горутина обновления
	// for {
	// 	time.Sleep(time.Second)
	// 	donation.balance++
	// 	donation.ch <- donation.balance
	// }

	// type Donation struct {
	// 	cond *sync.Cond // Добавление *sync.Cond
	// 	balance int
	// }
	// donation := &Donation {
	// 	cond: sync.NewCond(&sync.Mutex{}), // sync.Cond  использует мьютекс
	// }
	// // горутины-слушатели
	// f := func(goal int) {
	// 	donation.cond.L.Lock()
	// 	for donation.balance < goal {
	// 		donation.cond.Wait()
	// 	}
	// 		fmt.Printf("%d$ goal reached\n", donation.balance)
	// 		donation.cond.L.Unlock()
	// }
	// go f(10)
	// go f(15)
	// // горутина обновления
	// for {
	// 	time.Sleep(time.Second)
	// 	donation.cond.L.Lock()
	// 	donation.balance++ // Увеличение баланса в рамках блокировки/разблокировки
	// 	donation.cond.L.Unlock()
	// 	donation.cond.Broadcast() // Трансляция факта выполнения условия (баланс обновлен)
	// }

	// fmt.Println("*** 9.14 ***")
	// type Counter struct {
	// 	mu sync.Mutex
	// 	counters map[string]int
	// }
	// func NewCounter() Counter { // Фабричная функция
	// 	return Counter{counters: map[string]int{}}
	// }
	// func (c Counter) Increment(name string) {
	// 	c.mu.Lock() // Увеличение значения счетчика в критической секции
	// 	defer c.mu.Unlock()
	// 	c.counters[name]++
	// }

	// type Counter struct {
	// 	mu sync.Mutex
	// 	counters map[string]int
	// }
	// func NewCounter() Counter { // Фабричная функция
	// 	return Counter{counters: map[string]int{}}
	// }
	// func (c *Counter) Increment(name string) {
	// 	c.mu.Lock() // Увеличение значения счетчика в критической секции
	// 	defer c.mu.Unlock()
	// 	c.counters[name]++
	// }

	// type Counter struct {
	// 	mu *sync.Mutex
	// 	counters map[string]int
	// }
	// func NewCounter() Counter { // Фабричная функция
	// 	return Counter{counters: map[string]int{}}
	// }
	// func (c Counter) Increment(name string) {
	// 	c.mu.Lock() // Увеличение значения счетчика в критической секции
	// 	defer c.mu.Unlock()
	// 	c.counters[name]++
	// }

	// counter := NewCounter()
	// go func() {
	// 	counter.Increment("foo")
	// }()
	// go func() {
	// 	counter.Increment("bar")
	// }()

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
