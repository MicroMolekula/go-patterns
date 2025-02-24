package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	templateBytes, err := os.ReadFile("templates/index.html")
	if err != nil {
		log.Fatal(err)
	}
	templateString := string(templateBytes)
	readyTemplate, err := template.New("test").Parse(templateString)
	if err != nil {
		log.Fatal(err)
	}
	err = readyTemplate.Execute(os.Stdout, map[string]string{
		"title": "Hello",
	})
	if err != nil {
		log.Fatal(err)
	}

}
