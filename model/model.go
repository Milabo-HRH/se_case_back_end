package model

import (
	"github.com/jinzhu/gorm"
)

type Register struct {
	gorm.Model
	Name       string `gorm:"type:varchar(40);not null"`
	Age        uint   `gorm:"type:int unsigned;"`
	Gender     string `gorm:"type:varchar(1);" json:"Gender"`
	Department string `gorm:"type:varchar(40);not null" json:"Department"`
	Status     string `gorm:"type:varchar(1);" json:"Status"`
	UserID     uint   `gorm:"type:int unsigned;notnull"`
}

type Case struct {
	gorm.Model
	RegisterID uint   `gorm:"type:int unsigned;notnull"`
	UserID     uint   `gorm:"type:int unsigned;notnull"`
	CC         string `gorm:"type:text;notnull"`
	HOPI       string `gorm:"type:text;notnull"`
	PMH        string `gorm:"type:text"`
	PE         string `gorm:"type:text;notnull"`
	PD         string `gorm:"type:text;notnull"`
	RC         string `gorm:"type:text;notnull"`
	EDU        string `gorm:"type:text;notnull"`
}

type Supplement struct {
	gorm.Model
	ClinicID  uint   `gorm:"type:int unsigned;notnull"`
	CheckName string `gorm:"type:varchar(40);"`
	Result    string `gorm:"type:text;notnull"`
}

type Treatment struct {
	gorm.Model
	ClinicID uint   `gorm:"type:int unsigned;notnull"`
	MedName  string `gorm:"type:varchar(40);"`
	Val      uint   `gorm:"type:int unsigned;notnull"`
	Unit     string `gorm:"type:varchar(40);"`
	Usage    string `gorm:"type:text;"`
}
