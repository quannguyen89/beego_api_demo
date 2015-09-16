package models


type Location struct {
	Id int `json:"id"`
	Longtitude float32 `json:"longtitude"`
	Latitude float32 `json:"latitude"`
	House *House `orm:"reverse(one)" json:"-"`
}