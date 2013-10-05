package controllers

import (
	"github.com/coopernurse/gorp"
	"github.com/robfig/revel"
	db "github.com/robfig/revel/modules/db/app"
	"leaderboard/app/models"
)

var (
	Dbm *gorp.DbMap
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func Init() {
	db.Init()
	Dbm = &gorp.DbMap{Db: db.Db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	Dbm.AddTableWithName(models.Leaderboard{}, "leaderboards")
	Dbm.AddTableWithName(Challenge{}, "challenges").SetKeys(true, "id")
}

func init() {
	revel.OnAppStart(Init)
}
