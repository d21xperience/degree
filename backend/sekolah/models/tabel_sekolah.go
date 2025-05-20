package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Sekolah struct {
	SekolahID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"sekolah_id"`
	Nama                string    `gorm:"type:varchar(100);not null" json:"nama"`
	NPSN                string   `gorm:"type:varchar(8);default:null" json:"npsn,omitempty"`
	NSS                 string   `gorm:"type:varchar(12);default:null" json:"nss,omitempty"`
	Alamat              *string   `gorm:"type:text;default:null" json:"alamat,omitempty"`
	KodePos             *string   `gorm:"column:kd_pos;type:varchar(6);default:null" json:"kd_pos,omitempty"`
	Telepon             *string   `gorm:"type:varchar(20);default:null" json:"telepon,omitempty"`
	Fax                 *string   `gorm:"type:varchar(20);default:null" json:"fax,omitempty"`
	Kelurahan           *string   `gorm:"type:varchar(60);default:null" json:"kelurahan,omitempty"`
	Kecamatan           *string   `gorm:"type:varchar(60);default:null" json:"kecamatan,omitempty"`
	KabKota             *string   `gorm:"column:kab_kota;type:varchar(60);default:null" json:"kab_kota,omitempty"`
	Propinsi            *string   `gorm:"type:varchar(60);default:null" json:"propinsi,omitempty"`
	Website             *string   `gorm:"type:varchar(100);default:null" json:"website,omitempty"`
	Email               *string   `gorm:"type:varchar(50);default:null" json:"email,omitempty"`
	NamaKepsek          *string   `gorm:"column:nm_kepsek;type:varchar(100);default:null" json:"nm_kepsek,omitempty"`
	NIPKepsek           *string   `gorm:"column:nip_kepsek;type:varchar(25);default:null" json:"nip_kepsek,omitempty"`
	NIYKepsek           *string   `gorm:"column:niy_kepsek;type:varchar(30);default:null" json:"niy_kepsek,omitempty"`
	StatusKepemilikanID *int32    `gorm:"column:status_kepemilikan_id;type:numeric(1,0);default:null" json:"status_kepemilikan_id,omitempty"`
	KodeAktivasi        *string   `gorm:"type:varchar(30);default:null" json:"kode_aktivasi,omitempty"`
	BentukPendidikanID  *int32   `gorm:"column:bentuk_pendidikan_id;type:smallint;default:null" json:"bentuk_pendidikan_id,omitempty"`
	JenjangPendidikanID *int32    `gorm:"column:jenjang_pendidikan_id;type:numeric(2,0);default:null" json:"jenjang_pendidikan_id,omitempty"`
}

type SekolahTenant struct {
	gorm.Model
	NamaSekolah     string `gorm:"column:nama_sekolah"`
	SekolahTenantId uint32 `gorm:"column:sekolah_tenant_id"`
	SchemaName      string `gorm:"column:schema_name"`
	// BentukPendidikanID uint16           `gorm:"not null;default:0;column:bentuk_pendidikan_id"`
	// BentukPendidikan   BentukPendidikan `gorm:"foreignKey:BentukPendidikanID;references:BentukPendidikanID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

type SchemaLog struct {
	gorm.Model
	SchemaName string
}

// func (SekolahTenant) TableName() string {
// 	return "sekolah_tenant"
// }
