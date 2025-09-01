package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("I want a cap of coffe!!!")
	a := 10
	if a < 18 {
		fmt.Println("Поешь растишку")
	}

	for i := 1; i < 5; i++ {
		fmt.Println("Поешь растишку!!!😉")
	}

	x := rand.Intn(10)

	fmt.Println(x)

	time.Sleep(1000 * time.Millisecond)

	fmt.Println(x + 100)
	d := foo1(600, 66)
	fmt.Println(d)
	nam := "america"
	fmt.Printf("type %T", nam)
	arr := [5]int{3, 10, 18, 9, 666}
	fmt.Println("\n")
	fmt.Println(arr)
	fmt.Println(3, arr[3])
}

func foo1(x int, y int) int {
	var z int = x + y
	return z
}
