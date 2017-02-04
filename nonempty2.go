package main

import (
	"fmt"
)

func nonempty2(strings []string) []string {
	var szStr []string
	for _, str := range strings {
		if str != "" {
			szStr = append(szStr, str)
		}
	}
	return szStr
}

func main() {
	str := []string{"daniel", "", "", "", "rangel"}
	fmt.Printf("%s\n", nonempty2(str))
	fmt.Printf("%s\n", str)

}

/*package main

import (
	"fmt"
)

func nonempty(strings []string) []string {
	i := 0
	for _, str := range strings {
		if str != "" {
			strings[i] = str
			i++
		}
	}
	return strings[:i]
}

func main() {
	str := []string{"daniel", "", "rangel"}
	fmt.Printf("%s\n", nonempty(str))
	fmt.Printf("%s\n", str)

}*/
