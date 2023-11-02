package main

import "github.com/negrel/assert"

func main() {
	value := false
	assert.Truef(value, "bool is %v", value)
}
