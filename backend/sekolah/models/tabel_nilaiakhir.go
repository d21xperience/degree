package models

import (
	"github.com/google/uuid"
)

type NilaiAkhir struct {
	IdNilaiAkhir    uuid.UUID `gorm:"column:id_nilai_akhir;primaryKey"` // Primary key
	AnggotaRombelId uuid.UUID `gorm:"column:anggota_rombel_id"`         // Foreign key ke tabel anggota_rombel
	MataPelajaranId *uint32   `gorm:"column:mata_pelajaran_id"`         // Foreign key ke tabel mata_pelajaran
	SemesterId      string    `gorm:"column:semester_id"`               // Semester
	NilaiPeng       *uint32   `gorm:"column:nilai_peng"`                // Nilai Pengetahuan
	PredikatPeng    string    `gorm:"column:predikat_peng"`             // Predikat Pengetahuan
	NilaiKet        *uint32   `gorm:"column:nilai_ket"`                 // Nilai Keterampilan
	PredikatKet     string    `gorm:"column:predikat_ket"`              // Predikat Keterampilan
	NilaiSik        *uint32   `gorm:"column:nilai_sik"`                 // Nilai Sikap
	PredikatSik     string    `gorm:"column:predikat_sik"`              // Predikat Sikap
	NilaiSikSos     *uint32   `gorm:"column:nilai_siksos"`              // Nilai Sikap Sosial
	PredikatSikSos  string    `gorm:"column:predikat_siksos"`           // Predikat Sikap Sosial
	PesertaDidikId  uuid.UUID `gorm:"column:peserta_didik_id"`          // Foreign key ke tabel peserta_didik
	IDMinat         string    `gorm:"column:id_minat"`                  // ID Minat
	Semester        *uint32   `gorm:"column:semester"`                  // Semester

	// Relasi
	// PesertaDidik     PesertaDidik //`gorm:"foreignKey:PesertaDidikId;references:PesertaDidikId"`
	// RombonganBelajar RombonganBelajar
	MataPelajaran MataPelajaran //`gorm:"foreignKey:MataPelajaranId;references:MataPelajaranID"`
	// AnggotaRombel RombelAnggota `gorm:"foreignKey:AnggotaRombelId;references:AnggotaRombelId"`
}

func (NilaiAkhir) TableName() string {
	return "tabel_nilaiakhir"
}

type SiswaNilai struct {
	MataPelajaran string `json:"mataPelajaran"`
	Semester1     *int   `json:"semester1,omitempty"`
	Semester2     *int   `json:"semester2,omitempty"`
	Semester3     *int   `json:"semester3,omitempty"`
	Semester4     *int   `json:"semester4,omitempty"`
	Semester5     *int   `json:"semester5,omitempty"`
	Semester6     *int   `json:"semester6,omitempty"`
	Semester7     *int   `json:"semester7,omitempty"`
	Semester8     *int   `json:"semester8,omitempty"`
	Semester9     *int   `json:"semester9,omitempty"`
	Semester10    *int   `json:"semester10,omitempty"`
	Semester11    *int   `json:"semester11,omitempty"`
	Semester12    *int   `json:"semester12,omitempty"`
}
