package main

import "sync"

var pc [256]byte
var once sync.Once

// init func invoked automatically to init complicated value
func load() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) (n int) {
	once.Do(load)
	return int(pc[x>>(0*8)] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))] +
		pc[byte(x>>(8*8))])
}
