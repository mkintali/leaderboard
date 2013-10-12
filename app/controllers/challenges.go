package controllers

import (
	"github.com/golang/glog"
	"github.com/robfig/revel"
	"leaderboard/app/models"
)

type Challenges struct {
	*revel.Controller
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
	challenge := &models.Challenge{
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
	challenge, err := Dbm.Get(models.Challenge{}, challengeId)
	if err != nil {
		glog.Error(err)
	}

	return c.RenderJson(challenge)
}

/**
 * Get all the challenges for a user ID
 * @param {int64} userId The user id to fetch the challenges for
 */
func (c Challenges) GetUserChallenges(userId int64) revel.Result {
	var challenges []*models.Challenge
	_, err := Dbm.Select(&challenges, "SELECT * FROM challenges WHERE toUser_id = ? OR fromUser_id = ?", userId, userId)

	if err != nil {
		glog.Error(err)
	}

	c.RenderArgs["userChallenges"] = challenges
	return c.RenderTemplate("Users/_userChallenges.html")
}

/**
 * Set the winner of the challenge
 */
func (c Challenges) setWinner(challengeId int64, winnerId int64, loserId int64) revel.Result {

	return c.RenderText("TBD")
}
