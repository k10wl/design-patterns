package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(">>> Facade demo")
	logger := NewLogger(os.Stdin)
	logger.Log("That's all, Folks!")
}
