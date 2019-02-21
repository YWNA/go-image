package main

import (
	"./clip"
	"fmt"
)

func main() {
	TestClip()
}

func TestClip() {
	fileName := clip.Resize("./orange.png")
	fmt.Println(fileName)
}
