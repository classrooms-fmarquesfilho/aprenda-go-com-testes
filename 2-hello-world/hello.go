package main

import "fmt"

const prefixoPortugues = "Olá, "

func Hello(name string) string {
	if name == "" {
		name = "Mundo"
	}
	return prefixoPortugues + name
}

func getTimeGreeting(hour int) string {
	if hour >= 5 && hour < 12 {
		return "Bom dia"
	} else if hour >= 12 && hour < 18 {
		return "Boa tarde"
	}
	return "Boa noite"
}

func HelloWithTime(name string, hour int) string {
	if name == "" {
		name = "Mundo"
	}

	return getTimeGreeting(hour) + ", " + name
}

func main() {
	fmt.Println(Hello("mundo"))
}
