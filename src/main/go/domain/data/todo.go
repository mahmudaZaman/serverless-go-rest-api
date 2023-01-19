package data

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
}

//type OrMapping struct {
//	OrId           int    `json:"or_id"`
//	SubSpecialtyId int    `json:"sub_specialty_id"`
//	Weekday        string `json:"weekday"`
//	OpeningTime    string `json:"opening_time"`
//	ClosingTime    string `json:"closing_time"`
//	AnesthType     string `json:"anesth_type"`
//	WeekId         []int  `json:"week_id"`
//}

type OrMapping struct {
	ID           uint   `gorm:"primarykey"`
	OrId         int    `json:"or_id"`
	SubSpecialty string `json:"sub_specialty"`
	Weekday      string `json:"weekday"`
	OpeningTime  string `json:"opening_time"`
	ClosingTime  string `json:"closing_time"`
	AnesthType   string `json:"anesth_type"`
	//WeekId       []int  `json:"week_id"`
	WeekId pq.Int32Array `json:"week_id" gorm:"type:int[]"`
}
