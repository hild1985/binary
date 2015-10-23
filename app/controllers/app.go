package controllers

import "github.com/revel/revel"
import "labix.org/v2/mgo"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {	
	//db = connect("192.168.59.103:27017/myDatabase")
	return c.Render()
}
