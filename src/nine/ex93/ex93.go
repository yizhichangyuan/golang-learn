package main

import (
	"flag"
	"os"
	"sync"
)

var mu sync.Mutex
var stopFlag = flag.Bool("done", false, "stop getting key from map manually by inputting any words")
var done = make(chan struct{})

type Memo struct {
	f     Func
	cache map[string]*result // 对应value为指针，便于更新值而不再需要再次访问该map
}

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
	ready chan struct{} // 表示该值是否已经计算完成的channel
}

func New(f Func) *Memo {
	return &Memo{f, make(map[string]*result)}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	if *stopFlag {
		select {
		case <-done:
			return nil, nil
		default:
		}
	}
	mu.Lock()
	res := memo.cache[key]
	if res == nil {
		e := result{ready: make(chan struct{})}
		memo.cache[key] = &e         // 直接先放入一个空调目
		mu.Unlock()                  // tips亮点：锁直接在这儿释放，不需要在计算过程中仍然维持该锁，后续直接根据指针将值赋上而不再需要访问map
		e.value, e.err = memo.f(key) // 计算耗时，其他goroutines也都在<-res.ready进行广播信号等待
		close(e.ready)
	} else {
		mu.Unlock()
		<-res.ready // 除了第一个goroutines之外，大量goroutines直接等待计算完成发出的广播信号
	}
	return res.value, res.err
}

func main() {
	flag.Parse()
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()
}
