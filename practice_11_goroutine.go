// goroutine
// 並行処理を簡単に描けるようになる (２つの処理を同時に走らせる）)
package main

import (
	"fmt"
	"time"
)

// 重い処理
func task1() {
	// 指定した時間だけ処理待ちをしてくれるというもの
	time.Sleep(time.Second * 2)
	fmt.Println("task1 finished!!")
}

// 軽い処理
func task2() {
	fmt.Println("task2 finished!!")
}

func main() {
	// goとつけるとgoroutine
	go task1()
	go task2()

	time.Sleep(time.Second * 6)
}
