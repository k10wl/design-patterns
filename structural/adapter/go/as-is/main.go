package main

import "fmt"

type Target interface {
	Operation()
}

func main() {
	fmt.Println("Adapter demo:")
	adapter := NewAdapter(&Adaptee{})
	adapter.Operation()
}
