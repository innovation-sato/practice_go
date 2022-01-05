/*チャネル
goroutineの時、各メソッドで受け取った値をmain関数で使いたいみたいになった時
return で普通にはできない、その時使うのがchannel
データの受け渡しをするパイプ見たいなイメージ
*/

package main

import (
	"fmt"
	"time"
)

func task1(result chan string) {
	// 指定した時間だけ処理待ちをしてくれるというもの
	time.Sleep(time.Second * 2)

	// チャネルに値を渡す
	result <- "task1 result"
}

func task2() {
	fmt.Println("task2 finished!!")
}

func main() {
	// channel作成
	result := make(chan string)

	// goとつけるとgoroutine
	go task1(result)
	go task2()

	fmt.Println(<-result)

	time.Sleep(time.Second * 3)

}
