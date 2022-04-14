package main

import "fmt"

var pc [256]byte

// init func invoked automatically to init complicated value
func init() {
	// range repeat ignore value only index for i, _ := range pc
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1) // 数字i的二进制中1的个数 = i右移一位后的个数 + i最低位是否为1
	}
}

func PopCount(x uint64) (n int) {
	for i := 0; i <= 8; i++ {
		n += int(pc[byte(x>>i*8)])
	}
	return n
	//return int(pc[x>>(0*8)] +
	//	pc[byte(x>>(1*8))] +
	//	pc[byte(x>>(2*8))] +
	//	pc[byte(x>>(3*8))] +
	//	pc[byte(x>>(4*8))] +
	//	pc[byte(x>>(5*8))] +
	//	pc[byte(x>>(6*8))] +
	//	pc[byte(x>>(7*8))])
}

var pc1 [256]byte = func() (pc1 [256]byte) {
	for i := range pc1 {
		pc1[i] = pc1[i/2] + byte(i&1)
	}
	return
}()

func PopCountByRightMostBit(x uint64) (n int) {
	for i := uint(0); i < 64; i++ {
		if x&1 != 0 {
			n++
		}
		x >>= 1
	}
	return
}

func PopCountByClearing(x uint64) (n int) {
	for x != 0 {
		x = x & (x - 1)
		n++
	}
	return
}

func main() {
	//fmt.Println(pc)  // already have init
	//fmt.Println(pc1) // already have init

	fmt.Println(PopCount(255))
}
