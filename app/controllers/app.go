package controllers

import "github.com/revel/revel"
import "gopkg.in/mgo.v2"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}