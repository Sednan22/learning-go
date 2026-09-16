package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("Code starting...")
	var now = time.Now()

	x := []int{1, 2, 3, 4, 5, 6, 7, 8}
	y := make([]int, 3)

	fmt.Println(x)
	fmt.Println(y)

	copy(y, x)
	fmt.Println(y)

	fmt.Println("Code finished...", time.Since(now))
	fmt.Println(strings.Repeat("-", 50))
}
