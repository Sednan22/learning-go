package main

import (
	"fmt"
)

const value = 20

func main() {

	var i int = 20
	var f float64 = float64(i)
	fmt.Println(i)
	fmt.Println(f)

	i = value
	f = value
	fmt.Println(i)
	fmt.Println(f)

	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615

	b += 1
	smallI += 1
	bigI += 1

	fmt.Println(b)
	fmt.Println(smallI)
	fmt.Println(bigI)

	testKeyboard("Testing keyboard...")

}

func testKeyboard(test string) {
	fmt.Println(test)
}
