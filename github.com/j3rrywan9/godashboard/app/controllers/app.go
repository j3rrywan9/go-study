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
	var myPlatform []models.Mad_platform
	myPlatform = models.Get_all_platforms()
	return c.Render(myPlatform)
}


func (c App) Database() revel.Result {
	var myDatabase []models.Mad_database
	myDatabase = models.Get_all_databases()
	return c.Render(myDatabase)
}

