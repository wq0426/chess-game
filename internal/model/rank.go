package model

type Rank struct {
	Id          int    `gorm:"primaryKey;autoIncrement;column:id;type:int(11);NOT NULL;" json:"id"`
	AccountName string `gorm:"column:account_name;type:varchar(100);NOT NULL;comment:账号名称" json:"account_name"` // 账号名称
	Region      string `gorm:"column:region;type:varchar(100);NOT NULL;comment:地区" json:"region"`               // 地区
	Date        string `gorm:"column:date;type:date;NOT NULL;comment:当天日期" json:"date"`                         // 当天日期
	CreatedAt   string `gorm:"column:created_at;type:datetime;NOT NULL;comment:创建时间" json:"created_at"`         // 创建时间
}

func (m *Rank) TableName() string {
	return "rank"
}

type UserRankParam struct {
	Region      string `json:"region" binding:"required"`
	AccountName string `json:"account_name" binding:"required"`
}

func (v *UserRankParam) Validate() bool {
	if len(v.Region) == 0 || len(v.AccountName) == 0 {
		return false
	}
	return true
}

type RankList struct {
	Region string `json:"region"` // 地区
	Total  int    `json:"total"`  // 总人数
	Rank   int    `json:"rank"`   // 总排名
}
