package services

import (
	"context"
	"fmt"
	"sekolah/models"
	"sekolah/repositories"
	"sekolah/template"
	"sekolah/utils"

	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

type ParamTemplate struct {
	templateType string
	schemaname   string
	filePath     string
	semesterId   string
	
}

// // RequestRequirement digunakan untuk menyimpan dependensi service dan protobuf
// type RequestRequirement[S any, P any] struct {
// 	service  S // Service untuk mengakses data
// 	protoBuf P // Protobuf atau struct untuk representasi data
// }

// func (r *RequestRequirement[S, P]) GetModelByID(ctx context.Context, id string, fetcher func(S, string) (*P, error)) (*P, error) {
// 	// Gunakan fetcher untuk mengambil data berdasarkan ID
// 	result, err := fetcher(r.service, id)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to get model by ID: %w", err)
// 	}

// 	return result, nil
// }

// func ConvertModelsToPB[T any, U any](models []*T, converter func(*T) *U) []*U {
// 	var pbList []*U
// 	for _, model := range models {
// 		pbList = append(pbList, converter(model))
// 	}
// 	return pbList
// }
// func ConvertPBToModels[T any, U any](pbs []*T, converter func(*T) *U) []*U {
// 	var modelList []*U
// 	for _, model := range pbs {
// 		modelList = append(modelList, converter(model))
// 	}
// 	return modelList
// }

// func ConvertModelToPB[T any, U any](model *T, converter func(*T) *U) *U {
// 	if model == nil {
// 		return nil
// 	}
// 	return converter(model)
// }

// Fungsi untuk membaca file Excel dan memproses data berdasarkan jenis
func BacaDataExcel(param *ParamTemplate) ([][]string, error) {
	f, err := excelize.OpenFile(param.filePath)
	if err != nil {
		return nil, fmt.Errorf("gagal membaca file Excel: %w", err)
	}
	defer f.Close()

	ret, err := f.GetDocProps()
	if err != nil {
		return nil, err
	}
	param.semesterId = ret.Category
	// param.schemaname = ret.Keywords
	rows, err := f.GetRows(f.GetSheetName(0))
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil data dari sheet: %w", err)
	}

	if len(rows) < 2 {
		return nil, fmt.Errorf("file Excel kosong atau tidak memiliki data yang valid")
	}
	// Kembalikan data mulai dari baris kedua (tanpa header)
	return rows[1:], nil
	// return rows, nil

}

// func BacaDataExcel(param *ParamTemplate) ([][]string, error) {
// 	f, err := excelize.OpenFile(param.filePath)
// 	if err != nil {
// 		return nil, fmt.Errorf("gagal membaca file Excel: %w", err)
// 	}
// 	defer f.Close()

// 	ret, err := f.GetDocProps()
// 	if err != nil {
// 		return nil, err
// 	}
// 	param.semesterId = ret.Category

// 	sheetName := f.GetSheetName(0)

// 	rows, err := f.GetRows(sheetName, excelize.Options{RawCellValue: true})
// 	if err != nil {
// 		return nil, fmt.Errorf("gagal mengambil data dari sheet: %w", err)
// 	}

// 	if len(rows) < 2 {
// 		return nil, fmt.Errorf("file Excel kosong atau tidak memiliki data yang valid")
// 	}

// 	var result [][]string

// 	for i, row := range rows {
// 		if i == 0 {
// 			// Skip header
// 			result = append(result, row)
// 			continue
// 		}

// 		var newRow []string
// 		for j, cell := range row {
// 			cellRef := fmt.Sprintf("%s%d", ToAlpha(j+1), i+1)
// 			rawVal, err := f.GetCellValue(sheetName, cellRef)
// 			if err != nil {
// 				newRow = append(newRow, cell)
// 				continue
// 			}

// 			// Coba konversi nilai ke float64, untuk deteksi serial date
// 			dateFloat, ok := isExcelDate(rawVal)
// 			if ok {
// 				t, parseErr := excelize.ExcelDateToTime(dateFloat, false)
// 				if parseErr == nil {
// 					rawVal = t.Format("02/01/2006") // Format DD/MM/YYYY
// 				}
// 			}

// 			newRow = append(newRow, rawVal)
// 		}

// 		result = append(result, newRow)
// 	}

// 	return result[1:], nil
// }

