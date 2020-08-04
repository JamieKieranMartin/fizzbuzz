package main

import (
	"fmt"
)

var i = 1
var output = ""

func main() {
	for ; i <= 100; i++ {
		mod(3, "Fizz")
		mod(5, "Buzz")

		if output == "" {
			output = fmt.Sprintf("%v", i)
		}

		fmt.Println(output)
		output = ""
	}
}

func mod(mod int, s string) {
	if i%mod == 0 {
		output += s
	}
}
