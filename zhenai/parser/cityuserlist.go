package parser

import (
	"spider/types"
	"regexp"
)

var cityUserListRe =regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9a-z])+"[^>]*>[^<]+</a>`)
func ParserCityUserList(contents []byte) types.ParserResult  {
	matches := cityUserListRe.FindAllSubmatch(contents,-1)
	result := types.ParserResult{}
	for _,m:=range matches {
		name := string(m[2])
		result.Request=append(result.Request,
			types.Request{
			Url:string(m[1]),
			ParserFunc: func(bytes []byte) types.ParserResult {
				return ParseProfile(bytes, name)
			},
		})
	}
}