// // Helper function untuk cek apakah string bisa jadi angka dan merepresentasikan Excel Date
// func isExcelDate(val string) (float64, bool) {
// 	var f float64
// 	_, err := fmt.Sscan(val, &f)
// 	if err != nil {
// 		return 0, false
// 	}
// 	// Excel dates biasanya antara 1 sampai ~2958465
// 	if f > 0 && f < 2958465 {
// 		return f, true
// 	}
// 	return f, false
// }

// // Fungsi bantu untuk ubah nomor kolom ke huruf (A, B, ..., Z, AA, dll.)
// func ToAlpha(colNumber int) string {
// 	var col string
// 	for colNumber > 0 {
// 		colNumber--
// 		letter := 'A' + rune(colNumber%26)
// 		col = string(letter) + col
// 		colNumber /= 26
// 	}
// 	return col
// }

// Fungsi generik untuk membaca file Excel dan memproses data berdasarkan jenis
func UploadDataSekolah[T any](param ParamTemplate) ([]T, error) {
	f, err := excelize.OpenFile(param.filePath)
	if err != nil {
		return nil, fmt.Errorf("gagal membaca file Excel: %w", err)
	}
	defer f.Close()

	ret, err := f.GetDocProps()
	if err != nil {
		return nil, err
	}
	// ContentStatus: param.templateType,
	// Category:      param.semesterId,
	// Keywords:      param.schemaname,
	if ret.Keywords != param.schemaname || ret.ContentStatus != param.templateType {
		return nil, err
	}

	rows, err := f.GetRows(f.GetSheetName(0))
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil data dari sheet: %w", err)
	}

	if len(rows) < 2 {
		return nil, fmt.Errorf("file Excel kosong atau tidak memiliki data yang valid")
	}

	// Parsing data berdasarkan uploadType
	switch param.templateType {
	case "siswa":
		if v, ok := any(parseSiswa(rows)).([]T); ok {
			return v, nil
		}
	case "guru":
		if v, ok := any(parseGuru(rows)).([]T); ok {
			return v, nil
		}
	case "kelas":
		if v, ok := any(parseKelas(rows)).([]T); ok {
			return v, nil
		}
	case "nilai_akhir":
		if v, ok := any(parseNilaiAkhir(rows)).([]T); ok {
			return v, nil
		}
	case "ijazah":
		if v, ok := any(parseNilaiAkhir(rows)).([]T); ok {
			return v, nil
		}
	default:
		return nil, fmt.Errorf("jenis unggahan tidak dikenali: %s", param.templateType)
	}

	return nil, fmt.Errorf("gagal memproses data dengan tipe yang diberikan")
}

// Fungsi parsing untuk siswa
func parseSiswa(rows [][]string) []models.PesertaDidik {
	var siswaList []models.PesertaDidik
	for _, row := range rows[1:] {
		if len(row) < 3 {
			continue
		}
		tglLahir, err := utils.StringToTime(row[4], "2006-01-02")
		if err != nil {
			return nil
		}
		tglDiterima, err := utils.StringToTime(row[4], "2006-01-02")
		if err != nil {
			return nil
		}
		siswa := models.PesertaDidik{
			Nis:             &row[0],
			Nisn:            &row[1],
			NmSiswa:         row[2],
			TempatLahir:     &row[3],
			TanggalLahir:    utils.TimeToPointer(tglLahir.Format("2006-01-02")),
			JenisKelamin:    &row[5],
			Agama:           &row[6],
			AlamatSiswa:     &row[7],
			TeleponSiswa:    &row[8],
			DiterimaTanggal: utils.TimeToPointer(tglDiterima.Format("2006-01-02")),
			NmAyah:          &row[10],
			NmIbu:           &row[11],
			PekerjaanAyah:   &row[12],
			PekerjaanIbu:    &row[13],
			NmWali:          &row[14],
			PekerjaanWali:   &row[15],
			Nik:             &row[16],
			// status_dalam_kel: &row[17],
			// Anak_ke:          &row[18],
			// sekolah_asal:     &row[19],
			// diterima_kelas:   &row[20],
			// alamat_ortu:      &row[21],
			// telepon_ortu:     &row[22],
			// alamat_wali:      &row[23],
			// telepon_wali:     &row[24],
		}
		siswaList = append(siswaList, siswa)
	}
	return siswaList
}

