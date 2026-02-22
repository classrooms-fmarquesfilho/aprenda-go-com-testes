package main

import (
	"fmt"
	"io"
)

func Saudar(escritor io.Writer, options ...HelloOption) {
	fmt.Fprintf(escritor, "%s", Hello(options...))
}
