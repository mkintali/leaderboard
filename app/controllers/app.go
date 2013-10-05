package controllers

import (
	"github.com/coopernurse/gorp"
	"github.com/robfig/revel"
	db "github.com/robfig/revel/modules/db/app"
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

type Leaderboard struct {
	id   int64
	name string
}

func Init() {
	db.Init()
	Dbm = &gorp.DbMap{Db: db.Db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	Dbm.AddTableWithName(Leaderboard{}, "leaderboard.leaderboards")
}
