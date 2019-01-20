package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

type Emoji struct {
	gorm.Model
	Title       string         `gorm:"type:varchar" json:"title"`
	CodePoints  string         `gorm:"type:varchar" json:"code_points"`
	Description string         `gorm:"type:varchar" json:"description"`
	ShortCodes  string         `gorm:"type:varchar" json:"short_codes"`
	Tags        pq.StringArray `json:"tags"`
	VersionID   uint
}

func (Emoji) TableName() string {
	return "emoji"
}

type Version struct {
	gorm.Model
	VersionName string `gorm:"type:varchar" json:"version_name"`
}

func (Version) TableName() string {
	return "emoji_unicode_version"
}
