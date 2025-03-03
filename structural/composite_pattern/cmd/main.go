package main

import (
	"composite_pattern/internal/composite"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("не достаточно аргументов")
	}
	rootPath := os.Args[1]
	root, err := composite.BuildTree(rootPath)
	if err != nil {
		log.Fatal("Ошибка чтения директории: ", err)
	}
	root.Print("", "", "")
}
