// Berkay Deniz
// Exercise 2.1, 2.2, 2.3

package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// Exercise 2.3
func PopCount2(x uint64) int {
	sum := 0
	for i := 0; i < 8; i++ {
		sum += int(pc[byte(x>>uint(i))])
	}
	return sum
}

// Exercise 2.4
func PopCount3(x uint64) int {
	count := 0
	y := uint64(1)
	for i := 0; i < 64; i++ {
		if x&y > 0 {
			count += 1
		}
		x >>= 1
	}
	return count
}

// Exercise 2.5
func PopCount4(x uint64) int {
	count := 0
	for x != 0 {
		x &= x - 1
		count += 1
	}
	return count
}
