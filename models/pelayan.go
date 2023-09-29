package models

import (
	"github.com/jinzhu/gorm"
	"github.com/josepnapitupulu/API_Pelayan/config"
	)

var db *gorm.DB

type Pelayan struct {
	gorm.Model
	NIK string `gorm:""json:"nik"`
	NamaPelayan string `json:"nama_pelayan"`
	Tahbisan string `json:"tahbisan"`
	Pendidikan string `json:"pendidikan"`
	BidangPendidikan string `json:"bidang_pendidikan"`
	Pekerjaan string `json:"pekerjaan"`
	Alamat string `json:"alamat"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Pelayan{})
}

func (b *Pelayan) CreatePelayan() *Pelayan {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllPelayans() []Pelayan {
	var Pelayans []Pelayan
	db.Find(&Pelayans)
	return Pelayans
}

func GetPelayanbyId(nik int64) (*Pelayan, *gorm.DB) {
	var getPelayan Pelayan
	db := db.Where("NIK=?", nik).Find(&getPelayan)
	return &getPelayan, db
}

func DeletePelayan(nik int64) Pelayan {
	var pelayan Pelayan
	db.Where("NIK=?", nik).Delete(pelayan)
	return pelayan 
}