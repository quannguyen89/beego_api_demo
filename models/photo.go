package models
type Photo struct {
	Id int `json:"id"`
	Url string `json:"url"`
	SmallHouse *House `orm:"reverse(one)" json:"-"`
	Agent *House `orm:"reverse(one)" json:"-"`
	House *House `orm:"null;rel(fk)"` // RelForeignKey relation
}