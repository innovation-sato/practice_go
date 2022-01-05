package main

import (
	"fmt"
)

func main() {
	print()

}

func print() {
	name := "miyu"
	fmt.Println(name)

	fmt.Printf("こんにちは%v\n", name)
}
