package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

type Emoji struct {
	gorm.Model
	Title      string         `gorm:"type:varchar" json:"title"`
	CodePoints string         `gorm:"type:varchar" json:"code_points"`
	ShortCodes string         `gorm:"type:varchar" json:"short_codes"`
	Tags       pq.StringArray `gorm:"type:varchar[]" json:"tags"`
	Url        string         `gorm:"type:varchar" json:"url"`
	VersionID  uint
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

func Collection() []interface{} {
	return []interface{}{
		&Emoji{},
		&Version{},
	}
}
