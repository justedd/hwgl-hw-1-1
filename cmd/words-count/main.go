package main

import (
	"github.com/justedd/hwgl-hw-1-1/internal/usecase/counter"
	"fmt"
)

func main() {
	counter := counter.NewCounter()
	top := counter.FileTop(5, "text.txt")

	for i := range top {
		line := fmt.Sprintf("%s: %d", top[i].Word, top[i].Count)
		fmt.Println(line)
	}
}
