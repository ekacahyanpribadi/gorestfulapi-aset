package models

import "gorm.io/datatypes"

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by a to b
func (Kategori_aset) TableName() string {
	return "kategori_aset"
}

type Kategori_aset struct {
	Id_kategori                string `json:"id_kategori" gorm:"primary_key"`
	Kategori                   string `json:"kategori"`
	Sub_kategori               string `json:"sub_kategori"`
	Keterangan                 string `json:"keterangan"`
	Jumlah_aset                string `json:"jumlah_aset"`
	Status_kategori            string `json:"status_kategori"`
	Masa_manfaat               string `json:"masa_manfaat"`
	Penyusutan_persen_pertahun string `json:"penyusutan_persen_pertahun"`
	Ins_user                   string `json:"ins_user"`
	Ins_date                   string `json:"ins_date"`
	Upd_user                   string `json:"upd_user"`
	Upd_date                   string `json:"upd_date"`
}

type kaliKepake struct {
	UpdDate datatypes.Time `json:"upd_date"`
}

type Token_access struct {
	Id_token    string `json:"id_token" gorm:"primary_key"`
	Token       string `json:"token"`
	Desc1_token string `json:"desc1_token"`
	Desc2_token string `json:"desc2_token"`
}

// type validation kategori input
type ValidateKategoriInput struct {
	Id_kategori                string `json:"id_kategori" binding:"required"`
	Kategori                   string `json:"kategori" binding:"required"`
	Sub_kategori               string `json:"sub_kategori" binding:"required"`
	Keterangan                 string `json:"keterangan" binding:"required"`
	Jumlah_aset                string `json:"jumlah_aset" binding:"required"`
	Status_kategori            string `json:"status_kategori" binding:"required"`
	Masa_manfaat               string `json:"masa_manfaat" binding:"required"`
	Penyusutan_persen_pertahun string `json:"penyusutan_persen_pertahun" binding:"required"`
	Ins_user                   string `json:"ins_user" binding:"required"`
	Ins_date                   string `json:"ins_date" binding:"required"`
	Upd_user                   string `json:"upd_user" binding:"required"`
	Upd_date                   string `json:"upd_date" binding:"required"`
}

// type error message
type ErrorMsgKategori struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type mKuntul struct {
	UpdDate datatypes.Time `json:"upd_date"`
}
