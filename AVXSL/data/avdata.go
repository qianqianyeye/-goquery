package data

import (
	"spider/AVXSL/model"
	"strings"
	"spider/db"
	"log"
)

func BatchInsertAvInfo(avinfos []model.AvInfo) {
	sqlStr := "INSERT INTO avinfo (bhref,href,img_src,title,types,tag," +
		"name,numbers,time,xf_play,summary,photo,type_big,type_small) VALUES "
	const rowSQL= "(?, ?, ?, ?, ?, ?, ?, ?, ?,?,?,?,?,?)"
	var inserts []string
	vals := []interface{}{}
	for _, elem := range avinfos {
		inserts = append(inserts, rowSQL)
		vals = append(vals, elem.Bhref, elem.Href, elem.ImgSrc, elem.Title, elem.Types,
			elem.Tag, elem.Name, elem.Numbers, elem.Time, elem.XfPlay, elem.Summary,
			elem.Photo, elem.TypeBig, elem.TypeSmall)
	}
	sqlStr = sqlStr + strings.Join(inserts, ",")
	err := db.SqlDB.Exec(sqlStr, vals...).Error
	if  err != nil {
		log.Println(err)
	}
}
