package models

var RankTbName = "rank"

type Rank struct {
    Id int `gorm:"primaryKey;autoIncrement;column:id;type:int(11);NOT NULL;" json:"id"` 
    AccountName string `gorm:"column:account_name;type:varchar(100);NOT NULL;comment:账号名称" json:"account_name"` // 账号名称
    Region string `gorm:"column:region;type:varchar(100);NOT NULL;comment:地区" json:"region"` // 地区
    Date time.Time `gorm:"column:date;type:date;NOT NULL;comment:当天日期" json:"date"` // 当天日期
    CreatedAt time.Time `gorm:"column:created_at;type:datetime;NOT NULL;comment:创建时间" json:"created_at"` // 创建时间
}