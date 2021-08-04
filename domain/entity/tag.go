package entity

import "time"

type Tag struct {
	ID        int64     `json:"tagid" db:"tagid"`
	TagName   string    `json:"tagname" db:"tagname"`
	AccountID int64     `json:"accountid" db:"accountid"`
	Created   time.Time `json:"created" db:"created"`
	Updated   time.Time `json:"updated" db:"updated"`
}

type CreateTag struct {
	ID        int64     `json:"tagid" db:"tagid"`
	TagName   string    `json:"tagname" db:"tagname"`
	AccountID int64     `json:"accountid" db:"accountid"`
	Created   time.Time `json:"created" db:"created"`
}

type Tags []Tag

type ProductTags struct {
	NmID int64
	Tags []int64
}
type TagsByProduct []ProductTags

