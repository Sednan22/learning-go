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
	/*--------------------------------------------*/

	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	var john = Employee{}

	raq := Employee{
		"raq",
		"test",
		235543543,
	}

	gus := Employee{
		firstName: "Gus",
		lastName:  "test",
		id:        13543,
	}

	fmt.Println(john)
	fmt.Println(raq)
	fmt.Println(gus)

	/*--------------------------------------------*/
	fmt.Println("Code finished...", time.Since(now))
	fmt.Println(strings.Repeat("-", 50))
}
