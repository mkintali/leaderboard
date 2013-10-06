package models

type Leaderboard struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
}
