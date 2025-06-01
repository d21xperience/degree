package models

import (
	"time"

	"github.com/google/uuid"
)

type TabelPTK struct {
	PtkID             uuid.UUID  `gorm:"column:ptk_id;primaryKey"` // UUID
	Nama              string     `gorm:"column:nama"`              // String
	JenisKelamin      *string    `gorm:"column:jenis_kelamin"`     // String
	Agama             *string    `gorm:"column:agama"`             // String
	TempatLahir       *string    `gorm:"column:tempat_lahir"`      // String
	TanggalLahir      *time.Time `gorm:"column:tanggal_lahir"`     // String (format tanggal, bisa diubah ke time.Time jika perlu)
	IsDapodik         bool       `gorm:"column:is_dapodik"`
	PtkIdDapodik      *uuid.UUID `gorm:"column:ptk_id_dapodik"`
	StatusKeaktifanID int32      `gorm:"column:status_keaktifan_id"` // String
	JenisPtkID        int32      `gorm:"column:jenis_ptk_id"`        // String
}

type PTKTerdaftar struct {
	PtkTerdaftarId        uuid.UUID  `gorm:"column:ptk_terdaftar_id;primaryKey"`
	PtkID                 uuid.UUID  `gorm:"column:ptk_id"`
	TahunAjaranId         string     `gorm:"column:tahun_ajaran_id"`
	JenisKeluarId         *string    `gorm:"column:jenis_keluar_id"`
	IsDapodik             bool       `gorm:"column:is_dapodik"`
	PtkTerdaftarIdDapodik *uuid.UUID `gorm:"column:ptk_terdaftar_id_dapodik"`
	// Relasi ke PTK
	PTK TabelPTK `gorm:"foreignKey:PtkID;references:PtkID"`
	// Pembelajaran []Pembelajaran `gorm:"foreignKey:PtkTerdaftarId;references:PtkTerdaftarId"`
	PTKPelengkap PtkPelengkap `gorm:"foreignKey:PtkID;references:PtkId"` // Menyertakan PTKPelengkap
}

// Menentukan nama tabel kustom
func (PTKTerdaftar) TableName() string {
	return "tabel_ptk_terdaftar"
}

// Menentukan nama tabel kustom
func (TabelPTK) TableName() string {
	return "tabel_ptk"
}

type PtkPelengkap struct {
	PtkPelengkapId uuid.UUID `gorm:"column:ptk_pelengkap_id"`
	PtkId          uuid.UUID `gorm:"column:ptk_id"`
	GelarDepan     *string   `gorm:"column:gelar_depan"`
	GelarBelakang  *string   `gorm:"column:gelar_belakang"`
	Nik            *string   `gorm:"column:nik"`
	NoKk           *string   `gorm:"column:no_kk"`
	Nuptk          *string   `gorm:"column:nuptk"`
	Niy            *string   `gorm:"column:niy"`
	Nip            *string   `gorm:"column:nip"`
	AlamatJalan    *string   `gorm:"column:alamat_jalan"`
	Rt             *string   `gorm:"column:rt"`
	Rw             *string   `gorm:"column:rw"`
	DesaKelurahan  *string   `gorm:"column:desa_kelurahan"`
	Kec            *string   `gorm:"column:kec"`
	KabKota        *string   `gorm:"column:kab_kota"`
	Propinsi       *string   `gorm:"column:propinsi"`
	KodePos        *string   `gorm:"column:kode_pos"`
	NoTeleponRumah *string   `gorm:"column:no_telepon_rumah"`
	NoHp           *string   `gorm:"column:no_hp"`
	Email          *string   `gorm:"column:email"`

	// Relasi opsional ke tabel PTK
	Ptk *TabelPTK `gorm:"foreignKey:PtkId;references:PtkID"`
}

func (PtkPelengkap) TableName() string {
	return "tabel_ptk_pelengkap"
}
