package patterns

import "builder/internal/utils"

func CreatePage(data, title string) string {
	templateResult, err := utils.CreateCommonPage("templates/common_page.html", map[string]string{
		"title_page": title,
		"data_page":  data,
	})
	if err != nil {
		return ""
	}
	return templateResult
}
