package controllers

import (
	"github.com/golang/glog"
	"github.com/robfig/revel"
	"leaderboard/app/models"
)

type LeaderBoards struct {
	*revel.Controller
}

// TODO (mkintali): This is garbage, clean this shit up
func (c LeaderBoards) ViewBoard(boardId int64) revel.Result {
	var boards []*models.Leaderboard
	_, err := Dbm.Select(&boards, `select * from leaderboard.leaderboards where id = ?`, boardId)
	if err != nil {
		glog.Error(err)
	}

	c.RenderArgs["board"] = boards[0]
	return c.RenderTemplate("Leaderboards/leaderboard.html")
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
