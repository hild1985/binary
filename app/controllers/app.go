package controllers

import (
        "fmt"
		"log"
        "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}