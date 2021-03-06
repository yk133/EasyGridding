package models

type Record struct {
	UserId    string  `gorm:"column:user_id"`
	Type      int     `gorm:"column:type" ` // Type 1 buy ，2 sell 3 other
	Name      string  `gorm:"column:name"`
	BuyPrice  float64 `gorm:"column:buy_price"`
	WantPrice float64 `gorm:"column:want_price"`
	Status    int    `gorm:"column:status"`

	MTime int `gorm:"column:mtime"`
	CTime int `gorm:"column:ctime"`
}

var createRecordTable = `
create Table if not exists record(
    user_id varchar(128) NOT NULL primary key,
    type integer NOT NULL,
    name varchar(255) NOT NULL,
    buy_price decimal(20) NOT NULL,
    want_price decimal(20) ,
    status int ,
    ctime bigint NOT NULL,
    mtime bigint not NULL
)`
