package controllers

import (
	"github.com/golang/glog"
	"github.com/robfig/revel"
	"leaderboard/app/models"
	"leaderboard/app/routes"
)

type LeaderBoards struct {
	*revel.Controller
}

// TODO (mkintali): This is garbage, clean this shit up
func (c LeaderBoards) ViewBoard(boardId int64) revel.Result {
	var boards []*models.Leaderboard
	var players []*models.User

	_, err := Dbm.Select(&boards, `select * from leaderboard.leaderboards where id = ?`, boardId)

	if err != nil {
		glog.Error(err)
	}

	// TODO (billy) : Update DB model (board_id, user_id) since they are foreign keys pointing to other tables
	_, err = Dbm.Select(&players, `
		select e.Id, e.FirstName, e.LastName, e.Email, e.Active
		from leaderboard_players as p
		join alpha.alpha_employees as e
			on p.userId = e.Id
		where e.active = 1
			and p.boardId = ?`, boardId)

	c.RenderArgs["board"] = boards[0]
	c.RenderArgs["players"] = players

	return c.RenderTemplate("Leaderboards/leaderboard.html")
}

func (c LeaderBoards) ViewBoards() revel.Result {
	users := GetAllUsers()
	var boards []*models.Leaderboard
	_, err := Dbm.Select(&boards, `select * from leaderboard.leaderboards`)
	if err != nil {
		glog.Error(err)
	}
	c.RenderArgs["boards"] = boards
	c.RenderArgs["users"] = users

	return c.RenderTemplate("Leaderboards/leaderboards.html")
}

func (c LeaderBoards) AddBoard() revel.Result {
	return c.RenderTemplate("Leaderboards/addleaderboard.html")
}

func (c LeaderBoards) InsertBoard(name string) revel.Result {
	board := &models.Leaderboard{
		Name: name,
	}
	err := Dbm.Insert(board)
	if err != nil {
		glog.Errorf("Error creating a new leaderboard", err)
	}
	return c.Redirect(routes.LeaderBoards.ViewBoards())
}

func (c LeaderBoards) JoinBoard(boardId int64, userId int64) revel.Result {
	boardPlayer := &models.LeaderboardPlayer{
		BoardId: boardId,
		UserId:  userId,
	}

	err := Dbm.Insert(boardPlayer)
	if err != nil {
		glog.Error("Error joining board", err)
	}

	return c.RenderJson(boardPlayer)
}
