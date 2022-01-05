package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	name := getName()
	fmt.Printf("ようこそ!%sさん!まもなくゲームが始まります...\n", name)

	// カウントダウン
	for i := 5; i < 10; i-- {
		if i == 0 {
			fmt.Println("START!!")
			break
		}
		countdown(&i)
	}

	// ゲーム実行
	totalScore := 0
	ask(1, "apple", &totalScore)
	ask(2, "melon", &totalScore)
	ask(3, "GoodMorning", &totalScore)

	// クイズの結果表示
	result(name, totalScore)

	// サーバー起動、おみくじ
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

// 名前を登録
func getName() string {
	var name string
	fmt.Println("ニックネームをを入力してね！")

	// 入力されたニックネームを取得
	fmt.Scan(&name)
	return name
}

// 開始前のカウントダウンを行う関数
func countdown(numPtr *int) {
	fmt.Println(*numPtr)
	time.Sleep(time.Second * 1)
}

// ゲームを行う関数
func ask(number int, question string, scorePtr *int) {
	var input string
	fmt.Printf("[質問%d]次の単語を入力せよ!: %s\n", number, question)

	fmt.Scan(&input)

	if question == input {
		fmt.Println("正解！")
		*scorePtr += 10
	} else {
		fmt.Println("不正解！")
	}
}

// ゲームの結果を表示する関数
func result(name string, totalScore int) {
	fmt.Printf("%sさんの結果は...\n", name)
	time.Sleep(time.Second * 3)

	fmt.Printf("スコア：%d点 (30点満点中)\n", totalScore)
	fmt.Println("お疲れさまです！下記URLにアクセスしておみくじに挑戦してみてね！")
	fmt.Printf("http://127.0.0.1:8080/%v\n", name)
}

// サーバー起動
// おみくじの結果表示
func handler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	num := rand.Intn(6)
	var omikuzi string

	switch num {
	case 0, 1:
		omikuzi = "大吉！おめでとう！！"
	case 2, 3:
		omikuzi = "中吉"
	case 4:
		omikuzi = "小吉"
	case 5:
		omikuzi = "すえきち"
	}

	fmt.Fprintf(w, "Hi %sさん!\n", r.URL.Path[1:])
	w.Write([]byte("あなたの今日の運勢は..."))
	w.Write([]byte(omikuzi))
	w.Write([]byte("\nまたチャレンジしてね！"))
}
