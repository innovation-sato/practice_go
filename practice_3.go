// ポインタについて

package main

import "fmt"

func main() {
	a := 5
	var pa *int
	pa = &a // &a = aのアドレス
	// paの値を取りたい時＝　*pa

	fmt.Println(pa)
	fmt.Println(*pa)
}
