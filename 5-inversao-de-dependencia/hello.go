package main

import (
	"fmt"
	"strings"
)

// HourNotSpecified indica que nenhuma hora foi fornecida.
// Horas válidas: [0, 23].
const HourNotSpecified = -1

const (
	RegionRN = "rn"
	RegionSP = "sp"
	RegionMG = "mg"
	RegionRS = "rs"
)

type HelloConfig struct {
	Name   string
	Hour   int
	Region string
}

type HelloOption func(*HelloConfig)

func WithName(name string) HelloOption {
	return func(cfg *HelloConfig) { cfg.Name = name }
}

func WithHour(hour int) HelloOption {
	return func(cfg *HelloConfig) { cfg.Hour = hour }
}

func WithRegion(region string) HelloOption {
	return func(cfg *HelloConfig) { cfg.Region = region }
}

// Hello retorna uma saudação personalizada.
// Use WithName, WithHour e WithRegion para customizar.
func Hello(options ...HelloOption) string {
	cfg := HelloConfig{
		Hour: HourNotSpecified,
	}

	for _, option := range options {
		option(&cfg)
	}

	cfg.Name = strings.TrimSpace(cfg.Name)

	var greeting string
	if cfg.Hour >= 0 && cfg.Hour <= 23 {
		greeting = getTimeGreeting(cfg.Hour)
	} else {
		greeting = "Olá"
	}

	vocative := getRegionalVocative(cfg.Region)

	if cfg.Name != "" {
		return greeting + ", " + cfg.Name
	} else if vocative != "" {
		return greeting + ", " + vocative
	}
	return greeting + ", Mundo"
}

func getTimeGreeting(hour int) string {
	switch {
	case hour >= 5 && hour < 12:
		return "Bom dia"
	case hour >= 12 && hour < 18:
		return "Boa tarde"
	default:
		return "Boa noite"
	}
}

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
		return ""
	}
}

func main() {
	fmt.Println(Hello())
	fmt.Println(Hello(WithName("Fernando")))
	fmt.Println(Hello(WithHour(0))) // Boa noite, Mundo
	fmt.Println(Hello(WithHour(9), WithName("Maria")))
	fmt.Println(Hello(WithRegion(RegionRN)))
	fmt.Println(Hello(
		WithName("João"),
		WithHour(20),
		WithRegion(RegionSP),
	))
}
