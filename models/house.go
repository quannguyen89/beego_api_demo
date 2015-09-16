package models


type House struct {
	Id int `json:"id"`
	Address string `json:"address"`
	HouseType string `json:"type"`
	Price int `json:"price"`
	Properties string `json:"properties"`
	PublishFrom int `json:"publishFrom"`
	Description string `orm:"size(1000)" json:"description"`
	Location *Location `orm:"null;rel(one);on_delete(set_null)" json:"location"`
	Thumbnail *Photo `orm:"null;rel(one);on_delete(set_null)" json:"thumbnail"`
	Photos []*Photo `orm:"reverse(many)" json:"photos"` // reverse relationship of fk
	ContactAgents []*Agent `orm:"reverse(many)" json:"contactAgents"` // reverse relationship of fk
}