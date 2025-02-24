package patterns

import "builder/internal/models"

type Report string

type ReportBuilder interface {
	SetFrontPage(data models.FrontPageData) ReportBuilder
	SetPurpose(purpose string) ReportBuilder
	SetExercise(exercise string) ReportBuilder
	SetTheoreticalInfo(info string) ReportBuilder
	SetResultWork(result string) ReportBuilder
	SetAnalyzeResult(result string) ReportBuilder
	SetConclusion(conclusion string) ReportBuilder
	Build() Report
}
