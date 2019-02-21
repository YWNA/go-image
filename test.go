package main

import (
	"./clip"
	"fmt"
)

func main() {
	TestClip()
}

func TestClip() {
	fileName := clip.Resize("./orange.png", 0, 0, 200, 200)
	//fileName := clip.Resize("./pineapple.jpeg", 0, 0, 200, 200)
	fmt.Println(fileName)
}
