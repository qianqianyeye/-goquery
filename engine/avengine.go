package engine

import (
	"log"
	"spider/fetcher"
	"spider/AVXSL/model"
	"spider/AVXSL/parser"
	"strconv"
	"fmt"
	"spider/util"
	"spider/AVXSL/data"
	"time"
)

type Engine struct {
	start chan interface{}
	done chan interface{}
	count chan interface{}
	WorkCount int
}


func (e *Engine)Run(request...model.SourceRequest)  {
	//控制线程数量
	e.start = make(chan interface{},e.WorkCount)
	e.done= make(chan interface{},e.WorkCount+1)
	e.count=make(chan interface{},e.WorkCount+1)
	go e.listenQueue()
	//go BatchResult()
	for _,r:=range request {
		doc,err:=fetcher.Fetcher(r.Url)
		if err!=nil {
			log.Println("err%v",err)
			continue
		}
		var info model.AvInfo
		sourceRequest:=r.ParserFunc(doc,r.Url,info)
		for _,req := range sourceRequest.AvRequest {
			 e.Work(req,r.Url)
		}
	}
	for  {
		if len(e.count) >0{
			continue
		}else {
			break
		}
	}
}
func (e *Engine)listenDone()  {
	for  {
		select {
		case done:=<-e.done:
		_:<-e.count
			fmt.Println(done)
		}
	}
}
func (e *Engine)listenQueue() {
	go e.listenDone()
	for  {
			select {
			case start:=<-e.start:
				parmMap :=start.(map[string]interface{})
				req :=parmMap["req"].(model.AvRequest)
				i:=parmMap["i"].(int)
				url :=parmMap["url"].(string)
				e.count<-1
				go e.DoWork(req,i,url)
			}

	}
}

func (e *Engine)Start(parmMap map[string]interface{}){
	e.start<-parmMap
}
func (e *Engine) Done() {
	e.done<-"done..."
}

func (e *Engine)Work(req model.AvRequest,url string)  {
	doc,err:=fetcher.Fetcher(req.Url)
	if err!=nil {
		log.Println("err%v",err)
		return
	}
	total :=parser.GetTotalPage(doc) //获取该分类的总页数
	//fmt.Println(total)
	for i:=1;i<=total;i++  {
		parmMap := make(map[string]interface{})
		parmMap["req"]=req
		parmMap["i"]=i
		parmMap["url"]=url
		e.Start(parmMap)
	}
}

func (e *Engine)DoWork(req model.AvRequest,i int,url string)  {
		doc,err:=fetcher.Fetcher(req.Url+"page"+strconv.Itoa(i)+"/")
		if err!=nil {
			log.Println("err%v",err)
			e.Done()
			return
		}
		req.ParserFunc(doc,url,req.Info)
		//fmt.Println("before done")
		e.Done()
		//fmt.Println("i done")
}

func BatchResult()  {
	model.Avmap=new(syncmap.Map)
	model.AvinfoChan=make(chan interface{},2)
	go listenAvinfoChan()
	for  {

		if model.Avmap.Length()!=nil {
			if *model.Avmap.Length() >=10{
				var avinfos []model.AvInfo
				model.Avmap.Range(func(key, value interface{}) bool {
					avinfos=append(avinfos,value.(model.AvInfo))
					model.Avmap.Delete(key)
					return true
				})
				data.BatchInsertAvInfo(avinfos)
				model.AvinfoChan<-"in"
			}
		}
		time.Sleep(3*time.Second)
	}
}

func listenAvinfoChan()  {
	for  {
		select {
		case <-model.AvinfoChan:
		case <-time.After(8*time.Minute):
			if model.Avmap.Length()!=nil {
				if *model.Avmap.Length()>0{
					var avinfos []model.AvInfo
					model.Avmap.Range(func(key, value interface{}) bool {
						avinfos=append(avinfos,value.(model.AvInfo))
						model.Avmap.Delete(key)
						return true
					})
					data.BatchInsertAvInfo(avinfos)
			}

			}
		}
	}
}