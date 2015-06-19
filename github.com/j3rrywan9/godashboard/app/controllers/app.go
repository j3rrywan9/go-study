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
	var MyPlatform []models.Mad_platform
	MyPlatform = models.Get_all_platforms()
	return c.Render(MyPlatform)
}

