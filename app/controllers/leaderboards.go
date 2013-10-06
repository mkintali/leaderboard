package controllers

import (
	"github.com/golang/glog"
	"github.com/robfig/revel"
	"leaderboard/app/models"
)

type LeaderBoards struct {
	*revel.Controller
}

func (c LeaderBoards) ViewBoards() revel.Result {
	var boards []*models.Leaderboard
	_, err := Dbm.Select(&boards, `select * from leaderboard.leaderboards`)
	if err != nil {
		glog.Error(err)
	}
	c.RenderArgs["boards"] = boards
	return c.RenderTemplate("Leaderboards/leaderboards.html")
}

func (c LeaderBoards) AddBoard() revel.Result {
	return c.RenderTemplate("Leaderboards/addleaderboard.html")
}
