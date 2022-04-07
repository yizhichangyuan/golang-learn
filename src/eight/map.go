package main

import (
	"fmt"
	"sync"
	"time"
)

const timeFormat = "2006-01-02 15:04:05.999999999 -0700 MST"

func main() {
	var m sync.Map
	m.Store(1, time.Now())
	v, _ := m.Load(1)
	fmt.Println(v.(time.Time))
}
