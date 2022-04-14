package ex116

var pc [256]byte

// init func invoked automatically to init complicated value
func init() {
	// range repeat ignore value only index for i, _ := range pc
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) (n int) {
	//return int(pc[x>>(0*8)] +
	//	pc[byte(x>>(1*8))] +
	//	pc[byte(x>>(2*8))] +
	//	pc[byte(x>>(3*8))] +
	//	pc[byte(x>>(4*8))] +
	//	pc[byte(x>>(5*8))] +
	//	pc[byte(x>>(6*8))] +
	//	pc[byte(x>>(7*8))])
	count := 0
	for x != 0 {
		count++
		x &= (x - 1)
	}
	return count
}
