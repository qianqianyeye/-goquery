package main

import (
	"spider/engine"
	"spider/AVXSL/parser"
	"spider/AVXSL/model"
	"time"
	"spider/db"
)

func main()  {
	db.InitDB() //初始化数据库
	e:=engine.Engine{WorkCount:2}
	go engine.BatchResult()
	e.Run(model.SourceRequest{
		Url:"http://www.avxsl5.com/",
		ParserFunc:parser.ParserSource,
	})
	for  {
		time.Sleep(2*time.Hour)
	}
}