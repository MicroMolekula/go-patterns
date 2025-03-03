package main

import (
	"composite_pattern/internal/composite"
	"log"
	"os"
)

func main() {
	var rootPath string
	switch len(os.Args) {
	case 2:
		rootPath = os.Args[1]
	default:
		rootPath = "."
	}
	root, err := composite.BuildTree(rootPath)
	if err != nil {
		log.Fatal("Ошибка чтения директории: ", err)
	}
	root.Print("", "", "")
}
