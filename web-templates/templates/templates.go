package templates

import (
	"log"
	"os"
	"strings"
	"text/template"
	"time"
)

type sage struct {
	Name, Motto string
}

type car struct {
	Manufacturer, Model string
	Doors               int16
}

var tpl *template.Template
var tpl2 *template.Template
var tpl3 *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/tpl.gohtml"))
	tpl2 = template.Must(template.New("").Funcs(fm2).ParseFiles("templates/tpl2.gohtml"))
	tpl3 = template.Must(template.New("").Funcs(fm3).ParseFiles("templates/tpl3.gohtml"))
}

func Template() {
	//passingDataToTemplate()
	//passingFuncToTemplate()
	passingDateFormatToTemplate()
}

func getData() any {
	b := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs"}
	g := sage{
		Name:  "Gandhi",
		Motto: "Be the change"}
	m := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed."}

	f := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2}
	t := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4}

	sages := []sage{b, g, m}
	cars := []car{f, t}

	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		sages,
		cars,
	}

	return data
}

func passingDataToTemplate() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", getData())
	if err != nil {
		log.Fatalln(err)
	}
}

var fm2 = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

var fm3 = template.FuncMap{
	"fdateMDY": monthDayYear,
}

func monthDayYear(t time.Time) string {
	return t.Format("02-01-2006") // https://go.dev/src/time/format.go
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func passingFuncToTemplate() {
	err := tpl2.ExecuteTemplate(os.Stdout, "tpl2.gohtml", getData())
	if err != nil {
		log.Fatalln(err)
	}
}

func passingDateFormatToTemplate() {
	err := tpl3.ExecuteTemplate(os.Stdout, "tpl3.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
