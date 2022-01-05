// range
// 配列やスライスマップと使える
package main

import (
	"fmt"
)

func main() {
	s := []int{2, 3, 8}

	// スライスの要素文だけ処理を繰り返したい時！
	for i, v := range s {
		fmt.Println(i, v)
	}

	// iはいらないって時は、アンダースコアをかく、要素だけ取り出せる！
	for _, v := range s {
		fmt.Println(v)
	}

	m := map[string]int{"miyu": 200, "sato": 300}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
