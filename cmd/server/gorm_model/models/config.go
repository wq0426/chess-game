package models

var ConfigTbName = "config"

type Config struct {
    Id int `gorm:"primaryKey;autoIncrement;column:id;type:int(10) unsigned;NOT NULL;comment:配置主键ID" json:"id"` // 配置主键ID
    Module string `gorm:"column:module;type:varchar(50);default:;NOT NULL;comment:所属模块" json:"module"` // 所属模块
    Key string `gorm:"column:key;type:varchar(50);default:;NOT NULL;comment:配置项键" json:"key"` // 配置项键
    Value string `gorm:"column:value;type:varchar(255);default:;NOT NULL;comment:配置项值" json:"value"` // 配置项值
}