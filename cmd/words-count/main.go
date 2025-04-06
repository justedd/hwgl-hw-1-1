package main

import (
	"fmt"

	"github.com/justedd/hwgl-hw-1-1/internal/usecase/counter"
)

func main() {
	counter, err := counter.NewCounter()
	if err != nil {
		fmt.Printf("Initialization error: %v", err)
		return
	}

	top, err := counter.FileTop(5, "text.txt")

	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	for i := range top {
		line := fmt.Sprintf("%s: %d", top[i].Word, top[i].Count)
		fmt.Println(line)
	}
}
