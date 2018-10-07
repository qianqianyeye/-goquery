package parser

import (
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"spider/AVXSL/model"
	"strings"
	"strconv"
	"spider/fetcher"
	"log"
	"fmt"
)

func ParserSource(doc *goquery.Document, url string, info model.AvInfo) model.SourceRequest {
	var sourceRequest model.SourceRequest
	allmenu := doc.Find("#allmenu")
	nav := allmenu.Find("#nav")
	//获取分类下的信息
	nav.Find("li").Each(func(i int, selection *goquery.Selection) {
		if i != 0 && i != 6 && i != 7 {
			id, _ := selection.Attr("id")
			BigName := selection.Find("a").Text()
			info.TypeBig = BigName
			//找到大分类下的小分类
			allmenu.Find("#bmenu").Find("ul").Each(func(i int, selection *goquery.Selection) {
				class, _ := selection.Attr("class")
				if class == id {
					selection.Find("li").Each(func(i int, selection *goquery.Selection) {
						SmallHref, _ := selection.Find("a").Attr("href")
						SmallHref = url + SmallHref
						SmallName := selection.Find("a").Text()
						info.TypeSmall = SmallName
						if SmallName != "" {
							sourceRequest.AvRequest = append(sourceRequest.AvRequest, model.AvRequest{
								BigName:   BigName,
								SmallName: SmallName,
								Url:       SmallHref,
								Info:      info,
								ParserFunc: func(doc *goquery.Document, url string, info model.AvInfo) []model.AvInfo{
									return ParserAvRequest(doc, url, info)
								}})
						}
					})
				}
			})
			//fmt.Printf("Review %d: %s - %s\n", i, href, title)
		}
	})
	return sourceRequest
}

func ParserAvRequest(doc *goquery.Document, url string, info model.AvInfo) []model.AvInfo {
	var infos []model.AvInfo
	list := doc.Find(".list")
	list.Each(func(i int, selection *goquery.Selection) {
		href, _ := selection.Find("a").Attr("href")
		info.Href = url + href
		info.ImgSrc, _ = selection.Find("a").Find("img").Attr("src")
		info.Title = selection.Find("p").Find("a").Eq(0).Text()
		info.Types = selection.Find("p").Find("a").Eq(1).Text()
		tempTag := selection.Find("div").Find(".tag").Text()
		tag := strings.Split(tempTag, "：")
		info.Tag = tag[1]
		pstring := selection.Find("p").Text()
		ps := regexp.MustCompile("\n\t[^\n\t]*([^\n\t]+)([^\n\t]+)([^\n\t]+)")
		matches := ps.FindAllString(pstring, -1)
		info.Name = strings.TrimSpace(matches[1])
		info.Numbers = strings.TrimSpace(matches[2])
		info.Time = strings.TrimSpace(matches[3])
		in,err:=ParserXQ(info,url)
		if err!=nil {
			infos = append(infos, info)
			model.Avmap.Store(info.Numbers,info)
			fmt.Println(info)
		}else {
			for _,info :=range in {
				infos = append(infos, info)
				model.Avmap.Store(info.Numbers,info)
				fmt.Println(info)
			}
		}
	})
	//fmt.Println(infos)
	return infos
}

func ParserXQ(info model.AvInfo,url string) ([]model.AvInfo,error){
	doc,err:=fetcher.Fetcher(info.Href)
	if err!=nil {
		log.Println("err%v",err)
		return nil,err
	}
	var resultInfo []model.AvInfo
	var hrefs []string
	doc.Find("#playlist").Find("a").Each(func(i int, selection *goquery.Selection) {
		temp,_:=selection.Attr("href")
		hrefs=append(hrefs,url+temp)
	})
	info.Summary=doc.Find(".pbox").Eq(1).Find("p").Text()
	info.Photo,_=doc.Find(".pbox").Find("a").Find("img").Attr("src")
	for  _,href := range hrefs{
		info.Bhref=href
		info,_:=GetXfPlay(info)
		resultInfo=append(resultInfo,info)
	}
	return resultInfo,nil
}

func GetXfPlay(info model.AvInfo) (model.AvInfo,error){
	doc,err:=fetcher.Fetcher(info.Bhref)
	if err!=nil {
		log.Println("err%v",err)
		return info,err
	}
	doc.Find("script").EachWithBreak(func(i int, selection *goquery.Selection) bool {
		nurl:=selection.Text()
		b:=strings.Contains(nurl,"var nurl = \"")
		if b==true {
			r:=strings.NewReplacer("var nurl = \"","","\";","")
			nurl=r.Replace(nurl)
			info.XfPlay=nurl
			return false
		}
		return true
	})
	return info,nil
}
//获取总页码
func GetTotalPage(doc *goquery.Document) int {
	var result int
	doc.Find(".pagenav").Find("a").Each(func(i int, selection *goquery.Selection) {
		temp:=selection.Text()
		page,err:=strconv.Atoi(temp)
		if err!=nil {
		}else {
			result=page
		}
	})
	return result
}
