package main

import "sync"

var mu sync.RWMutex

type Memo struct {
	f     Func
	cache map[string]result
}

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f, make(map[string]result)}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	// 读锁，保证大量读可以并发而不需要顺序抢占sync.Mutex排他锁
	// f计算耗时，加上了写锁，但是一旦完成大量读锁就可以并发获取计算好的值而不需要挨个抢占排他锁
	mu.RLock()
	res, ok := memo.cache[key]
	mu.RUnlock()
	if !ok {
		// 写
		mu.Lock()
		res, ok = memo.cache[key]
		if !ok {
			res.value, res.err = memo.f(key)
			memo.cache[key] = res
		}
		mu.Unlock()
	}
	return res.value, res.err
}
