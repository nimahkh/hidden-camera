package main

import (
	"cardetector/capture"
	"cardetector/mail"
	"fmt"
	"github.com/eiannone/keyboard"
	"os"
)

func main() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	if len(os.Args) < 1 {
		fmt.Println("How to run:\n\tsaveimage [camera ID] [image file]")
		return
	}

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		fmt.Printf("You pressed: rune %q, key %X\r\n", char, key)
		capture.Doit()
		mail.Send()
		if key == keyboard.KeyEsc {
			break
		}
	}
}
