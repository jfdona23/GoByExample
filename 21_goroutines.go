package main

import (
	"fmt"
	"time"
)

func f(from string) {

	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	f("direct")

	go f("goroutine")

	go f("testing")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Duration(2 * time.Second))
	// time.Sleep(time.Second)
	fmt.Println("Done")
}
