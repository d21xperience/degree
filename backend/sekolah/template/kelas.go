package template

func GetKelasColumns() []TemplateColumn {
	return []TemplateColumn{
		{Name: "Nama", Example: "XI TKJ A"},
		{Name: "Tingkat Pendidikan", Example: "10", ColumnWidth: 16},
		{Name: "Jurusan Id", Example: JurusanList, ColumnWidth: 30},
	}
}