// Fungsi parsing untuk guru
func parseGuru(rows [][]string) []*models.TabelPTK {
	var guruList []*models.TabelPTK
	for _, row := range rows[1:] {
		if len(row) < 3 {
			continue
		}
		guru := &models.TabelPTK{
			// Nama:   row[0],
			// Mapel:  row[1],
			// Alamat: row[2],
		}
		guruList = append(guruList, guru)
	}
	return guruList
}

// Fungsi parsing untuk kelas
func parseKelas(rows [][]string) []*models.RombonganBelajar {
	var kelasList []*models.RombonganBelajar
	for _, row := range rows[1:] {
		if len(row) < 2 {
			continue
		}
		kelas := &models.RombonganBelajar{
			// NamaKelas:  row[0],
			// JumlahSiswa: parseInt(row[1]),
		}
		kelasList = append(kelasList, kelas)
	}
	return kelasList
}

// Fungsi parsing untuk kelas
func parseNilaiAkhir(rows [][]string) []*models.NilaiAkhir {
	var kelasList []*models.NilaiAkhir
	for _, row := range rows[1:] {
		if len(row) < 2 {
			continue
		}
		kelas := &models.NilaiAkhir{
			// AnggotaRombelID: (uuid.UUID).row[0],
			// NamaKelas:  row[0],
			// JumlahSiswa: parseInt(row[1]),
			// NilaiPeng: int32(parseInt(row[2])),
		}
		kelasList = append(kelasList, kelas)
	}
	return kelasList
}

// GenerateTemplate membuat template XLSX untuk berbagai jenis data
func GenerateTemplate(param ParamTemplate, db *gorm.DB) error {
	f := excelize.NewFile()
	sheetName := "Template"
	f.SetSheetName("Sheet1", sheetName)

	// Tentukan header berdasarkan jenis template
	var headers []string
	templates := map[string][]string{
		"siswa": {"NIS", "NISN", "Nama", "Tempat Lahir", "Tanggal Lahir", "JK", "Agama", "Alamat", "Telepon Siswa", "Diterima Tanggal", "Nama Ayah", "Nama Ibu", "Pekerjaan Ayah", "Pekerjaan Ibu", "Nama Wali", "Pekerjaan Wali", "NIK Siswa", "status_dalam_kel", "anak_ke", "sekolah_asal", "diterima_kelas", "alamat_ortu", "telepon_ortu", "alamat_wali", "telepon_wali"},

		"nilai_akhir": {"peserta_didik_id(uuid)", "nm_siswa", "mata_pelajaran_id"},
		"ijazah":      {"peserta_didik_id(uuid)", "nm_siswa", "nis", "nomor_ijazah", "tahun_lulus"},

		"kelas": {"Nama Kelas", "Nama Wali Kelas", "Tingkat", "Kurikulum ID"},

		"guru": {"nama", "jenis_kelamin", "tempat_lahir", "tanggal_lahir", "agama", "gelar_depan", "gelar_belakang", "nik", "no_kk", "nuptk", "niy", "nip", "alamat_jalan", "rt", "rw", "desa_kelurahan", "kec", "kab_kota", "propinsi", "kode_pos", "no_telepon_rumah", "no_hp", "email"},
	}
	headers, exists := templates[param.templateType]
	if !exists {
		return fmt.Errorf("template type %s tidak ditemukan", param.templateType)
	}

	for col, header := range headers {
		cell := fmt.Sprintf("%s1", getExcelColumnName(col))
		f.SetCellValue(sheetName, cell, header)
	}
	err := f.SetDocProps(&excelize.DocProperties{
		ContentStatus: param.templateType,
		Category:      param.semesterId,
		Keywords:      param.schemaname,
	})
	if err != nil {
		return fmt.Errorf("gagal membuat properties %s: %w", param.templateType, err)
	}
	// Simpan ke file
	err = f.SaveAs(param.filePath)
	if err != nil {
		return fmt.Errorf("gagal membuat template %s: %w", param.templateType, err)
	}

	return nil
}
func GenerateTemplateV2(param ParamTemplate, db *gorm.DB) error {
	f := excelize.NewFile()
	sheetName := "Template"
	f.SetSheetName("Sheet1", sheetName)

	columns, exists := template.GetTemplateColumns(param.templateType)
	if !exists {
		return fmt.Errorf("template type %s tidak ditemukan", param.templateType)
	}
	headerStyle, _ := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#4F81BD"},
			Pattern: 1,
		},
		Font: &excelize.Font{
			Bold:  true,
			Color: "#FFFFFF",
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
			WrapText:   true,
		},
	})
	err1 := f.SetRowHeight(sheetName, 1, 50)
	if err1 != nil {
		return err1
	}
	// Buat kolom header dan contoh input
	for i, col := range columns {
		colLetter, _ := excelize.ColumnNumberToName(i + 1)
		cellHeader := fmt.Sprintf("%s1", colLetter)
		cellExample := fmt.Sprintf("%s2", colLetter)

		// Header
		f.SetCellValue(sheetName, cellHeader, col.Name)
		f.SetCellStyle(sheetName, cellHeader, cellHeader, headerStyle)

		// Contoh isi baris kedua
		f.SetCellValue(sheetName, cellExample, col.Example)

		// Validasi (jika ada)
		if col.Validation != nil {
			f.AddDataValidation(sheetName, col.Validation)
		}

		// Format tanggal (jika ada)
		if col.FormatStyle != nil {
			styleID, _ := f.NewStyle(col.FormatStyle)
			f.SetCellStyle(sheetName, cellExample, fmt.Sprintf("%s100", colLetter), styleID)
		}
		if col.ColumnWidth != 0 {
			f.SetColWidth(sheetName, colLetter, colLetter, col.ColumnWidth)
		}
	}

	// Freeze baris header
	f.SetPanes(sheetName, &excelize.Panes{Freeze: true, Split: true, XSplit: 0, YSplit: 1})

	// Set properties dokumen
	err := f.SetDocProps(&excelize.DocProperties{
		ContentStatus: param.templateType,
		Category:      param.semesterId,
		Keywords:      param.schemaname,
	})
	if err != nil {
		return fmt.Errorf("gagal membuat properties %s: %w", param.templateType, err)
	}

	// Simpan ke file
	if err := f.SaveAs(param.filePath); err != nil {
		return fmt.Errorf("gagal menyimpan template %s: %w", param.templateType, err)
	}

	return nil
}

