package model

import (
	"gorm.io/gorm"
)

type Pemeriksaan struct {
	gorm.Model
	NamaPasien string `gorm:"type:varchar(100)"`
	SequenceSample string `gorm:"type:text"`
	PenyakitID uint
	Similarity float64
	Penyakit Penyakit
}

func (Pemeriksaan) TableName() string {
	return "pemeriksaan"
}
