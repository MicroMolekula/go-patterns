package patterns

import (
	"bufio"
	"builder/internal/models"
	"builder/internal/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Director struct {
	builder ReportBuilder
}

func NewDirector(builder ReportBuilder) *Director {
	return &Director{
		builder: builder,
	}
}

func (d *Director) CreateReport(
	dataFrontPage models.FrontPageData,
	purpose string,
	exercise string,
	theoreticalInfo string,
	result string,
	analyzeResult string,
	conclusion string) Report {
	d.builder.SetFrontPage(dataFrontPage).SetPurpose(purpose).SetExercise(exercise).SetTheoreticalInfo(theoreticalInfo)
	d.builder.SetResultWork(result).SetAnalyzeResult(analyzeResult).SetConclusion(conclusion)
	return d.builder.Build()
}

func (d *Director) CreateFromTerminal() Report {
	frontPage := getDataFrontPage()
	purpose := getPurpose()
	exercise := getExercise()
	theoreticalInfo := getTheoreticalInfo()
	result := getResult(d.builder)
	analyzeResult := getAnalyzeResult()
	conclusion := getConclusion()
	return d.CreateReport(
		frontPage,
		purpose,
		exercise,
		theoreticalInfo,
		result,
		analyzeResult,
		conclusion,
	)
}

func getDataFrontPage() models.FrontPageData {
	data := models.FrontPageData{}
	fmt.Println("Ввод данных для титульного листа")
	fmt.Println("Фамилия Студента: ")
	data.Student.Surname = readString()
	fmt.Println("Инициалы Студента: ")
	data.Student.Name = readString()
	fmt.Println("Группа студента: ")
	data.Student.Group = readString()
	fmt.Println("Фамилия Преподователя: ")
	data.Teacher.Surname = readString()
	fmt.Println("Инициалы Студента: ")
	data.Teacher.Name = readString()
	fmt.Println("Номер лабораторной работы: ")
	numberWork, err := strconv.Atoi(readString())
	if err != nil {
		fmt.Println(err)
		data.NumberWork = 1
		return data
	}
	data.NumberWork = numberWork
	return data
}

func getPurpose() string {
	fmt.Println("Ввод данных для Цели работы")
	fmt.Println("Цель работы: ")
	return readString()
}

func getExercise() string {
	fmt.Println("Ввод данных для Задания кафедры")
	fmt.Println("Задание кафедры: ")
	return readString()
}

func getTheoreticalInfo() string {
	fmt.Println("Ввод данных для Теоретических сведений")
	fmt.Println("Теоретические сведения: ")
	return readString()
}

func getAnalyzeResult() string {
	fmt.Println("Ввод данных для Анализа работы")
	fmt.Println("Анализ работы: ")
	return readString()
}

func getConclusion() string {
	fmt.Println("Ввод данных для Вывода")
	fmt.Println("Вывод: ")
	return readString()
}

func getResult(builder ReportBuilder) string {
	var result string
	fmt.Println("Ввод результата работы")
	switch builder.(type) {
	case *DBBuilder:
		return getScheme("Базы данных")
	case *NetBuilder:
		return getScheme("Компьютерной сети")
	case *ProgBuilder:
		return getCode()
	}
	return result
}

func getScheme(subject string) string {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Ошибка")
		}
	}()
	var scheme string
	var path string
	fmt.Println(fmt.Sprintf("Ввод схемы %s", subject))
	fmt.Println("Полный путь до изображения: ")
	path = readString()
	suffixImage := strings.Split(path, ".")[1]
	scheme = "images/result." + suffixImage
	err := utils.MoveFile(path, "result/"+scheme)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Ошибка чтения файла")
	}
	return scheme
}

func getCode() string {
	fmt.Println(fmt.Sprintf("Ввод кода программы"))
	fmt.Println("Путь к файлу с кодом программы: ")
	path := readString()
	code, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	return string(code)
}

func readString() string {
	reader := bufio.NewReader(os.Stdin)
	result, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	text, _ := strings.CutSuffix(result, "\n")
	return text
}
