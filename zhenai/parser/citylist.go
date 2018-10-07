package parser

import (
	"spider/types"
	"regexp"
)

var CityListRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z])+"[^>]*>([^<]+)</a>`)
func ParserCityList(contents []byte)  types.ParserResult{
	matches := CityListRe.FindAllSubmatch(contents,-1)
	result := types.ParserResult{}
	for _,m:=range matches {
		result.Request = append(result.Request,types.Request{Url:string(m[1]),ParserFunc:ParserCityUserList})
	}
	return result
}
