package main

import "fmt"

// switch 的坑

func main() {
	f := func() bool {
		return false
	}

	switch _ = f(); {
	// switch {
	case false:
		fmt.Println("假")
	case true:
		fmt.Println("真")
	}
}
