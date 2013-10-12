package controllers

import (
	"github.com/coopernurse/gorp"
	"github.com/robfig/revel"
	db "github.com/robfig/revel/modules/db/app"
	"leaderboard/app/models"
	"leaderboard/app/routes"
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

func checkUser(c *revel.Controller) revel.Result {
	if c.Request.URL.String() == routes.Users.Login() {
		return nil
	}

	if len(c.Session["userId"]) == 0 {
		c.Flash.Error("You must be logged in to use the leaderboard.")
		return c.Redirect(routes.Users.Login())
	}

	c.RenderArgs["session"] = c.Session
	return nil
}

func Init() {
	db.Init()
	Dbm = &gorp.DbMap{Db: db.Db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	Dbm.AddTableWithName(models.Leaderboard{}, "leaderboards").SetKeys(true, "id")
	Dbm.AddTableWithName(models.Challenge{}, "challenges").SetKeys(true, "id")
	Dbm.AddTableWithName(models.User{}, "alpha.alpha_employees").SetKeys(true, "Id")
	Dbm.AddTableWithName(models.LeaderboardPlayer{}, "leaderboard_players")
}

func init() {
	revel.InterceptFunc(checkUser, revel.AFTER, &revel.Controller{})
	revel.OnAppStart(Init)
}
