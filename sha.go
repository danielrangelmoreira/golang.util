package main

import (
	"crypto/sha256"
	//"crypto/sha512"
	"flag"
	"fmt"
	"hash"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

var pc [256]byte

func init() {

	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)

	}

}
func ShaCompare(chksum1, chksum2 []byte) int {
	var count int
	for i, v := range chksum1 {
		count += PopCount(uint64(chksum2[i] ^ v))
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

type intFlag struct {
	set   bool
	value int
}

type stringFlag struct {
	set   bool
	value string
}

func (self *intFlag) Set(str string) error {
	i, _ := strconv.Atoi(str)
	self.value = i
	self.set = true
	return nil
}

func (self *intFlag) String() string {
	return string(self.value)
}
func (self *stringFlag) Set(str string) error {
	self.value = str
	self.set = true
	return nil
}

func (self *stringFlag) String() string {
	return self.value
}

var (
	filename1 stringFlag
	filename2 stringFlag
	algorithm intFlag
)

func init() {
	flag.Var(&filename1, "f", "filename")
	flag.Var(&filename2, "c", "compare with this file hash")
	flag.Var(&algorithm, "a", "algorithm to be used 256, 384 or 512")
}

func main() {

	var myhash []hash.Hash
	flag.Parse()

	if filename1.set {
		h := sha256.New()
		b, err := ioutil.ReadFile(filename1.value)
		if err != nil {
			log.Panic("Error opening file1")
		}
		h.Write(b)
		myhash = append(myhash, h)

	} else {
		h := sha256.New()
		b, _ := ioutil.ReadAll(os.Stdin)
		h.Write(b)
		myhash = append(myhash, h)

	}
	if filename2.set {
		h := sha256.New()
		b, err := ioutil.ReadFile(filename2.value)
		if err != nil {
			log.Panic("Error opening file2")
		}
		h.Write(b)
		myhash = append(myhash, h)

	}
	/*	if !algorithm.set {
			myhash = sha256.New()
		} else if algorithm.value == 256 {
			myhash = sha256.New()
		} else if algorithm.value == 384 {
			myhash = sha512.New384()
		} else if algorithm.value == 512 {
			myhash = sha512.New()
		} else {
			log.Panic("Unrecognized algorithm")
		}*/

	for _, h := range myhash {
		fmt.Printf("%x\n", h.Sum(nil))
	}

	//fmt.Printf("%x\n %[1]t\n", myhash[0].Sum(nil))
	fmt.Printf("number of different bits: %d\n", ShaCompare(myhash[0].Sum(nil), myhash[1].Sum(nil)))
	/*fmt.Printf("%t\n", filename2.set)

	if filename1.set {
		fmt.Printf("%x\n", myhash[0].Sum(nil))

	} else {
		fmt.Printf("%x  %s\n", myhash[0].Sum(nil), "-")
	}

	if filename2.set {
		fmt.Printf("%x\n", myhash[1].Sum(nil))
	}*/

}
