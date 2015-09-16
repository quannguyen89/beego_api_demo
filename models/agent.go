package models

type Agent struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	AverageRating float32 `json:"averageRating"`
	NumRates int `json:"numRates"`
	RecentSales int `json:"recentSales"`
	House *House `orm:"rel(fk)"` // RelForeignKey relation
	Avatar *Photo `orm:"null;rel(one);on_delete(set_null)" json:"avatar"`
}