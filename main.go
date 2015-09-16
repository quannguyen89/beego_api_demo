package main

import (
	"fmt"
	_"go_api/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"go_api/models"
	_"github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	orm.Debug = true
	orm.RegisterModel(new(models.Location))
	orm.RegisterModel(new(models.Photo))
	orm.RegisterModel(new(models.Agent))
	orm.RegisterModel(new(models.House))
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "root@/go_api_db?charset=utf8", 30)
}

func main() {
	// Database alias.
	name := "default"

	// Drop table and re-create.
	force := true

	// Print log.
	verbose := true

	// Error.
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
	o := orm.NewOrm()
	o.Using("default")
	seed(o)

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
	}))

	beego.Run()
}


func seed(o orm.Ormer) {
	arr := [10]map[string]interface{}{}
	arr[0] = map[string]interface{}{
		"address": "589 Coleridge Ave, Hanoi, Vietnam",
		"houseType": "House for sale",
		"price": 8950000,
		"properties": "3 bds, 3.5 ba, 4464 sqft, 0.28 ac lot, Built 1997",
		"publishFrom": 104,
		"thumbnail": "http://photos2.zillowstatic.com/p_g/ISpxm1nzj9y1xo1000000000.jpg",
		"location": []float32{21.009326, 105.857682},
		"description": "Custom designed home on a large lot in the Old Palo Alto neighborhood. Vaulted ceilings and large windows fill the home with natural light. Clean lines and craftsman details create a bond between the shingle-style neighborhood and contemporary style. Three level home, with basement. Possibility to create a fourth bedroom on either the ground floor or at the basement level.",
	}
	arr[1] = map[string]interface{}{
		"address": "1565 Webster St,  Hanoi, Vietnam",
		"houseType": "House for sale",
		"price": 6950000,
		"properties": "3 bds, 3.5 ba, 4464 sqft, 0.28 ac lot, Built 1997",
		"publishFrom": 200,
		"thumbnail": "http://photos3.zillowstatic.com/p_g/ISxzt2p00tw62m1000000000.jpg",
		"location": []float32{21.009356, 105.855442},
		"description": "Custom designed home on a large lot in the Old Palo Alto neighborhood. Vaulted ceilings and large windows fill the home with natural light. Clean lines and craftsman details create a bond between the shingle-style neighborhood and contemporary style. Three level home, with basement. Possibility to create a fourth bedroom on either the ground floor or at the basement level.",
	}
	arr[2] = map[string]interface{}{
		"address": "1061 Alma Street, Hanoi, Vietnam",
		"houseType": "Make me move",
		"price": 550000,
		"properties": "3 bds, 3.5 ba, 4464 sqft, 0.28 ac lot, Built 1997",
		"publishFrom": 222,
		"thumbnail": "http://photos3.zillowstatic.com/p_g/IS5mozylseymxd1000000000.jpg",
		"location": []float32{21.011072, 105.842455},
		"description": "Custom designed home on a large lot in the Old Palo Alto neighborhood. Vaulted ceilings and large windows fill the home with natural light. Clean lines and craftsman details create a bond between the shingle-style neighborhood and contemporary style. Three level home, with basement. Possibility to create a fourth bedroom on either the ground floor or at the basement level.",
	}
	arr[3] = map[string]interface{}{
		"address": "589 Coleridge Ave, Hanoi, Vietnam",
		"houseType": "House for sale",
		"price": 1250000,
		"properties": "3 bds, 3.5 ba, 4464 sqft, 0.28 ac lot, Built 1997",
		"publishFrom": 204,
		"thumbnail": "http://photos2.zillowstatic.com/p_g/ISpxm1nzj9y1xo1000000000.jpg",
		"location": []float32{21.027816, 105.852268},
		"description": "Custom designed home on a large lot in the Old Palo Alto neighborhood. Vaulted ceilings and large windows fill the home with natural light. Clean lines and craftsman details create a bond between the shingle-style neighborhood and contemporary style. Three level home, with basement. Possibility to create a fourth bedroom on either the ground floor or at the basement level.",
	}
	arr[4] = map[string]interface{}{
		"address": "589 Coleridge Ave, Hanoi, Vietnam",
		"houseType": "House for sale",
		"price": 1250000,
		"properties": "3 bds, 3.5 ba, 4464 sqft, 0.28 ac lot, Built 1997",
		"publishFrom": 204,
		"thumbnail": "http://photos3.zillowstatic.com/p_g/ISxzt2p00tw62m1000000000.jpg",
		"location": []float32{21.054544, 105.820344},
		"description": "Custom designed home on a large lot in the Old Palo Alto neighborhood. Vaulted ceilings and large windows fill the home with natural light. Clean lines and craftsman details create a bond between the shingle-style neighborhood and contemporary style. Three level home, with basement. Possibility to create a fourth bedroom on either the ground floor or at the basement level.",
	}
	arr[5] = map[string]interface{}{
		"address": "589 Coleridge Ave, Hanoi, Vietnam",
		"houseType": "House for sale",
		"price": 1250000,
		"properties": "3 bds, 3.5 ba, 4464 sqft, 0.28 ac lot, Built 1997",
		"publishFrom": 204,
		"thumbnail": "http://photos3.zillowstatic.com/p_g/IS5mozylseymxd1000000000.jpg",
		"location": []float32{21.042675, 105.791481},
		"description": "Custom designed home on a large lot in the Old Palo Alto neighborhood. Vaulted ceilings and large windows fill the home with natural light. Clean lines and craftsman details create a bond between the shingle-style neighborhood and contemporary style. Three level home, with basement. Possibility to create a fourth bedroom on either the ground floor or at the basement level.",
	}
	arr[6] = map[string]interface{}{
		"address": "589 Coleridge Ave, Hanoi, Vietnam",
		"houseType": "House for sale",
		"price": 1250000,
		"properties": "3 bds, 3.5 ba, 4464 sqft, 0.28 ac lot, Built 1997",
		"publishFrom": 204,
		"thumbnail": "http://photos2.zillowstatic.com/p_g/ISpxm1nzj9y1xo1000000000.jpg",
		"location": []float32{21.028501, 105.782255},
		"description": "Custom designed home on a large lot in the Old Palo Alto neighborhood. Vaulted ceilings and large windows fill the home with natural light. Clean lines and craftsman details create a bond between the shingle-style neighborhood and contemporary style. Three level home, with basement. Possibility to create a fourth bedroom on either the ground floor or at the basement level.",
	}
	arr[7] = map[string]interface{}{
		"address": "589 Coleridge Ave, Hanoi, Vietnam",
		"houseType": "House for sale",
		"price": 1250000,
		"properties": "3 bds, 3.5 ba, 4464 sqft, 0.28 ac lot, Built 1997",
		"publishFrom": 204,
		"thumbnail": "http://photos3.zillowstatic.com/p_g/IS5mozylseymxd1000000000.jpg",
		"location": []float32{21.028309, 105.790856},
		"description": "Custom designed home on a large lot in the Old Palo Alto neighborhood. Vaulted ceilings and large windows fill the home with natural light. Clean lines and craftsman details create a bond between the shingle-style neighborhood and contemporary style. Three level home, with basement. Possibility to create a fourth bedroom on either the ground floor or at the basement level.",
	}
	arr[8] = map[string]interface{}{
		"address": "589 Coleridge Ave, Hanoi, Vietnam",
		"houseType": "House for sale",
		"price": 1250000,
		"properties": "3 bds, 3.5 ba, 4464 sqft, 0.28 ac lot, Built 1997",
		"publishFrom": 204,
		"thumbnail": "http://photos2.zillowstatic.com/p_g/ISpxm1nzj9y1xo1000000000.jpg",
		"location": []float32{21.027705, 105.811117},
		"description": "Custom designed home on a large lot in the Old Palo Alto neighborhood. Vaulted ceilings and large windows fill the home with natural light. Clean lines and craftsman details create a bond between the shingle-style neighborhood and contemporary style. Three level home, with basement. Possibility to create a fourth bedroom on either the ground floor or at the basement level.",
	}
	arr[9] = map[string]interface{}{
		"address": "589 Coleridge Ave, Hanoi, Vietnam",
		"houseType": "House for sale",
		"price": 1250000,
		"properties": "3 bds, 3.5 ba, 4464 sqft, 0.28 ac lot, Built 1997",
		"publishFrom": 204,
		"thumbnail": "http://photos3.zillowstatic.com/p_g/ISxzt2p00tw62m1000000000.jpg",
		"location": []float32{21.019992, 105.814706},
		"description": "Custom designed home on a large lot in the Old Palo Alto neighborhood. Vaulted ceilings and large windows fill the home with natural light. Clean lines and craftsman details create a bond between the shingle-style neighborhood and contemporary style. Three level home, with basement. Possibility to create a fourth bedroom on either the ground floor or at the basement level.",
	}
	var house *models.House
	for i := 0; i < len(arr); i++ {
		temp := createHouse(arr[i], o)
		if i == 0 {
			house = temp
		}
	}
	thumbnail := models.Photo{}
	thumbnail.Url = "wtf"
	o.Insert(&thumbnail)

	photos := []string{
		"http://photos1.zillowstatic.com/p_h/IS9tsmiolp6hev1000000000.jpg",
		"http://photos2.zillowstatic.com/p_h/IS1nqa0ihhyjev1000000000.jpg",
		"http://photos3.zillowstatic.com/p_h/IStgoyhbd9qmev1000000000.jpg",
		"http://photos1.zillowstatic.com/p_h/ISlammz491ipev1000000000.jpg",
		"http://photos1.zillowstatic.com/p_h/IS9hr1kevjvkkl0000000000.jpg",
		"http://photos1.zillowstatic.com/p_h/IS1rxujhu5v30o1000000000.jpg",
		"http://photos1.zillowstatic.com/p_h/ISp1utk34801pv1000000000.jpg",
		"http://photos3.zillowstatic.com/p_h/IShvvoa7kf49w70000000000.jpg",
		"http://photos2.zillowstatic.com/p_h/ISphqpbkrfi2oe0000000000.jpg",
		"http://photos2.zillowstatic.com/p_h/IShzafgq64f3ka0000000000.jpg",
		"http://photos3.zillowstatic.com/p_h/IS9t83yj2w66ka0000000000.jpg",
		"http://photos3.zillowstatic.com/p_h/ISx7sev6epwv2j0000000000.jpg",
		"http://photos2.zillowstatic.com/p_h/ISp5xaj1u5nbev1000000000.jpg",
	}
	for i := 0; i < len(photos); i++ {
		photo := models.Photo{Url: photos[i], House: house}
		o.Insert(&photo)
	}

	contactAgents := [3]map[string]interface{}{}
	contactAgents[0] = map[string]interface{}{
		"name": "Michael Dreyfus",
		"phone": "(650) 644-3474",
		"averageRating": 4.7,
		"numRates": 32,
		"avatar": "http://photos1.zillowstatic.com/h_n/ISptj0herzprn40000000000.jpg",
		"recentSales": 44,
	}

	contactAgents[1] = map[string]interface{}{
		"name": "Katy Thielke Straser",
		"phone": "(650) 941-8003",
		"averageRating": 4.9,
		"numRates": 24,
		"avatar": "http://photos2.zillowstatic.com/h_n/IS-1h1vglrnwo5el.jpg",
		"recentSales": 33,
	}

	contactAgents[2] = map[string]interface{}{
		"name": "Karishma & Deepak Chandani",
		"phone": "(650) 941-8003",
		"averageRating": 4.6,
		"numRates": 31,
		"avatar": "http://photos1.zillowstatic.com/h_n/IS1fw1nb6twj6i1000000000.jpg",
		"recentSales": 49,
	}

	for i := 0; i < len(contactAgents); i++ {
		createAgent(contactAgents[i], o, house)
	}
}
func createHouse(m map[string]interface{}, o orm.Ormer) *models.House {
	location := models.Location{}
	locationData := m["location"].([]float32)
	location.Latitude = locationData[0]
	location.Longtitude = locationData[1]
	o.Insert(&location)

	thumbnail := models.Photo{}
	thumbnail.Url = m["thumbnail"].(string)
	o.Insert(&thumbnail)


	house := models.House{
		Address: m["address"].(string),
		HouseType: m["houseType"].(string),
		Price: m["price"].(int),
		Properties: m["properties"].(string),
		PublishFrom: m["publishFrom"].(int),
		Description: m["description"].(string),
		Location: &location,
		Thumbnail: &thumbnail,
	}
	o.Insert(&house)
	return &house
}

func createAgent(m map[string]interface{}, o orm.Ormer, h *models.House) {
	photo := models.Photo{}
	photo.Url = m["avatar"].(string)
	o.Insert(&photo)

	agent := models.Agent{
		Name: m["name"].(string),
		Phone: m["phone"].(string),
		AverageRating: float32(m["averageRating"].(float64)),
		NumRates: m["numRates"].(int),
		RecentSales: m["recentSales"].(int),
		House: h,
		Avatar: &photo,
	}
	o.Insert(&agent)
}
