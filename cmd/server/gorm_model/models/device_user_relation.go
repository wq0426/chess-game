package models

var DeviceUserRelationTbName = "device_user_relation"

// DeviceUserRelation 设备用户关联表
type DeviceUserRelation struct {
    Id int `gorm:"primaryKey;autoIncrement;column:id;type:int(13);NOT NULL;comment:主键" json:"id"` // 主键
    Uid int `gorm:"column:uid;type:int(13);NOT NULL;comment:用户ID" json:"uid"` // 用户ID
    Did int `gorm:"column:did;type:int(13);NOT NULL;comment:设备ID" json:"did"` // 设备ID
}