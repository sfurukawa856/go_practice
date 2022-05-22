package main

import (
	// importする順番をアルファベット順にすると読みやすいかも
	"fmt"
	"os/user"
	"time"
)

func main() {
	a := [1][2]string{{"foo", "bar"}}
	a[0][0] = "furukawa"

	fmt.Println(a)
	fmt.Println(time.Now())
	fmt.Println(user.Current())

}
