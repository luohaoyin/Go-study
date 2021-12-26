package main

import "fmt"

func main() {
	ch := make(chan int,1)
	ch <- 1
	go func() {
		fmt.Println("下山的路又堵起了")
		<- ch
	}()
	ch <- 2
}