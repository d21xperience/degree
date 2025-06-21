package template

import (
	"context"
	"fmt"
	"sekolah/repositories"
	"time"

	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

var KurikulumList string
var JurusanList string

type TemplateColumn struct {
	Name        string
	Example     string
	Validation  *excelize.DataValidation
	FormatStyle *excelize.Style // e.g., {"number_format": 14} for date
	ColumnWidth float64
}

type DataTemplate struct {
	TemplateType  string
	Schemaname    string
	TahunAjaranId string
}

// Mengembalikan daftar kolom berdasarkan tipe template
func GetTemplateColumns(param *DataTemplate) ([]TemplateColumn, bool) {
	switch param.TemplateType {
	case "siswa":
		return GetSiswaColumns(), true
	case "guru":
		return GetGuruColumns(), true
	case "kelas":
		return GetKelasColumns(), true
	// case "ijazah":
	// 	return GetIjazahColumns(), true
	case "nilai_akhir":
		return GetNilaiAkhirColumns(), true
	default:
		return nil, false
	}
}

// func initKelas(param *DataTemplate, db *gorm.DB) error {
// 	// var err error
// 	repoKategoriSekolah := repositories.NewKategoriSekolahRepository(db)
// 	exactConditions := map[string]any{
// 		"tabel_kategori_sekolah.tahun_ajaran_id": param.TahunAjaranId,
// 	}
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	kategoriSekolah, err := repoKategoriSekolah.FindWithRelations(ctx, param.Schemaname, nil, nil, exactConditions, nil, nil, nil)
// 	if err != nil {
// 		return err
// 	}
// 	// namaKurikulumList := []string{}
// 	namaJurusanList := []string{}
// 	for _, k := range kategoriSekolah {
// 		// if k.NamaKurikulum != nil {
// 		// 	namaKurikulumList = append(namaKurikulumList, *k.NamaKurikulum)
// 		// }
// 		// if k.NamaJurusan != nil {
// 		// 	namaJurusanList = append(namaJurusanList, *k.NamaJurusan)
// 		// }
// 		namaJurusanList = append(namaJurusanList, fmt.Sprintf("%d untuk jurusan", k.KategorisekolahId), *k.NamaJurusan)
// 	}

// 	// KurikulumList = fmt.Sprintf("`\"%s\"`", strings.Join(namaKurikulumList, ","))
// 	// fmt.Println(KurikulumList)
// 	// JurusanList = fmt.Sprintf("`\"%s\"`", strings.Join(namaJurusanList, ";"))
// 	JurusanList = strings.Join(namaJurusanList, ";")
// 	// KurikulumList = `"` + strings.Join(namaKurikulumList, ",") + `"`
// 	// JurusanList = `"` + strings.Join(namaJurusanList, ",") + `"`

// 	return nil
// }

func GetPetunjukSheet(f *excelize.File, param *DataTemplate, db *gorm.DB) error {
	switch param.TemplateType {
	case "siswa":
		return getKelas(f, param, db)
	// case "guru":
	// 	return GetGuruColumns()
	// case "kelas":
	// 	return GetKelasColumns()
	// // case "ijazah":
	// // 	return GetIjazahColumns()
	// case "nilai_akhir":
	// 	return getNilaiAkhir(f, param, db)
	default:
		return nil
	}
}

// revisi 1
// func getNilaiAkhir(f *excelize.File, param *DataTemplate, db *gorm.DB) error {
// 	repoKelas := repositories.NewrombonganBelajarRepository(db)
// 	repoSiswa := repositories.NewSiswaRepository(db)
// 	repoMapel := repositories.NewKategoriMapelRepository(db)
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	conditions := map[string]any{
// 		"tabel_kelas.semester_id": fmt.Sprintf("%s1", param.TahunAjaranId),
// 	}
// 	resp
// 	kelasList, err := repoKelas.FindAllByConditions(ctx, param.Schemaname, conditions, 100, 0, nil)
// 	if err != nil {
// 		return err
// 	}
// 	mapelList, err := repoMapel.FindAllByConditions(ctx, param.Schemaname, nil, 0, 1000, nil)
// 	if err != nil {
// 		return err
// 	}

// 	infoSheet := "Nilai Akhir"
// 	f.NewSheet(infoSheet)

// 	// Sembunyikan kolom G
// 	f.SetColVisible(infoSheet, "C", false)

// 	f.SetCellValue(infoSheet, "A1", "Peserta Didik ID")
// 	f.SetCellValue(infoSheet, "B1", "Nama")
// 	f.SetCellValue(infoSheet, "C1", "Mapel")
// 	for i, mapel := range mapelList {
// 		kolom, err := excelize.ColumnNumberToName(i + 3)
// 		if err != nil {
// 			return err
// 		}
// 		for j := 0; j < len(mapelList); j++ {
// 			cell := fmt.Sprintf("%s%d",kolom, j)
// 			f.SetCellValue(infoSheet, cell, mapel.NmMapel)
// 		}
// 	}
// 	var cell string
// 	for i, kelas := range siswaList {
// 		cell = fmt.Sprintf("A%d", i+2)
// 		f.SetCellValue(infoSheet, cell, kelas.NmKelas)

// 		cell = fmt.Sprintf("B%d", i+2)
// 		f.SetCellValue(infoSheet, cell, kelas.TingkatPendidikanId)

//			cell = fmt.Sprintf("C%d", i+2)
//			f.SetCellValue(infoSheet, cell, kelas.RombonganBelajarId)
//		}
//		return nil
//	}
func getKelas(f *excelize.File, param *DataTemplate, db *gorm.DB) error {
	repoKelas := repositories.NewrombonganBelajarRepository(db)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conditions := map[string]any{
		"tabel_kelas.semester_id": fmt.Sprintf("%s1", param.TahunAjaranId),
	}
	kelasList, err := repoKelas.FindAllByConditions(ctx, param.Schemaname, conditions, 100, 0, nil)
	if err != nil {
		return err
	}
	infoSheet := "Petunjuk pengisian"
	f.NewSheet(infoSheet)

	// Sembunyikan kolom G
	f.SetColVisible(infoSheet, "C", false)

	f.SetCellValue(infoSheet, "A1", "Nama Kelas")
	f.SetCellValue(infoSheet, "B1", "Tingkat")
	f.SetCellValue(infoSheet, "C1", "Rombel Id")

	var cell string
	for i, kelas := range kelasList {
		cell = fmt.Sprintf("A%d", i+2)
		f.SetCellValue(infoSheet, cell, kelas.NmKelas)

		cell = fmt.Sprintf("B%d", i+2)
		f.SetCellValue(infoSheet, cell, kelas.TingkatPendidikanId)

		cell = fmt.Sprintf("C%d", i+2)
		f.SetCellValue(infoSheet, cell, kelas.RombonganBelajarId)
	}
	return nil
}
