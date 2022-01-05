// マップ
package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["miyu"] = 200
	m["sato"] = 300

	// 宣言しながら値を入れるには
	m2 := map[string]int{"miyu2": 100, "sato2": 200}

	fmt.Println(m)
	fmt.Println(m2)
	fmt.Println(len(m2))

	//要素を削除したいときは
	delete(m2, "miyu2")
	fmt.Println(m2)

	//キーの存在を調べるには
	v, ok := m2["sato2"]
	fmt.Println(v)  // 値
	fmt.Println(ok) // 存在してればtrue しなければfalse
}
