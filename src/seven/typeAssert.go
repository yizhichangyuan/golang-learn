package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout
	fmt.Printf("%T\n", w)
	// 检查w动态类型是否是类型T，结果返回对象os.Stdout，类型就是os.File, 可以获得os.File的方法
	if w, ok := w.(*os.File); ok {
		w.WriteString("abc")
	}
	//fmt.Println(ok)
	// 检查w动态类型是否满足接口类型T，返回具备相同断言类型的类型和值部分的接口值，本例中就是变成了io.ReadWriter接口
	s, _ := w.(io.ReadWriter)
	s.Read([]byte("abc"))
	_, ok2 := w.(*bytes.Buffer)
	fmt.Println(ok2)
}
