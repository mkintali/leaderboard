package controllers

import (
	"github.com/golang/glog"
	"github.com/robfig/revel"
	"leaderboard/app/models"
	"leaderboard/app/routes"

	"strconv"
)

type Users struct {
	*revel.Controller
}

/**
 * Render the Login View, but redirect to dashboard if already
 * logged in
 */
func (c Users) Login() revel.Result {
	return c.RenderTemplate("Users/login.html")
}

func (c Users) Logout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}

	return c.Redirect(routes.Users.Login())
}

/**
 * Handle post request for login and authetication
 * Redirect to user view dashboard if successful
 * and set the Session["userId"] to be the userId
 */
func (c Users) Auth(email string) revel.Result {
	var users []*models.User

	_, err := Dbm.Select(&users, `
	select id, firstname, lastname, email, active
	from alpha.alpha_employees
	where active = 1
	and email = ?
	limit 1`, email)

	if err != nil {
		glog.Error("Error accessing database")
		c.Flash.Error("Error accessing database. Please contact an admin.")
		return c.Redirect(routes.Users.Login())
	}

	if len(users) == 0 {
		c.Flash.Error("User does not exist. Please try again")
		return c.Redirect(routes.Users.Login())
	}

	// TODO(billy) can we make this one var?
	c.Session["userId"] = strconv.FormatInt(users[0].Id, 10)
	c.Session["userEmail"] = users[0].Email
	c.Session["userFirstName"] = users[0].FirstName
	c.Session["userLastName"] = users[0].LastName

	return c.Redirect(routes.App.Index())
}
