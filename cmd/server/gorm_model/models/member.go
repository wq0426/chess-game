package models

var MemberTbName = "member"

type Member struct {
    Mid int `gorm:"primaryKey;column:mid;type:int(13);NOT NULL;comment:会员主键ID" json:"mid"` // 会员主键ID
    Type int8 `gorm:"column:type;type:tinyint(4);NOT NULL;comment:会员类型（1:免费 2:月会员 3:年会员 4:永久 5:自定义）" json:"type"` // 会员类型（1:免费 2:月会员 3:年会员 4:永久 5:自定义）
    Limit int `gorm:"column:limit;type:int(14);NOT NULL;comment:token上限" json:"limit"` // token上限
    Expired int64 `gorm:"column:expired;type:bigint(14);NOT NULL;comment:有效期时长（单位:s）" json:"expired"` // 有效期时长（单位:s）
}