package main

import "fmt"

func main() {
	channel := make(chan string,2)	
       channel <- "hello"
       channel <- "hi"
       fmt.Println(<-channel)
	   fmt.Println(<-channel)
	
}

