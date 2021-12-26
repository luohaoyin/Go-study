package main

import (
	"fmt"
	"sync"
)
/*
报错原因：main函数里面，声明一个变量后开启子协程，子协程需要完成输出后并执行lock操作；
        与此同时，主协程继续向下执行unlock操作。
        主协程完成操作后，子协程还没完成就退出了，即unlock了一个未lock的变量，因此报错。
*/
func main() {
	var mu sync.Mutex

	go func() {
		fmt.Println("有点强人锁男")
		mu.Lock()
	}()

	mu.Unlock()
}