package model

import (
	"gorm.io/gorm"
)

type Penyakit struct {
	gorm.Model
	Nama string `gorm:"type:varchar(100)"`
	DnaSequence string `gorm:"type:text"`
}