func templateNilaiAkhir(ctx context.Context, db *gorm.DB, f *excelize.File, param ParamTemplate, sheetName string) error {
	// TODO: Implementasi template nilai akhir
	// sheetName := "Template"
	f.SetSheetName("Sheet1", sheetName)
	repoAnggotaKelas := repositories.NewRombelAnggotaRepository(db)
	joins := []string{
		"JOIN tabel_siswa ON tabel_siswa.peserta_didik_id = tabel_anggotakelas.peserta_didik_id",
		// "JOIN ref.jurusan ON tabel_kelas.jurusan_id = ref.jurusan.jurusan_id",
	}
	preloads := []string{"PesertaDidik"}
	var conditions = map[string]interface{}{
		"semester_id": param.semesterId,
	}
	groupByColumns := []string{"tabel_anggotakelas.anggota_rombel_id"} // Hindari duplikasi
	anggotaKelasModel, err := repoAnggotaKelas.FindWithPreloadAndJoins(ctx, param.schemaname, joins, preloads, conditions, groupByColumns, nil, false)
	if err != nil {
		return err
	}

	for row, pd := range anggotaKelasModel {
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row+2), pd.PesertaDidik.PesertaDidikId)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row+2), pd.PesertaDidik.NmSiswa)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row+2), pd.SemesterId)
	}
	f.SetColVisible(sheetName, "A", false)
	return nil
}
func templateIjazah(ctx context.Context, db *gorm.DB, f *excelize.File, param ParamTemplate, sheetName string) error {
	// TODO: Implementasi template nilai akhir
	// sheetName := "Template"
	f.SetSheetName("Sheet1", sheetName)
	repoAnggotaKelas := repositories.NewRombelAnggotaRepository(db)
	joins := []string{
		"JOIN tabel_siswa ON tabel_siswa.peserta_didik_id = tabel_anggotakelas.peserta_didik_id",
		// "JOIN ref.jurusan ON tabel_kelas.jurusan_id = ref.jurusan.jurusan_id",
	}
	preloads := []string{"PesertaDidik"}
	var conditions = map[string]interface{}{
		"semester_id": param.semesterId,
	}
	groupByColumns := []string{"tabel_anggotakelas.anggota_rombel_id"} // Hindari duplikasi
	anggotaKelasModel, err := repoAnggotaKelas.FindWithPreloadAndJoins(ctx, param.schemaname, joins, preloads, conditions, groupByColumns, nil, false)
	if err != nil {
		return err
	}

	for row, pd := range anggotaKelasModel {
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row+2), pd.PesertaDidik.PesertaDidikId)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row+2), pd.PesertaDidik.NmSiswa)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row+2), pd.SemesterId)
	}
	f.SetColVisible(sheetName, "A", false)
	return nil
}

