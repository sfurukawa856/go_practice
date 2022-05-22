package main

import (
	"fmt"
)

func main() {
	// 変数宣言一括
	var (
		i      int     = 0
		num    float64 = 100
		str    string  = "test"
		tr, fl bool    = true, false
	)
	// 出力一括
	fmt.Println(i, num, str, tr, fl)
}
