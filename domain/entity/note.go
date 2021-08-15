package entity

import "time"

type Note struct {
	AccountID int64     `json:"accountid" db:"accountid"`
	ProductID int64     `json:"productid" db:"nmid"`
	Text      string    `json:"text" db:"note"`
	Created   time.Time `json:"created" db:"created"`
	Updated   time.Time `json:"updated" db:"updated"`
}

type Notes []Note

type CreateNote struct {
	AccountID int64  `json:"accountid" db:"accountid"`
	ProductID int64  `json:"productid" db:"nmid"`
	Text      string `json:"text" db:"note"`
}
