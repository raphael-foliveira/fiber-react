package models

type RefreshToken struct {
	ID     int
	UserID int
	Token  string
}
