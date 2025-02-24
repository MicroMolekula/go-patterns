package patterns

import (
	"builder/internal/models"
	"builder/internal/utils"
)

type NetBuilder struct {
	frontPage       string
	purpose         string
	exercise        string
	theoreticalInfo string
	resultWork      string
	analyzeResult   string
	conclusion      string
}

func NewNetBuilder() *NetBuilder {
	return &NetBuilder{}
}

func (nb *NetBuilder) SetFrontPage(data models.FrontPageData) ReportBuilder {
	frontPage, err := utils.CreateFrontPage(&data, "Компьютерные сети")
	if err != nil {
		nb.frontPage = ""
		return nb
	}
	nb.frontPage = frontPage
	return nb
}

func (nb *NetBuilder) SetPurpose(purpose string) ReportBuilder {
	nb.purpose = CreatePage(purpose, "Цель работы")
	return nb
}

func (nb *NetBuilder) SetExercise(exercise string) ReportBuilder {
	nb.exercise = CreatePage(exercise, "Задание кафедры")
	return nb
}

func (nb *NetBuilder) SetTheoreticalInfo(info string) ReportBuilder {
	nb.theoreticalInfo = CreatePage(info, "Теоретические сведения")
	return nb
}

func (nb *NetBuilder) SetResultWork(result string) ReportBuilder {
	imageTemplate, err := utils.CreateCommonPage("templates/scheme.html", map[string]string{
		"image_path":         result,
		"description_scheme": "Сети",
	})
	if err != nil {
		nb.resultWork = ""
		return nb
	}
	resultTemplate, err := utils.CreateCommonPage("templates/result_work.html", map[string]string{
		"title_page": "Схема Сети",
		"data_page":  imageTemplate,
	})
	if err != nil {
		nb.resultWork = ""
		return nb
	}
	nb.resultWork = resultTemplate
	return nb
}

func (nb *NetBuilder) SetAnalyzeResult(result string) ReportBuilder {
	nb.analyzeResult = CreatePage(result, "Анализ работы")
	return nb
}

func (nb *NetBuilder) SetConclusion(conclusion string) ReportBuilder {
	nb.conclusion = CreatePage(conclusion, "Вывод")
	return nb
}

func (nb *NetBuilder) Build() Report {
	result, err := utils.MergePages([]string{
		nb.frontPage,
		nb.purpose,
		nb.exercise,
		nb.theoreticalInfo,
		nb.resultWork,
		nb.analyzeResult,
		nb.conclusion,
	}, "Компьютерным сетям")
	if err != nil {
		return ""
	}
	return Report(result)
}
