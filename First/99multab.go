package main

import "fmt"

func multab() {
	for a := 1; a <= 9; a++ {
		for b := 1; b <= a; b++ {
				c := a * b
				fmt.Printf("%d * %d = %d \t",b,a,c)
		}
		fmt.Printf("\n")
	}
}

func main() {
	multab()
}
