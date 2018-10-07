package fetcher

import (
	"net/http"
	"fmt"
	"bufio"
	"golang.org/x/text/encoding"
	"github.com/labstack/gommon/log"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/net/html/charset"
	"github.com/PuerkitoBio/goquery"
)

//获取网页内容
func Fetcher (urls string)(*goquery.Document,error){
	//ip :=iputil.ReturnIp()
	//proxy := func(_ *http.Request) (*url.URL, error) {
	//	return url.Parse("http://178.128.104.199:8080")//根据定义Proxy func(*Request) (*url.URL, error)这里要返回url.URL
	//}
	//transport := &http.Transport{Proxy: proxy}
	//client := &http.Client{Transport: transport}
	//resp, err := client.Get(urls) //请求并获取到对象,使用代理
	//if err != nil {
	//	log.Fatal(err)
	//	return nil,err
	//}
	resp,err:=http.Get(urls)
	if err!=nil {
		return nil,err
	}
	defer resp.Body.Close()
	if resp.StatusCode !=http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	bodyReader := bufio.NewReader(resp.Body)
	doc, err := goquery.NewDocumentFromReader(bodyReader)
	if err != nil {
		log.Fatal(err)
		return nil,err
	}

	return doc,nil
}

//转码转成UTF-8
func determineEncoding(r *bufio.Reader) encoding.Encoding{
	bytes, err := r.Peek(1024)
	if err!=nil {
		log.Print("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
