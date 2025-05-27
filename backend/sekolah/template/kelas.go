package template

import (
	"sekolah/utils"

	"github.com/xuri/excelize/v2"
)

func GetKelasColumns() []TemplateColumn {
	// formula1 := fmt.Sprintf(`"%s"`, strings.Join(kurikulums, ","))
	return []TemplateColumn{
		{Name: "Nama", Example: "XI TKJ A"},
		{Name: "Tingkat Pendidikan", Example: "10", ColumnWidth: 16},
		{Name: "Kurikulum", Example: "SMK Merdeka Teknik Komputer dan Jaringan (k)", ColumnWidth: 30, Validation: &excelize.DataValidation{
			Type:             "list",
			Formula1:         `"KurikulumList"`,
			ShowDropDown:     true,
			ShowErrorMessage: true,
			ErrorTitle:       utils.StringToPointer("Input Kurikulum Salah!"),
			Error:            utils.StringToPointer("Isi sesuai dengan kurikulum yang digunakan"),
			Sqref:            "C2:C1048576", // Kolom JK,
			PromptTitle:      utils.StringToPointer("Input kurikulum"),
			ShowInputMessage: true,
			Prompt:           utils.StringToPointer("Isi sesuai dengan kurikulum yang digunakan"),
		}},
		// {Name: "Jurusan", Example: "Teknik Komputer dan Jaringan", ColumnWidth: 30, Validation: &excelize.DataValidation{
		// 	Type:             "list",
		// 	Formula1:         JurusanList,
		// 	ShowDropDown:     true,
		// 	ShowErrorMessage: true,
		// 	ErrorTitle:       utils.StringToPointer("Input jurusan Salah!"),
		// 	Error:            utils.StringToPointer("Isi sesuai dengan kurikulum yang digunakan"),
		// 	Sqref:            "D2:D1048576", // Kolom JK,
		// 	PromptTitle:      utils.StringToPointer("Input jurusan"),
		// 	ShowInputMessage: true,
		// 	Prompt:           utils.StringToPointer("Isi sesuai dengan jurusan yang digunakan"),
		// }},
	}
}
