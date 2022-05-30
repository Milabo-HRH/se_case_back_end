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
	DoctorID   uint   `gorm:"type:int unsigned;notnull"`
	DoctorName string `gorm:"type:varchar(40);not null"`
}

type Case struct {
	gorm.Model
	RegisterID uint   `gorm:"type:int unsigned;notnull" json:"registerID"`
	CC         string `gorm:"type:text;notnull" json:"cc"`
	HOPI       string `gorm:"type:text;notnull" json:"hopi"`
	PMH        string `gorm:"type:text" json:"pmh"`
	PE         string `gorm:"type:text;notnull" json:"pe"`
	PD         string `gorm:"type:text;notnull" json:"pd"`
	RC         string `gorm:"type:text;notnull" json:"rc"`
	EDU        string `gorm:"type:text;notnull" json:"edu"`
}

type Supplement struct {
	gorm.Model
	ClinicID  uint   `gorm:"type:int unsigned;notnull" json:"clinicID"`
	CheckName string `gorm:"type:varchar(40);" json:"checkName"`
	Result    string `gorm:"type:text;notnull" json:"result"`
}

type Treatment struct {
	gorm.Model
	ClinicID uint   `gorm:"type:int unsigned;notnull" json:"clinicID"`
	MedName  string `gorm:"type:varchar(40);" json:"medName"`
	Val      uint   `gorm:"type:int unsigned;notnull" json:"val"`
	Unit     string `gorm:"type:varchar(40);" json:"unit"`
	Usage    string `gorm:"type:text;" json:"usage"`
}

type Sps struct {
	Sps []Supplement `json:"sps"`
}

type Trs struct {
	Trs []Treatment `json:"trs"`
}
