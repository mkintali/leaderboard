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
 * Check to see if a user session exists
 * @returns {bool} true if the user is logged in
 */
func (c Users) IsLoggedIn() bool {
	if len(c.Session["userId"]) == 0 {
		c.Flash.Error("You must be logged in to use the leaderboard.")
		return false
	}

	return true
}

/**
 * Render the Login View, but redirect to dashboard if already
 * logged in
 */
func (c Users) Login() revel.Result {
	if c.IsLoggedIn() {
		return c.Redirect(routes.Users.ViewDashboard())
	}
	return c.RenderTemplate("Users/login.html")
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

	c.Session["userId"] = strconv.FormatInt(users[0].Id, 10)
	return c.Redirect(routes.Users.ViewDashboard())
}

/**
 * Renders the Dashboard view and redirects user to
 * login page if they are not already logged in
 */
func (c Users) ViewDashboard() revel.Result {
	if c.IsLoggedIn() {
		return c.RenderTemplate("Users/dashboard.html")
	}
	return c.Redirect(routes.Users.Login())
}
