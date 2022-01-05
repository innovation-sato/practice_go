// スライス

package main

import (
	"fmt"
)

func main() {
	a := [5]int{2, 10, 8, 15, 4}
	// aの配列から切り出したりできる
	s: = a[2:4] // [8, 15]
	s[1] //15
	s[1] = 12 //こうすると、元の配列も12になる！

	fmt.Println(s)
	fmt.Println(len(s)) //2
	fmt.Println(cap(s)) // スライスの最大容量。２から切り出してるから、最後まで行ったら３つだから３


	// 配列から取り出したけど、いきなりスライスを作ることもできる！！

	s2 := make([]int, 3, 3) // lenとcapをしている（capは省略可）
	// ただ０で初期化するんじゃなくて、いきなり値を割り当てたい時は
	s3 := []int{1, 3, 5}

	s3 = append(s3, 8, 2, 10)

	fmt.Println(s3)


	t := make([]int, len(s))
	n := copy(t, s3)

	fmt.Println(s)
	fmt.Println(t)
	fmt.Println(n)
}
