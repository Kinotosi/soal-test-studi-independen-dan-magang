package main

import "fmt"

func main() {
	// Perulangan dari 1 sampai 100
	for i := 1; i <= 100; i++ {
		// Jika nilai i habis dibagi 3 dan 5, maka menampilkan "Hello World"
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Hello World")
			// Jika nilai i habis dibagi 3, maka menampilkan "Hello"
		} else if i%3 == 0 {
			fmt.Println("Hello")
			// Jika nilai i habis dibagi 5, maka menampilkan "World"
		} else if i%5 == 0 {
			fmt.Println("Wolrd")
		} else {
			fmt.Println(i)
		}
	}
}
