package main

import (
	"github.com/justedd/hwgl-hw-1-1/internal/usecase/counter"
	"fmt"
)

func main() {
	counter := counter.NewCounter()
	top, err := counter.FileTop(5, "text.txt")

	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	for i := range top {
		line := fmt.Sprintf("%s: %d", top[i].Word, top[i].Count)
		fmt.Println(line)
	}
}
