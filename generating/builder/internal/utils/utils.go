package utils

import (
	"builder/internal/models"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"
)

func CreateFrontPage(data *models.FrontPageData, subject string) (string, error) {
	templateFrontPage, err := GetTemplateFromFile("templates/front_page.html")
	if err != nil {
		return "", err
	}
	frontPage := new(bytes.Buffer)
	err = templateFrontPage.Execute(frontPage, map[string]interface{}{
		"numberWork":     data.NumberWork,
		"subject":        subject,
		"studentGroup":   data.Student.Group,
		"studentName":    data.Student.Name,
		"studentSurname": data.Student.Surname,
		"teacherName":    data.Teacher.Name,
		"teacherSurname": data.Teacher.Surname,
	})
	if err != nil {
		return "", err
	}
	return frontPage.String(), nil
}

func CreateCommonPage(file string, data map[string]string) (string, error) {
	purposeTemplate, err := GetTemplateFromFile(file)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	purposeResult := new(bytes.Buffer)
	err = purposeTemplate.Execute(purposeResult, data)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return purposeResult.String(), nil
}

func GetTemplateFromFile(file string) (*template.Template, error) {
	templateBytes, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	templateString := string(templateBytes)
	templateName, _ := strings.CutSuffix(strings.ToTitle(file), ".html")
	readyTemplate, err := template.New(templateName).Parse(templateString)
	if err != nil {
		return nil, err
	}
	return readyTemplate, nil
}

func MergePages(pages []string, title string) (string, error) {
	mainTemplate, err := GetTemplateFromFile("templates/index.html")
	if err != nil {
		return "", err
	}
	result := new(bytes.Buffer)
	content := ""
	for _, page := range pages {
		content += page + "<br/><br/><br/><br/><br/>"
	}
	err = mainTemplate.Execute(result, map[string]string{
		"content": content,
		"subject": title,
	})
	if err != nil {
		return "", err
	}
	return result.String(), nil
}

func MoveFile(srcPath, dstPath string) error {
	inputFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	outputFile, err := os.Create(dstPath)
	if err != nil {
		inputFile.Close()
		return err
	}
	defer outputFile.Close()
	_, err = io.Copy(outputFile, inputFile)
	inputFile.Close()
	if err != nil {
		return err
	}
	return nil
}
