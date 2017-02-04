package main

import (
	"bytes"
	"fmt"
	"os"
)

func comma(str string) string {
	var (
		lt          bool = false
		sz          int  = len(str)
		r           int  = sz % 3
		buffer, res bytes.Buffer
	)

	buffer.WriteString(str)

	if sz <= 3 {
		return str
	}
	for {

		if r == 0 {
			res.Write(buffer.Next(3))
			if lt {
				break
			}
			res.WriteByte(',')

		} else if r == 1 {
			res.Write(buffer.Next(1))
			res.WriteByte(',')

		} else if r == 2 {
			res.Write(buffer.Next(2))
			res.WriteByte(',')
		}

		r = buffer.Len() % 3
		if buffer.Len() == 3 {
			lt = true
		}
	}
	return res.String()
}

/*
func comma(s string) string {
	n := len(s)
	i := n % 3
	if i == 0 {
		i += 3
	}
	for i < n {
		s = s[:i] + "," + s[i:]
		i += 4
		n++
	}
	return s
}
*/
/*func main() {
	fmt.Println(comma("12"))
	fmt.Println(comma("123"))
	fmt.Println(comma("1234"))
	fmt.Println(comma("12345"))
	fmt.Println(comma("123456"))
	fmt.Println(comma("1234567"))
	fmt.Println(comma("12345678"))
	fmt.Println(comma("123456789"))
	fmt.Println(comma("1234567891"))
	fmt.Println(comma("12345678912"))
	fmt.Println(comma("123456789123"))
	fmt.Println(comma("1234567891234"))
}
*/
func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}
