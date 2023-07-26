package main

import "fmt"

func main() {
	divTres()
}
func divTres() {
	for i := 1; i < 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("pin pan")
		} else if i%3 == 0 {
			fmt.Println("pin")
		} else if i%5 == 0 {
			fmt.Println("pan")
		} else {
			fmt.Println(i)
		}
	}
}
