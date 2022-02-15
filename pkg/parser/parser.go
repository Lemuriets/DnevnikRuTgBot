package parser

type AbsParser interface {
	ParseMarks()
	GetLessonMarks()
	GetAllMarks()
	GetWeekShedule()
}

type Parser struct {
	Url string
}

func (p *Parser) ParseMarks() {

}

func (p *Parser) GetLessonMarks() {

}

func (p *Parser) GetAllMarks() {

}

func (p *Parser) GetWeekShedule() {

}

func (p *Parser) GetTomorrowShedule() {

}
