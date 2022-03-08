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
	Body   string
}

func NewParser(body string) *Parser {
	par := geziyor.NewGeziyor(&geziyor.Options{})

	return &Parser{
		Parser: par,
		Body:   body,
	}
}
