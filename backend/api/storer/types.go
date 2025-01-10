package storer

import "time"

type Account struct {
	ID        string    `db:"id"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
}

type Session struct {
	ID          string    `db:"id" json:"id"`
	AdminID     string    `db:"admin_id" json:"admin_id"`
	Name        string    `db:"name" json:"name"`
	Picture     string    `db:"picture" json:"picture"`
	Seed        string    `db:"seed" json:"seed"`
	PlayerCount int64     `db:"player_count" json:"player_count"`
	IsArchived  bool      `db:"is_archived" json:"is_archived"`
	CreatedAt   time.Time `db:"created_at"`
}

type PlayerCard struct {
	ID          int64  `db:"id"`
	AccountID   string `db:"account_id"`
	SessionID   string `db:"session_id"`
	Nickname    string `db:"nickname"`
	Preferences string `db:"preferences"`
}
