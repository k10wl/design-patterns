package main

import "fmt"

func main() {
	fmt.Println(">>> Command demo")
	pasteController := PasteReceiver{}
	pasteCommand := PasteCommand{receiver: pasteController}
	button := Button{command: pasteCommand}
	shortcut := Shortcut{command: pasteCommand}
	button.execute()
	shortcut.execute()
}
