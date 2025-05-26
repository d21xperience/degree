package template

import (
	"sekolah/utils"

	"github.com/xuri/excelize/v2"
)

func GetSiswaColumns() []TemplateColumn {
	return []TemplateColumn{
		{Name: "Tingkat Pendidikan", Example: "10", ColumnWidth: 11},
		{Name: "Nama Kelas", Example: "X TKJ A", ColumnWidth: 11},
		{Name: "NIS", Example: "123456"},
		{Name: "NISN", Example: "99887766"},
		{Name: "Nama", Example: "Ahmad"},
		{Name: "JK", Example: "L", ColumnWidth: 3, Validation: &excelize.DataValidation{
			Type:             "list",
			Formula1:         `"L,P"`,
			ShowDropDown:     true,
			ShowErrorMessage: true,
			ErrorTitle:       utils.StringToPointer("Input JK Salah!"),
			Error:            utils.StringToPointer("Isi hanya L atau P"),
			Sqref:            "F2:F1048576", // Kolom JK,
			PromptTitle:      utils.StringToPointer("Input Jk"),
			ShowInputMessage: true,
			Prompt:           utils.StringToPointer("Isi L atau P"),
		}},
		{Name: "Tempat Lahir", Example: "Jakarta"},
		{Name: "Tanggal Lahir", Example: "09/12/2016", FormatStyle: &excelize.Style{
			// NumFmt: 14,

		}},
		{Name: "Agama", Example: "Islam", Validation: &excelize.DataValidation{
			Type:             "list",
			Formula1:         `"Islam,Kristen,Katolik,Hindu,Buddha,Khonghucu"`,
			ShowDropDown:     true,
			ShowErrorMessage: true,
			ErrorTitle:       utils.StringToPointer("Input Agama Salah!"),
			Error:            utils.StringToPointer("Pilih agama dari daftar"),
			Sqref:            "G2:G1048576",
			PromptTitle:      utils.StringToPointer("Input Agama"),
			ShowInputMessage: true,
			Prompt:           utils.StringToPointer("Tulis agama: Islam,Kristen,Katolik,Hindu,Buddha,Khonghucu"),
		}},
		{Name: "Alamat", Example: "Jl. Merdeka 10"},
		{Name: "Telepon Siswa", Example: "08123456789"},
		{Name: "Diterima Tanggal", Example: "01/01/2022", FormatStyle: &excelize.Style{NumFmt: 14}},
		{Name: "Nama Ayah", Example: "Bapak Ahmad"},
		{Name: "Nama Ibu", Example: "Ibu Siti"},
		{Name: "Pekerjaan Ayah", Example: "Petani", ColumnWidth: 11},
		{Name: "Pekerjaan Ibu", Example: "Ibu Rumah Tangga", ColumnWidth: 11},
		{Name: "Nama Wali", Example: "Paman Budi", ColumnWidth: 13},
		{Name: "Pekerjaan Wali", Example: "Guru"},
		{Name: "NIK Siswa", Example: "3276010101010001"},
		{Name: "Status Dalam Kel", Example: "Anak Kandung", ColumnWidth: 14},
		{Name: "Anak ke", Example: "2"},
		{Name: "Sekolah Asal", Example: "SDN 1 Jakarta"},
		{Name: "Diterima di Kelas", Example: "7"},
		{Name: "Alamat Ortu", Example: "Jl. Sudirman"},
		{Name: "Telepon Ortu", Example: "0811223344"},
		{Name: "Alamat Wali", Example: "Jl. Gajah Mada"},
		{Name: "Telepon Wali", Example: "0813556677"},
	}
}
