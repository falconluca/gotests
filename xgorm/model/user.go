package model

import "time"

type User struct {
	Id        int64  `xorm:"not null pk autoincr"`
	Name      string `xorm:"varchar(25) notnull unique 'user_name' comment('姓名')"`
	Salt      string
	Age       int
	Passwd    string    `xorm:"varchar(200)"`
	CreatedAt time.Time `xorm:"created"` // FIXME created, updated tags
	UpdatedAt time.Time `xorm:"updated"`
}

type Detail struct {
	Id     int64
	UserId int64 `xorm:"index"`
}

type UserDetail struct {
	User   `xorm:"extends"`
	UserId int64
}