func getExcelColumnName(col int) string {
	result := ""
	for col >= 0 {
		result = string(rune('A'+(col%26))) + result
		col = col/26 - 1
	}
	return result
}

// ============================
// s header ke Excel
// for i, header := range headers {
// 	col := string(rune('A'+i)) + "1"
// 	f.SetCellValue(sheetName, col, header)
// }
// f.SetCellValue(sheetName, "Y1", "Timestamp")
// f.SetCellValue(sheetName, "Y2", time.Now().Format("2006-01-02 15:04:05")) // Isi dengan waktu saat ini

// f.SetCellValue(sheetName, "Z1", "UserID")
// f.SetCellValue(sheetName, "Z2", "123456") // Bisa diisi dengan ID pengguna yang mengunduh template

// f.SetColVisible(sheetName, "Y", false) // Sembunyikan kolom Timestamp
// f.SetColVisible(sheetName, "Z", false) // Sembunyikan kolom UserID

// sampleData := []interface{}{"12345", "987654321", "Budi Santoso", "Jakarta", "2005-08-10", "Laki-laki", "Islam"}
// for i, data := range sampleData {
// 	col := string(rune('A'+i)) + "2"
// 	f.SetCellValue(sheetName, col, data)
// }

// // Buat validasi dropdown untuk JenisKelamin
// dv := excelize.NewDataValidation(true)
// dv.Sqref = "F2:F1000" // Rentang sel yang divalidasi
// dv.SetDropList([]string{"L", "P"})

// if err := f.AddDataValidation(sheetName, dv); err != nil {
// 	return fmt.Errorf("gagal menambahkan validasi: %w", err)
// }

// // Validasi NilaiAkhir hanya angka 0-100
// dvNilai := excelize.NewDataValidation(true)
// dvNilai.Sqref = "D2:D100"
// // dvNilai.SetWholeNumber(0, 100)

// if err := f.AddDataValidation(sheetName, dvNilai); err != nil {
// 	return fmt.Errorf("gagal menambahkan validasi angka: %w", err)
// }

// err := f.ProtectSheet(sheetName, &excelize.SheetProtectionOptions{
// 	FormatCells:      false, // Tidak bisa ubah format
// 	FormatColumns:    false, // Tidak bisa ubah kolom
// 	FormatRows:       false, // Tidak bisa ubah baris
// 	InsertRows:       true,  // Bisa menambah baris baru
// 	InsertColumns:    false, // Tidak bisa menambah kolom baru
// 	InsertHyperlinks: false,
// 	DeleteRows:       false, // Tidak bisa hapus baris
// 	DeleteColumns:    false, // Tidak bisa hapus kolom
// 	// SelectLockedCells:   false,
// 	// SelectUnlockedCells: true,
// })
// if err != nil {
// 	return err
// }

// f.NewSheet("Panduan")
// f.SetCellValue("Panduan", "A1", "Panduan Pengisian Template")
// f.SetCellValue("Panduan", "A2", "1. Isi semua kolom sesuai dengan contoh yang diberikan.")
// f.SetCellValue("Panduan", "A3", "2. Gunakan dropdown untuk memilih data yang tersedia.")
// f.SetCellValue("Panduan", "A4", "3. Pastikan semua data terisi sebelum mengunggah ke sistem.")

// ============================

// func parseData[T any](rows [][]string, parser func([]string) *T) []*T {
// 	var dataList []*T

// 	// Pastikan ada lebih dari satu baris (karena baris pertama biasanya header)
// 	if len(rows) < 2 {
// 		return dataList
// 	}

// 	for _, row := range rows[1:] { // Lewati header
// 		if len(row) == 0 {
// 			continue
// 		}

// 		item := parser(row)
// 		if item != nil {
// 			dataList = append(dataList, item)
// 		}
// 	}

// 	return dataList
// }
