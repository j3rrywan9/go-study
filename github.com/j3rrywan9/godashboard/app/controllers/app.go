package controllers

import (
	"github.com/revel/revel"
	"github.com/j3rrywan9/godashboard/app/models"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Platform() revel.Result {
	myPlatform := models.Get_platform_name_by_id(50)
	Id := myPlatform.Id
	Name := myPlatform.Name
	return c.Render(Id, Name)
}