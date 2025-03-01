package models

import "gorm.io/datatypes"

type Category struct {
	Id_Category    int            `json:"id_kategori" gorm:"primary_key"`
	Category       string         `json:"kategori"`
	SubCategory    string         `json:"sub_kategori"`
	Keterangan     string         `json:"keterangan"`
	JumlahAset     string         `json:"jumlah_aset"`
	StatusCategory string         `json:"status_kategori"`
	MasaManfaat    string         `json:"masa_manfaat"`
	Penyusutan     string         `json:"penyusutan_persen_pertahun"`
	InsUser        string         `json:"ins_user"`
	InsDate        datatypes.Time `json:"ins_date"`
	UpdUser        string         `json:"upd_user"`
	UpdDate        datatypes.Time `json:"upd_date"`
}
