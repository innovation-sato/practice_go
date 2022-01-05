package main

import (
	"fmt"
)

fuc main() {
	const (
		sun = iota //0 iotaで定数を宣言すると、それが0になって、自動的に下の定数の値が１、２と連番になる
		mon //1
		tue //2
	)
	fmt.Println(sun, mon, tue)s
}
