package patterns

import (
	"builder/internal/models"
	"builder/internal/utils"
)

type ProgBuilder struct {
	frontPage       string
	purpose         string
	exercise        string
	theoreticalInfo string
	resultWork      string
	analyzeResult   string
	conclusion      string
}

func NewProgBuilder() *ProgBuilder {
	return &ProgBuilder{}
}

func (pb *ProgBuilder) SetFrontPage(data models.FrontPageData) ReportBuilder {
	frontPage, err := utils.CreateFrontPage(&data, "Программирование")
	if err != nil {
		pb.frontPage = ""
		return pb
	}
	pb.frontPage = frontPage
	return pb
}

func (pb *ProgBuilder) SetPurpose(purpose string) ReportBuilder {
	pb.purpose = CreatePage(purpose, "Цель работы")
	return pb
}

func (pb *ProgBuilder) SetExercise(exercise string) ReportBuilder {
	pb.exercise = CreatePage(exercise, "Задание кафедры")
	return pb
}

func (pb *ProgBuilder) SetTheoreticalInfo(info string) ReportBuilder {
	pb.theoreticalInfo = CreatePage(info, "Теоретические сведения")
	return pb
}

func (pb *ProgBuilder) SetResultWork(result string) ReportBuilder {
	codeTemplate, err := utils.CreateCommonPage("templates/code.html", map[string]string{
		"code": result,
	})
	if err != nil {
		pb.resultWork = ""
		return pb
	}
	resultTemplate, err := utils.CreateCommonPage("templates/result_work.html", map[string]string{
		"title_page": "Код программы",
		"data_page":  codeTemplate,
	})
	if err != nil {
		pb.resultWork = ""
		return pb
	}
	pb.resultWork = resultTemplate
	return pb
}

func (pb *ProgBuilder) SetAnalyzeResult(result string) ReportBuilder {
	pb.analyzeResult = CreatePage(result, "Анализ работы")
	return pb
}

func (pb *ProgBuilder) SetConclusion(conclusion string) ReportBuilder {
	pb.conclusion = CreatePage(conclusion, "Вывод")
	return pb
}

func (pb *ProgBuilder) Build() Report {
	result, err := utils.MergePages([]string{
		pb.frontPage,
		pb.purpose,
		pb.exercise,
		pb.theoreticalInfo,
		pb.resultWork,
		pb.analyzeResult,
		pb.conclusion,
	}, "Программированию")
	if err != nil {
		return ""
	}
	return Report(result)
}
