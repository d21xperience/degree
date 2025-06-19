package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Sekolah struct {
	SekolahID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"sekolah_id"`
	Nama                string    `gorm:"type:varchar(100);not null" json:"nama"`
	NPSN                string    `gorm:"type:varchar(8);default:null" json:"npsn,omitempty"`
	NSS                 string    `gorm:"type:varchar(12);default:null" json:"nss,omitempty"`
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
	StatusKepemilikanID *uint32   `gorm:"column:status_kepemilikan_id;type:numeric(1,0);default:null" json:"status_kepemilikan_id,omitempty"`
	KodeAktivasi        *string   `gorm:"type:varchar(30);default:null" json:"kode_aktivasi,omitempty"`
	BentukPendidikanID  *uint32   `gorm:"column:bentuk_pendidikan_id;type:smallint;default:null" json:"bentuk_pendidikan_id,omitempty"`
	JenjangPendidikanID *uint32   `gorm:"column:jenjang_pendidikan_id;type:numeric(2,0);default:null" json:"jenjang_pendidikan_id,omitempty"`
	LamaPendidikan      *uint32   `gorm:"column:lama_pendidikan"`

	BentukPendidikan  BentukPendidikan  `gorm:"foreignKey:BentukPendidikanID;references:BentukPendidikanID"`
	JenjangPendidikan JenjangPendidikan `gorm:"foreignKey:JenjangPendidikanID;references:JenjangPendidikanID"`
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

type KategoriSekolah struct {
	KategorisekolahId   int32   `gorm:"column:kategori_sekolah_id;primaryKey;autoIncrement"`
	KurikulumId         int32   `gorm:"column:kurikulum_id"`
	JurusanId           *string `gorm:"column:jurusan_id"`
	NamaKurikulum       *string `gorm:"column:nama_kurikulum"`
	NamaBidangKeahlian  *string `gorm:"column:nama_bidang_keahlian"`
	NamaProgramKeahlian *string `gorm:"column:nama_program_keahlian"`
	NamaJurusan         *string `gorm:"column:nama_jurusan"`
	JenjangPendidikanId *int32  `gorm:"column:jenjang_pendidikan_id"`
	TingkatId           int32   `gorm:"column:tingkat_id"`
	Jumlah              *int32  `gorm:"jumlah"`
	TahunAjaranId       int32   `gorm:"column:tahun_ajaran_id"`
}
type KategoriMapel struct {
	Id                int32          `gorm:"column:id;primaryKey;autoIncrement"`
	KategorisekolahId int32          `gorm:"column:kategori_sekolah_id"`
	MataPelajaranId   int32          `gorm:"column:mata_pelajaran_id"`
	KurikulumId       int32          `gorm:"column:kurikulum_id"`
	JurusanId         string         `gorm:"jurusan_id"`
	NmMapel           string         `gorm:"nama_mapel"`
	TingkatPendidikan string         `gorm:"tingkat_pendidikan"`
	TahunAjaranId     string         `gorm:"column:tahun_ajaran_id"`
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}

type KategoriSekolahLog struct {
	LogId             int32   `gorm:"column:log_id;primaryKey;autoIncrement"`
	ActionType        *string `gorm:"column:action_type"`
	KategorisekolahId int32   `gorm:"column:kategori_sekolah_id"`
	NamaKurikulum     *string `gorm:"column:nama_kurikulum"`
	NamaJurusan       *string `gorm:"column:nama_jurusan"`
	JurusanId         *string `gorm:"column:jurusan_id"`
	KurikulumId       *int32  `gorm:"column:kurikulum_id"`
	TahunAjaranId     int32   `gorm:"column:tahun_ajaran_id"`
}

func (Sekolah) TableName() string {
	return "tabel_sekolah"
}
func (KategoriSekolah) TableName() string {
	return "tabel_kategori_sekolah"
}
func (KategoriMapel) TableName() string {
	return "tabel_kategori_mapel"
}
