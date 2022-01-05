// 構造体
package main

import (
	"fmt"
)

type user struct {
	name  string
	score int
}

// メソッド
func (u user) show() {
	fmt.Printf("name:%s, score:%d\n", u.name, u.score)
}

// 参照渡しにすることで、実際の値の書き換えも可能
func (u *user) hit() {
	u.score++
}

func main() {
	// まず初期化
	u := new(user)
	u.name = "miyu"
	u.score = 20
	fmt.Println(u)

	u.hit()
	u.show()

	//　別の書き方色々！　こっちの書き方はポインタが帰ってくるわけじゃないから&つかない構造体のデータが入る
	u2 := user{"miyu", 50}
	fmt.Println(u2)
	u3 := user{name: "miyu", score: 60}
	fmt.Println(u3)
}
