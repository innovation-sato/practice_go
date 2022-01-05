// インターフェース
// メソッドの一覧を定義したデータ型
package main

import (
	"fmt"
)

// show関数の引数を空のインターフェースにすると 　=> あらゆるデータ型を受け取ることができる
func show(t interface{}) {
	// データの型が何か知るための手法２つ
	//　1.型アサーション
	// t.とすると、値が二つ帰ってくる　　　　値とjapaneseかどうかの真偽値
	// _, ok := t.(japanese)
	// if ok {
	// 	fmt.Println("i am japanese")
	// } else {
	// 	fmt.Println("i am not japanese")
	// }

	// 2.型swith
	switch t.(type) {
	case japanese:
		fmt.Println("i am japanese")
	default:
		fmt.Println("i am not japanese")
	}
}

// 空のインターフェースも定義できる
// interface {}

// 挨拶をするインターフェース
type greeter interface {
	greet()
}

// 構造体を作る
type japanese struct{}
type american struct{}

func (j japanese) greet() {
	fmt.Println("こんにちわ〜")
}
func (a american) greet() {
	fmt.Println("Hello")
}

func main() {
	greeters := []greeter{japanese{}, american{}}
	for _, greeter := range greeters {
		greeter.greet()
		show(greeter)
	}
}
