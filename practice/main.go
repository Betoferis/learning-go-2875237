package main

import (
	"fmt"
	"time"
)


// Define a struct
type Person struct {
	Name string
	Age  int
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	fmt.Println("Hello from Go!")

	var a string = "Hello"
	var b, c int = 1, 2
	d := 3.5 // Type inference
	fmt.Println(a, b, c, d)

	// If-else statement
	if x := 42; x > 10 {
		fmt.Println("x is greater than 10")
	} else {
		fmt.Println("x is less than or equal to 10")
	}

	// For loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Arrays
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println(arr)

	// Slices
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	// Map declaration
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	fmt.Println(m)

	// Create an instance of the struct
	p := Person{Name: "Alice", Age: 30}
	fmt.Println(p)

	go say("Hello")
	say("World")
}
