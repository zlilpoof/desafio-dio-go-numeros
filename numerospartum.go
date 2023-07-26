package main

import "fmt"

func main() {
	divTres()
}
func divTres() {
	for i := 1; i < 100; i++ {

		if i%3 == 0 {
			fmt.Println(i)
		}

	}
}
