package utils

import (
	"builder/internal/models"
	"bytes"
	"os"
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

func GetTemplateFromFile(file string) (*template.Template, error) {
	templateBytes, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	templateString := string(templateBytes)
	readyTemplate, err := template.New("FrontPage").Parse(templateString)
	if err != nil {
		return nil, err
	}
	return readyTemplate, nil
}
