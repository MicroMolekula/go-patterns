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

func (db *DBBuilder) SetFrontPage(data models.FrontPageData) ReportBuilder {
	frontPage, err := utils.CreateFrontPage(&data, "Базы данных")
	if err != nil {
		db.frontPage = ""
		return db
	}
	db.frontPage = frontPage
	return db
}

func (db *DBBuilder) SetPurpose(purpose string) ReportBuilder {
	db.purpose = CreatePage(purpose, "Цель работы")
	return db
}

func (db *DBBuilder) SetExercise(exercise string) ReportBuilder {
	db.exercise = CreatePage(exercise, "Задание кафедры")
	return db
}

func (db *DBBuilder) SetTheoreticalInfo(info string) ReportBuilder {
	db.theoreticalInfo = CreatePage(info, "Теоретические сведения")
	return db
}

func (db *DBBuilder) SetResultWork(result string) ReportBuilder {
	imageTemplate, err := utils.CreateCommonPage("templates/scheme.html", map[string]string{
		"image_path":         result,
		"description_scheme": "Базы данных",
	})
	if err != nil {
		db.resultWork = ""
		return db
	}
	resultTemplate, err := utils.CreateCommonPage("templates/result_work.html", map[string]string{
		"title_page": "Схема базы данных",
		"data_page":  imageTemplate,
	})
	if err != nil {
		db.resultWork = ""
		return db
	}
	db.resultWork = resultTemplate
	return db
}

func (db *DBBuilder) SetAnalyzeResult(result string) ReportBuilder {
	db.analyzeResult = CreatePage(result, "Анализ работы")
	return db
}

func (db *DBBuilder) SetConclusion(conclusion string) ReportBuilder {
	db.conclusion = CreatePage(conclusion, "Вывод")
	return db
}

func (db *DBBuilder) Build() Report {
	result, err := utils.MergePages([]string{
		db.frontPage,
		db.purpose,
		db.exercise,
		db.theoreticalInfo,
		db.resultWork,
		db.analyzeResult,
		db.conclusion,
	}, "Базам данных")
	if err != nil {
		return ""
	}
	return Report(result)
}
