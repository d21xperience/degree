package template

import (
	"github.com/xuri/excelize/v2"
)

type TemplateColumn struct {
	Name        string
	Example     string
	Validation  *excelize.DataValidation
	FormatStyle *excelize.Style // e.g., {"number_format": 14} for date
	ColumnWidth float64
}

// Mengembalikan daftar kolom berdasarkan tipe template
func GetTemplateColumns(templateType string) ([]TemplateColumn, bool) {
	switch templateType {
	case "siswa":
		return GetSiswaColumns(), true
	case "guru":
		return GetGuruColumns(), true
	case "kelas":
		return GetKelasColumns(), true
	// case "ijazah":
	// 	return GetIjazahColumns(), true
	// case "nilai_akhir":
	// 	return GetNilaiAkhirColumns(), true
	default:
		return nil, false
	}
}

// func getSekolah (){

// }