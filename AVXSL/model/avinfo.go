package model
type AvInfo struct {
	ID int64 `gorm:"column:id" json:"id"`
	Bhref string `gorm:"column:bhref" json:"bhref"`
	Href string `gorm:"column:href" json:"href"`
	ImgSrc string `gorm:"column:img_src" json:"img_src"`
	Title string `gorm:"column:title" json:"title"`
	Types string `gorm:"column:types" json:"types"`
	Tag string `gorm:"column:tag" json:"tag"`
	Name string `gorm:"column:name" json:"name"`
	Numbers string `gorm:"column:numbers" json:"numbers"`
	Time string `gorm:"column:time" json:"time"`
	XfPlay string `gorm:"column:xf_play" json:"xf_play"`
	Summary string `gorm:"column:summary" json:"summary"`
	Photo string `gorm:"column:photo" json:"photo"`
	TypeBig string `gorm:"column:type_big" json:"type_big"`
	TypeSmall string `gorm:"column:type_small" json:"type_small"`
}
