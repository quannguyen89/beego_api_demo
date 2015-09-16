package controllers

import (
	"github.com/astaxie/beego"
	"go_api/models"
	"github.com/astaxie/beego/orm"
)

type HouseController struct {
	beego.Controller
}

func (h *HouseController) Get() {
	houseId, err := h.GetInt("id")
	o := orm.NewOrm()
	if err == nil {
		house := models.House{}
		o.QueryTable("house").Filter("Id", houseId).RelatedSel().One(&house)
		o.LoadRelated(&house, "Photos")
		var agents []*models.Agent
		o.QueryTable("agent").Filter("House", houseId).RelatedSel().All(&agents)
		house.ContactAgents = agents
		h.Data["json"] = house
	} else {
		var houses []*models.House
		o.QueryTable("house").RelatedSel().All(&houses)
		h.Data["json"] = houses
	}
	h.ServeJson()
}