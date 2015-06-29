package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/revel/revel"
	"github.com/j3rrywan9/godashboard/app/models"
)

var dbC gorm.DB

func initDB() gorm.DB {
	sqlConnection := "user=postgres dbname=postgres password=postgres host=10.43.0.157 port=5432 sslmode=disable"
	dbConn, err := gorm.Open("postgres", sqlConnection)
	if err != nil {
		panic(err)
	}
	dbConn.SingularTable(true)
	return dbConn
}

func init() {
	dbC = initDB()
}

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Platform() revel.Result {
	var myPlatform []models.Mad_platform
	myPlatform = models.Get_all_platforms(dbC)
	return c.Render(myPlatform)
}

func (c App) Database() revel.Result {
	var myDatabase []models.Mad_database
	myDatabase = models.Get_all_databases(dbC)
	return c.Render(myDatabase)
}

func (c App) Build() revel.Result {
	var myBuild []models.Mad_build
	myBuild = models.Get_all_builds(dbC)
	return c.Render(myBuild)
}

type Bbrun_with_version struct {
	Id int
	Git_hash string
	Run_type string
	Version string
	Jenkins_build_url string
	Duration string
	Email string
}

func convertElapsedSecondsToString(seconds int64) string {
	time_elapsed_str := strconv.FormatInt(seconds, 10) + "s"
	d, err := time.ParseDuration(time_elapsed_str)
	if err != nil {
		panic(err)
	}
	hour := int64(d/time.Hour)
	min := int64(d/time.Minute) - hour * 60
	sec := seconds - hour * 60 * 60 - min * 60
	h := strconv.FormatInt(hour, 10)
	m := strconv.FormatInt(min, 10)
	s := strconv.FormatInt(sec, 10)
	return fmt.Sprintf("%s:%s:%s", h, m, s)
}

func (c App) BbRun() revel.Result {
	var myBbrun []models.Mad_bbrun
	myBbrun = models.Get_all_bbruns(dbC)
	var myBuild []models.Mad_build
	myBuild = models.Get_all_builds(dbC)
	myBbrunWithVersion := make([]Bbrun_with_version, len(myBbrun))
	for i := 0; i < len(myBbrun); i++ {
		myBbrunWithVersion[i].Id = myBbrun[i].Id
		myBbrunWithVersion[i].Git_hash = myBbrun[i].Git_hash
		myBbrunWithVersion[i].Run_type = myBbrun[i].Run_type
		for j := 0; j < len(myBuild); j ++ {
			if myBuild[j].Id == myBbrun[i].Build_id {
				myBbrunWithVersion[i].Version = myBuild[j].Version
				break
			}
		}
		myBbrunWithVersion[i].Jenkins_build_url = myBbrun[i].Jenkins_build_url
		myBbrunWithVersion[i].Duration = convertElapsedSecondsToString(int64(myBbrun[i].Time_elapsed))
		myBbrunWithVersion[i].Email = myBbrun[i].Email
	}
	return c.Render(myBbrunWithVersion)
}

type MyHtml string

func (r MyHtml) Apply(req *revel.Request, resp *revel.Response) {
	resp.WriteHeader(http.StatusOK, "text/html")
	resp.Out.Write([]byte(r))
}

func (c App) List_Platform() revel.Result {
	var myPlatform []models.Mad_platform
	myPlatform = models.Get_all_platforms(dbC)
	var Html string
	Html = "<table width=\"100%\"><tr><th>ID</th><th>Name</th>"
	for i := 0; i < len(myPlatform); i++ {
		Html += "<tr align=\"center\" width=\"100%\"><td>" + fmt.Sprintf("%d", myPlatform[i].Id) + "</td><td>" + myPlatform[i].Name + "</td></tr>"
	}
	Html += "</table>"
	return MyHtml(Html)
}
