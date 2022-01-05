// for文　繰り返し
package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		if i == 3 {
			break
		}

		if i == 1 {
			continue
		}
		fmt.Println(i)
	}

	// for文の省略
	fmt.Println("while文風にこうもかける")
	i2 := 0
	for i2 < 10 {
		fmt.Println(i2)
		i2++
	}

	// 無限ループ
	i3 := 0
	fmt.Println("無限ループ")
	for {
		fmt.Println(i3)
		i3++
		if i3 == 3 {
			break
		}

	}

}
