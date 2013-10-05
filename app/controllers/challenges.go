package controllers

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/robfig/revel"
)

type Challenges struct {
	*revel.Controller
}

type Challenge struct {
	Id            int64  `db:"id"`
	LeaderBoardId int64  `db:"leaderBoard_id"`
	ToUserId      int64  `db:"toUser_id"`
	FromUserId    int64  `db:"fromUser_id"`
	Message       string `db:"message"`
	WinnerId      int64  `db:"winner_id"`
	LoserId       int64  `db:"loser_id"`
}

/**
 * Create a challenge between two people on a particular leaderboard
 *
 * @param {int64} boardId The leaderboad where the challenge should take place
 * @param {int64} fromUserId The challenger
 * @param {int64} toUserId The challenged
 * @param {string} message The smack talk message
 */
func (c Challenges) Create(leaderBoardId int64, fromUserId int64, toUserId int64, msg string) revel.Result {
	challenge := &Challenge{
		LeaderBoardId: leaderBoardId,
		ToUserId:      toUserId,
		FromUserId:    fromUserId,
		Message:       msg,
	}

	err := Dbm.Insert(challenge)

	if err != nil {
		glog.Error(err)
	}

	return c.RenderJson(challenge)
}

/**
 * Get a challenge based on the challenge ID
 * @param {int64} id The ID of the challenge to retrieve
 */
func (c Challenges) Get(challengeId int64) revel.Result {
	challenge, err := Dbm.Get(Challenge{}, challengeId)
	if err != nil {
		glog.Error(err)
	}

	return c.RenderJson(challenge)
}

/**
 * Get all the challenges for a user ID
 * @param {int64} userId The user id to fetch the challenges for
 */
func (c Challenges) GetAll(userId int64) revel.Result {
	var challenges []*Challenge
	_, err := Dbm.Select(&challenges, "SELECT * FROM challenges WHERE toUser_id = ? OR fromUser_id = ?", userId, userId)

	if err != nil {
		glog.Error(err)
	}

	fmt.Println(challenges)

	return c.RenderJson(challenges)
}

/**
 * Set the winner of the challenge
 */
func (c Challenges) setWinner(challengeId int64, winnerId int64, loserId int64) revel.Result {

	return c.RenderText("TBD")
}
