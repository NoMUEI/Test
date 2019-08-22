package main

import "fmt"

const helloPrefix = "Hello, "

func main() {
	fmt.Println(Hello("world",""))
}


func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return helloPrefix + name
}