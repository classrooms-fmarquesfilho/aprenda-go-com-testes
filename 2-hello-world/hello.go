package main

import "fmt"

const prefixoPortugues = "Olá, "

const (
	RegionRN = "rn"
	RegionSP = "sp"
	RegionMG = "mg"
	RegionRS = "rs"
)

func getRegionalVocative(region string) string {
	switch region {
	case RegionRN:
		return "boy"
	case RegionSP:
		return "mano"
	case RegionMG:
		return "sô"
	case RegionRS:
		return "tchê"
	default:
		return "" // sem vocativo regional
	}
}

func HelloWithRegion(hour int, region string) string {
	greeting := getTimeGreeting(hour)
	vocative := getRegionalVocative(region)

	if vocative != "" {
		return greeting + ", " + vocative
	}
	return greeting + ", Mundo"
}

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
