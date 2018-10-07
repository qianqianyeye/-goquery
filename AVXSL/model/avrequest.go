package model

import "github.com/PuerkitoBio/goquery"


type AvRequest struct {
	BigName string
	SmallName string
	Url string
	Info AvInfo
	ParserFunc func(doc *goquery.Document,url string,info AvInfo) []AvInfo
}

type SourceRequest struct {
	AvRequest []AvRequest
	Url string
	ParserFunc func(doc *goquery.Document,url string,info AvInfo) SourceRequest
}


