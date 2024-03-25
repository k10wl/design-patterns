package main

func main() {
	getSingleton()
	for i := 0; i < 3; i++ {
		getSingleton().Operation()
	}
	for i := 0; i < 3; i++ {
		go getSingleton().Operation()
	}
}
