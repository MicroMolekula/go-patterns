package main

import (
	"builder/internal/patterns"
	"errors"
	"fmt"
	"os"
)

func main() {
	var typeReport int
	fmt.Println("Программа по построению отчета")
	fmt.Print("1 - Базы данных\n2 - Компьютерные сети\n3 - Программирование\n> ")
	for _, err := fmt.Scan(&typeReport); err != nil; _, err = fmt.Scan(&typeReport) {
		fmt.Println(err)
		fmt.Print("1 - Базы данных\n2 - Компьютерные сети\n3 - Программирование\n> ")
	}
	var builder patterns.ReportBuilder
	switch typeReport {
	case 1:
		builder = patterns.NewDBBuilder()
	case 2:
		builder = patterns.NewNetBuilder()
	case 3:
		builder = patterns.NewProgBuilder()
	default:
		fmt.Println(errors.New("не правильный вариант"))
		return
	}
	director := patterns.NewDirector(builder)
	report := director.CreateFromTerminal()
	err := os.WriteFile("result/index.html", []byte(report), 0777)
	if err != nil {
		fmt.Println(err)
	}
}
