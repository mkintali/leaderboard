package controllers

import (
	"github.com/golang/glog"
	"leaderboard/app/models"
)

func GetAllUsers() []*models.User {
	var users []*models.User
	_, err := Dbm.Select(&users,
		`SELECT 
			Id,Active,FirstName,LastName
		FROM 
			alpha.alpha_employees
		WHERE
			Active = 1`,
	)
	if err != nil {
		glog.Error(err)
	}

	return users
}
