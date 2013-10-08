package models

type Leaderboard struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
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

type User struct {
	Id        int64
	Active    bool
	Email     string
	FirstName string
	LastName  string
}

type LeaderboardPlayer struct {
	BoardId int64
	UserId  int64
}
