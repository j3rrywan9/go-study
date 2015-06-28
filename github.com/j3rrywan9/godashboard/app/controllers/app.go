package controllers

import (
	"fmt"
	"net/http"
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

func (c App) Build() revel.Result {
	var myBuild []models.Mad_build
	myBuild = models.Get_all_builds()
	return c.Render(myBuild)
}

func (c App) BbRun() revel.Result {
	var myBbrun []models.Mad_bbrun
	myBbrun = models.Get_all_bbruns()
	return c.Render(myBbrun)
}

type MyHtml string

func (r MyHtml) Apply(req *revel.Request, resp *revel.Response) {
	resp.WriteHeader(http.StatusOK, "text/html")
	resp.Out.Write([]byte(r))
}

func (c App) List_Platform() revel.Result {
	var myPlatform []models.Mad_platform
	myPlatform = models.Get_all_platforms()
	var Html string
	Html = "<table width=\"100%\"><tr><th>ID</th><th>Name</th>"
	for i := 0; i < len(myPlatform); i++ {
		Html += "<tr align=\"center\" width=\"100%\"><td>" + fmt.Sprintf("%d", myPlatform[i].Id) + "</td><td>" + myPlatform[i].Name + "</td></tr>"
	}
	Html += "</table>"
	return MyHtml(Html)
}
