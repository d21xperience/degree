package template

import (
	"fmt"

	"gorm.io/gorm"
)

//	func GetNilaiColumns() []TemplateColumn {
//		return []TemplateColumn{
//			{Name: "Nama", Example: ""},
//			{Name: "Nama Kelas", Example: "", ColumnWidth: 10},
//			// {Name: "Tingkat Pendidikan", Example: "10", ColumnWidth: 8},
//			{Name: "Semester", Example: "", ColumnWidth: 10},
//		}
//	}
func GetNilaiTemplateColumns(dataTemplate *DataTemplate, db *gorm.DB) ([]TemplateColumn, bool, error) {
	// Ambil mata pelajaran dari database
	var mataPelajaran []struct {
		MataPelajaranId     string `gorm:"column:mata_pelajaran_id"`
		NmMapel             string `gorm:"column:nm_mapel"`
		TingkatPendidikanId string `gorm:"column:tingkat_pendidikan"`
		// Urutan string
	}

	// Query untuk mendapatkan mata pelajaran berdasarkan rombongan belajar
	tx := db.Session(&gorm.Session{NewDB: true})
	if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", dataTemplate.Schemaname)).Error; err != nil {
		return nil, false, err
	}

	if err := tx.Table("tabel_kategori_mapel mp").
		Select("mp.mata_pelajaran_id, mp.nm_mapel").
		// Joins("JOIN rombongan_belajar rb ON mp.rombel_id = rb.id").
		Where("mp.kurikulum_id = ? AND mp.tahun_ajaran_id = ? AND mp.tingkat_pendidikan = ?",
			dataTemplate.KurikulumId,
			dataTemplate.TahunAjaranId,
			dataTemplate.TingkatPendidikanId,
		).
		Order("mp.nm_mapel ASC").
		Find(&mataPelajaran).Error; err != nil {
		return nil, false, err
	}

	// Template dasar
	columns := []TemplateColumn{
		{Name: "Nama", Example: ""},
		{Name: "Nama Kelas", Example: "", ColumnWidth: 10},
		{Name: "Semester", Example: "", ColumnWidth: 10},
	}

	// Tambahkan kolom untuk setiap mata pelajaran
	for _, mp := range mataPelajaran {
		columns = append(columns, TemplateColumn{
			Name:        mp.MataPelajaranId,
			Example:     "",
			ColumnWidth: 16,
		})
		columns = append(columns, TemplateColumn{
			Name:        mp.NmMapel,
			Example:     "",
			ColumnWidth: 16,
		})
	}

	// Tambahkan kolom untuk nilai akhir dan rata-rata
	columns = append(columns, []TemplateColumn{
		{Name: "Nilai Akhir", Example: "", ColumnWidth: 12},
		{Name: "Rata-rata", Example: "", ColumnWidth: 12},
	}...)

	return columns, true, nil
}
