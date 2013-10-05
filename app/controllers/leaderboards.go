package controllers

import "github.com/robfig/revel"

type LeaderBoards struct {
	*revel.Controller
}

func (c LeaderBoards) ViewBoards() revel.Result {
	return c.RenderTemplate("App/leaderboards.html")
}

func (c LeaderBoards) AddBoard() revel.Result {
	return c.RenderTemplate("App/addleaderboard.html")
}
