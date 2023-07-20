package main

import (
	"fmt"
	"time"
)

func f(s string) {
	for i := 0; i <= 100; i++ {
		fmt.Println(s, ":", i)
	}
}
func main() {
	f("ping")
	go f("pong")
	go func (msg string) {
		fmt.Println("tick")
	}("tick")
	time.Sleep(2 * time.Second)
	fmt.Println("completed")
}
