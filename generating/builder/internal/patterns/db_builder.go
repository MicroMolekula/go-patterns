package patterns

import (
	"builder/internal/models"
	"builder/internal/utils"
)

type DBBuilder struct {
	frontPage       string
	purpose         string
	exercise        string
	theoreticalInfo string
	resultWork      string
	analyzeResult   string
	conclusion      string
}

func NewDBBuilder() *DBBuilder {
	return &DBBuilder{}
}

func (db *DBBuilder) SetFrontPage(data models.FrontPageData) *DBBuilder {
	frontPage, err := utils.CreateFrontPage(&data, "Базы данных")
	if err != nil {
		frontPage = ""
	}
	db.frontPage = frontPage
	return db
}
