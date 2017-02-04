//teste do sha
package main

import (
	"fmt"
	"strings"
)

func IsAnag(str1, str2 string) bool {

	if len(str1) != len(str2) {
		return false
	}

	for _, char := range strings.ToLower(str1) {
		if strings.ContainsRune(strings.ToLower(str2), char) {
			continue

		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Printf("Hello e olleH: %v\n", IsAnag("Hello", "olleh"))
	fmt.Printf("Daniel e leinaD: %v\n", IsAnag("Daniel", "leinaD"))
	fmt.Printf("Amor e Roma: %v\n", IsAnag("Amor", "Roma"))
	fmt.Printf("Rato e otar: %v\n", IsAnag("Rato", "otar"))
	fmt.Printf("Hello e olleH : %v\n", IsAnag("Hello", "olleH "))
}
