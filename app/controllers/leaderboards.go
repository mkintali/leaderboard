package controllers

import (
	"github.com/robfig/revel"
	"fmt"
)

type LeaderBoards struct {
	*revel.Controller
}

//func (c Application) CreateLeaderBoard(leaderBoard models.leaderBoard) revel.Result
//{

//}

func (c LeaderBoards) ViewLeaderBoards() revel.Result {
	return c.RenderTemplate("App/leaderboards.html")
}

func (c LeaderBoards) AddLeaderBoard() revel.Result {
	return c.RenderTemplate("App/addleaderboard.html")
}

func (c LeaderBoards) InsertLeaderBoard(newLeaderBoard string) revel.Result {
	fmt.Println("inserting", newLeaderBoard);
	return c.RenderTemplate("App/leaderboards.html")
	
}
