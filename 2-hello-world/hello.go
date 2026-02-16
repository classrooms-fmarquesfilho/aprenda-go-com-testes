package main

import (
	"fmt"
	"strings"
)

type HelloConfig struct {
	Name   string
	Hour   int
	Region string
}

type HelloOption func(*HelloConfig)

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

func WithName(name string) HelloOption {
	return func(cfg *HelloConfig) {
		cfg.Name = name
	}
}

func WithHour(hour int) HelloOption {
	return func(cfg *HelloConfig) {
		cfg.Hour = hour
	}
}

func WithRegion(region string) HelloOption {
	return func(cfg *HelloConfig) {
		cfg.Region = region
	}
}

func HelloWithConfig(cfg HelloConfig) string {
	if cfg.Name == "" {
		cfg.Name = "Mundo"
	}
	if cfg.Hour <= 0 && cfg.Hour > 24 {
		cfg.Hour = 12
	}

	greeting := getTimeGreeting(cfg.Hour)
	vocative := getRegionalVocative(cfg.Region)

	if vocative != "" {
		return greeting + ", " + vocative
	}
	return greeting + ", " + cfg.Name
}

func HelloWithRegion(hour int, region string) string {
	greeting := getTimeGreeting(hour)
	vocative := getRegionalVocative(region)

	if vocative != "" {
		return greeting + ", " + vocative
	}
	return greeting + ", Mundo"
}

func Hello(options ...HelloOption) string {
	cfg := HelloConfig{
		Name:   "",
		Hour:   0,
		Region: "",
	}

	for _, option := range options {
		option(&cfg)
	}
	// limpa espaços em branco
	cfg.Name = strings.TrimSpace(cfg.Name)

	var greeting string
	if cfg.Hour >= 1 && cfg.Hour <= 24 {
		greeting = getTimeGreeting(cfg.Hour)
	} else {
		greeting = "Olá"
	}

	vocative := getRegionalVocative(cfg.Region)

	// Lógica de prioridade:
	// 1. Se tem nome não vazio -> usa nome
	// 2. Se nome vazio mas tem vocativo -> usa vocativo
	// 3. Se ambos vazios -> usa "Mundo"

	if cfg.Name != "" {
		return greeting + ", " + cfg.Name
	} else if vocative != "" {
		return greeting + ", " + vocative
	} else {
		// ambos vazios
		return greeting + ", Mundo"
	}
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
	fmt.Println(Hello(WithName("mundo"), WithHour(14)))
}
