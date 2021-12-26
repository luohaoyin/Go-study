package main

import (
	"fmt"
	"time"
)
/*
**基于管道**，我们可以把打印的协程拓展为N个。

请在``main``函数中开启10个协程输出一段话，要求10行话全部输出完毕后再结束``main``函数。
*/
func hello(){
	fmt.Println("hello,world.")
}


func main(){
   for i:=1;i<=10;i++{
	   go hello()
   }
   time.Sleep(time.Second)
}
