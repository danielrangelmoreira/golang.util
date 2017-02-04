package popcount

/*
import (
	"fmt"
)*/

var pc [256]byte

func init() {

	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)

	}

}
func ShaCompare(chksum1, chksum2 [32]byte) int {
	var count int
	for i, v := range chksum1 {
		count += PopCount(uint64(chksum2[i] | v))
	}
	return count
}

func BitCount(x uint64) int {
	var count int
	for count = 0; x > 0; count++ {
		x &= (x - 1)
	}
	return count
}

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

/*
func main() {

	fmt.Printf("% 0x\n", pc)
	fmt.Printf("bitCount %b: %d\n", 100, bitCount(100))
	//fmt.Printf("PopCount %b: %d\n", 100, PopCount(100))
	//fmt.Printf("PopCount %b: %d\n", 333, PopCount(333))
	//fmt.Printf("PopCount %b: %d\n", 5, PopCount(5))
	//	fmt.Printf("pc[%d]: %b \n", i, pc[i])
	//fmt.Printf("byte(10&1): %b\n", byte(i&1))

}
*/
