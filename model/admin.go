package model

import "time"

type Admin struct {
	AdminID    int       `xorm:"pk autoincr" json:"id"`
	AdminName  string    `xorm:"varchar(32)" json:"admin_name"`
	CreateTime time.Time `xorm:"DateTime" json:"create_time"`
	Status     int       `xorm:"default 0" json:"status"`
	Avatar     string    `xorm:"varchar(22)" json:"avatar"`
	Pwd        string    `xorm:"varchar(33)" json:"pwd"`
	CityName   string    `xorm:"varchar(33)" json:"city_name"`
	CityId     int       `xorm:"index" json:"city_id"`
	//City       *City     `xorm:"-"`
}
