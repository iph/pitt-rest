package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Course struct {
	Subject       string
	CatalogNumber int
	ClassNumber   int
	Title         string
	Instructor    string
	Credits       string
	Term          Term
}

type Term struct {
	Semester string
	Year     int
	ID       int
}

func NewTerm(termId string) Term {
	semester := ""
	if termId[3] == '1' {
		semester = "Fall"
	} else if termId[3] == '4' {
		semester = "Spring"
	} else if termId[3] == '7' {
		semester = "Summer"
	}
	// This code will last 100 years I say!!
	year, _ := strconv.Atoi("20" + string(termId[1]) + string(termId[2]))
	termNum, _ := strconv.Atoi(termId)
	return Term{Semester: semester, Year: year, ID: termNum}
}

// Ugliness. RIP Go parsing. TODO - clean this up.
func ParseCourses(term string, courseType string) []Course {
	courses := []Course{}
	url := fmt.Sprintf("http://www.courses.as.pitt.edu/results-subja.asp?TERM=%s&SUBJ=%s", term, courseType)
	doc, _ := goquery.NewDocument(url)
	doc.Find("tr.odd").Each(func(i int, s *goquery.Selection) {
		course := ParseCourse(s)
		courses = append(courses, course)
	})
	doc.Find("tr.even").Each(func(i int, s *goquery.Selection) {
		course := ParseCourse(s)
		courses = append(courses, course)
	})

	return courses
}

func ParseCourse(s *goquery.Selection) Course {
	subject := strings.TrimSpace(s.Find("td").Eq(0).Text())
	catalog := strings.TrimSpace(s.Find("td").Eq(1).Text())
	termStr := strings.TrimSpace(s.Find("td").Eq(2).Text())
	class := strings.TrimSpace(s.Find("td").Eq(3).Text())
	title := strings.TrimSpace(s.Find("td").Eq(4).Text())
	instructor := strings.TrimSpace(s.Find("td").Eq(5).Text())
	credits := strings.TrimSpace(s.Find("td").Eq(6).Text())
	catalogNum, _ := strconv.Atoi(catalog)
	classNum, _ := strconv.Atoi(strings.TrimSpace(class))

	// Damn you unicode NBSP!!!
	filter := strings.Replace(termStr, "\u0020", "", -1)
	termCleaned := strings.Split(filter, "\u00A0")[0]

	course := Course{
		Subject:       subject,
		CatalogNumber: catalogNum,
		ClassNumber:   classNum,
		Title:         title,
		Instructor:    instructor,
		Credits:       credits,
		Term:          NewTerm(termCleaned),
	}
	return course
}
