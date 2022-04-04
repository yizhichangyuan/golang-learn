package main

import "fmt"

// panic发生异常会停止runtime运行，但是加上defer延迟函数以及recover
// 可以从panic异常恢复并且应发panic异常函数不会继续运行但能正常返回
func fd() (result int) {
	defer func() {
		result = 8
		_ = recover()
	}()
	panic("panic")
}

func main() {
	fmt.Println(fd())
}
