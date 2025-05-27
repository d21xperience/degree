package template

import (
	"context"
	"fmt"
	"sekolah/repositories"
	"strings"
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
func GetTemplateColumns(param *DataTemplate, db *gorm.DB) ([]TemplateColumn, bool) {
	switch param.TemplateType {
	case "siswa":
		return GetSiswaColumns(), true
	case "guru":
		return GetGuruColumns(), true
	case "kelas":
		initKelas(param, db)
		return GetKelasColumns(), true
	// case "ijazah":
	// 	return GetIjazahColumns(), true
	case "nilai_akhir":
		return GetNilaiAkhirColumns(), true
	default:
		return nil, false
	}
}

func initKelas(param *DataTemplate, db *gorm.DB) error {
	// var err error
	repoKategoriSekolah := repositories.NewKategoriSekolahRepository(db)
	exactConditions := map[string]any{
		"tabel_kategori_sekolah.tahun_ajaran_id": param.TahunAjaranId,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	kategoriSekolah, err := repoKategoriSekolah.FindWithRelations(ctx, param.Schemaname, nil, nil, exactConditions, nil, nil, nil)
	if err != nil {
		return err
	}
	namaKurikulumList := []string{}
	namaJurusanList := []string{}
	for _, k := range kategoriSekolah {
		if k.NamaKurikulum != nil {
			namaKurikulumList = append(namaKurikulumList, *k.NamaKurikulum)
		}
		if k.NamaJurusan != nil {
			namaJurusanList = append(namaJurusanList, *k.NamaJurusan)
		}
	}

	KurikulumList = fmt.Sprintf("`\"%s\"`", strings.Join(namaKurikulumList, ","))
	fmt.Println(KurikulumList)
	JurusanList = fmt.Sprintf("`\"%s\"`", strings.Join(namaJurusanList, ","))
	// KurikulumList = `"` + strings.Join(namaKurikulumList, ",") + `"`
	// JurusanList = `"` + strings.Join(namaJurusanList, ",") + `"`

	return nil
}
