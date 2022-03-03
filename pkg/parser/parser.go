package parser

import (
	"github.com/geziyor/geziyor"
)

type AbsParser interface {
	ParseMarks()
	GetLessonMarks()
	GetAllMarks()
	GetWeekShedule()
	GetTomorrowShedule()
}

type Parser struct {
	Parser *geziyor.Geziyor
	Url    string
}

func NewParser(url string) *Parser {
	par := geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{url},
	})

	return &Parser{
		Parser: par,
		Url:    url,
	}
}
