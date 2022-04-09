package memo6

type Memo struct {
	f        Func
	requests chan request
}

type request struct {
	key      string
	response chan result
}

type result struct {
	value interface{}
	err   error
	ready chan struct{}
}

type Func func(key string) (interface{}, error)

func (memo *Memo) server() {
	cache := make(map[string]*result) // 指针用于在另一个goroutines cal更新而不需要访问map，并且respond也不需要访问map
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			e = &result{ready: make(chan struct{})}
			cache[req.key] = e
			go memo.cal(e, req.key)
		}
		go respond(req, e)
	}
}

func (memo *Memo) cal(res *result, key string) {
	res.value, res.err = memo.f(key)
	close(res.ready)
}

func respond(r request, res *result) {
	<-res.ready
	r.response <- *res
}

func (memo *Memo) Get(key string) result {
	resp := make(chan result)
	memo.requests <- request{key, resp}
	return <-resp
}
