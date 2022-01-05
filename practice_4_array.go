// 配列

package main

import "fmt"

func main() {
	var a [5]int // a[0] - a[4]
	a[2] = 3
	a[4] = 10
	fmt.Println(a)
	fmt.Println(a[2])

	// 宣言と代入をまとめると...
	b := [3]int{1, 3, 5}
	// さらにそれはこれでOK
	b := [...]int{1, 3, 5}

	fmt.Println(b)
	// 配列の個数を知りたいときは
	fmt.Println(len(b))
}
