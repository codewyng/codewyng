package main

import "fmt"

// Does this program leak the private key? 🤔
// How much harder would it be to determine that without CodeWyng?

// Try these features:
// - Hover documentation
// - Jump-to-definition
// - Find-references
// - Quick search with Cmd+/ (or Win+/)

type Public struct { key string }
type Private struct {
	// Is the private key ever leaked? Click to find references! 🔍
	key string
}

var key = Private{key: "abcsecret!"}
var pub = Public{key: "abcpublic"}

func main() {
	initialize() // What does this do? Hover to see documentation 📜
	fmt.Println(getLength(pub))
}

// checkLength returns whether or not the key is 64 characters long.
func checkLength(key string) bool {
	return len(key) == 64
}

// getLength returns the length of the key, or -1 if invalid.
func getLength(pub Public) int {
	if checkLength(pub.key) {
		key := pub.key
		return len(key)
	} else {
		// Does this leak the private key? Which key is it?
		fmt.Println(key.key)
		return -1
	}
}

// scroll
//   |
//   |
//   v










// a bit further...














// init panics if the public or private key are empty.
//
// Click on a call site to jump to it.
// It's defined at the bottom of the file to demo jump-to-definition.
func initialize() {
	if len(pub.key) == 0 || len(key.key) == 0 {
		panic("empty!")
	}
}
