package main

import "fmt"

const prefixoPortugues = "Olá, "

func Hello(name string) string {
	return prefixoPortugues + name
}

func main() {
	fmt.Println(Hello("mundo"))
}
