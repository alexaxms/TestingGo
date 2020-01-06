package main

import(
	"fmt"
)
func main() {
	fmt.Println(HelloWorld("Alex"))
}

const HELLO = "Hello, "
func HelloWorld(name string) string {
	if name == "" {
		name = "world"
	}
	return HELLO + name
}