package utils

import (
	"fmt"
	"regexp"
	"sekolah/models"
	"strings"
	"unicode"
)

func SearchString(input string) string {
	// input := "invalid request: missing fields - TahunAjaran"

	// Gunakan regex untuk mencari kata setelah "- "
	re := regexp.MustCompile(`- (\w+)`)
	match := re.FindStringSubmatch(input)

	if len(match) > 1 {
		field := match[1]
		fmt.Println("Field:", field) // Output: TahunAjaran
		return field
	} else {
		fmt.Println("Tidak ada field ditemukan")
		return ""
	}
}

// versi 3
func ClassifyEducationForm(schoolName string, educationForms []models.BentukPendidikan) *models.BentukPendidikan {
	schoolName = strings.ToLower(schoolName)

	var bestMatch *models.BentukPendidikan
	bestIndex := len(schoolName) + 1 // index besar, pasti dikalahkan

	for _, form := range educationForms {
		formName := strings.ToLower(form.Nama)
		idx := strings.Index(schoolName, formName)

		if idx != -1 {
			// Prioritaskan yang:
			// 1. Muncul lebih awal (idx kecil)
			// 2. Jika idx sama, pilih nama form yang lebih panjang (lebih spesifik)
			if bestMatch == nil || idx < bestIndex || (idx == bestIndex && len(formName) > len(bestMatch.Nama)) {
				bestMatch = &form
				bestIndex = idx
			}
		}
	}

	return bestMatch
}

func ClassifyEducationGrade(educationForm *models.BentukPendidikan) uint32 {
	if educationForm == nil {
		return 1
	}
	if educationForm.JenjangPaud == 1 {
		return 1
	} else if educationForm.JenjangTk == 1 {
		return 2
	} else if educationForm.JenjangSd == 1 {
		return 4
	} else if educationForm.JenjangSmp == 1 {
		return 5
	} else if educationForm.JenjangSma == 1 {
		return 6
	} else {
		return 1
	}
}

// Daftar bentuk pendidikan yang dikenali
// var educationForms = []string{
// 	"SMAK", "MAK", "Paket C", "SMA", "Uttama Dhammasekha", "Uttama Dhammasekha Kejuruan",
// 	"Satap SD SMP SMA", "SPK SMA", "SMLB", "SMK", "MA", "Utama W P", "SMTK", "SLB", "SMAg.K",
// 	"Satap TK SD SMP", "RA", "SMP Terbuka", "SMPTK", "SDTK", "SPK PG", "SPK TK", "SPK SD",
// 	"SPK SMP", "Pratama W P", "Adi W P", "Madyama W P", "SKB", "Taman Seminari", "TKLB",
// 	"Pondok Pesantren", "SPK KB", "Mula Dhammasekha", "Muda Dhammasekha", "TK", "KB", "TPA",
// 	"SPS", "SD", "SMP", "SDLB", "SMPLB", "MTs", "MI", "Paket A", "Paket B", "Akademik",
// 	"Politeknik", "Sekolah Tinggi", "Institut", "Universitas", "Kursus", "Keaksaraan", "TBM",
// 	"PKBM", "Life Skill", "Satap TK SD", "Satap SD SMP",
// }

// // Fungsi utama untuk klasifikasi bentuk pendidikan
// func FindBentukPendidikan(schoolName string) string {
// 	schoolName = strings.ToUpper(schoolName)

// 	for _, form := range educationForms {
// 		if strings.Contains(schoolName, form) {
// 			return form
// 		}
// 	}

//		return "Tidak Diketahui"
//	}
// func BuatSingkatan(nama string) string {
// 	var kataYangDiabaikan = map[string]bool{
// 		"dan":    true,
// 		"dari":   true,
// 		"ke":     true,
// 		"yang":   true,
// 		"untuk":  true,
// 		"dengan": true,
// 		"di":     true,
// 	}

// 	// Singkatan khusus (override)
// 	var singkatanKhusus = map[string]string{
// 		"akuntansi": "AK",
// 	}

// 	// Cek override khusus
// 	normalisasi := strings.ToLower(strings.TrimSpace(nama))
// 	if val, ok := singkatanKhusus[normalisasi]; ok {
// 		return val
// 	}

// 	kata := strings.Fields(normalisasi)
// 	singkatan := ""

// 	for _, k := range kata {
// 		if kataYangDiabaikan[k] {
// 			continue
// 		}
// 		for _, r := range k {
// 			if unicode.IsLetter(r) {
// 				singkatan += strings.ToUpper(string(r))
// 				break
// 			}
// 		}
// 	}

// 	return singkatan
// }
// VERSI 2
func BuatSingkatan(nama string) string {
	var kataYangDiabaikan = map[string]bool{
		"dan":    true,
		"dari":   true,
		"ke":     true,
		"yang":   true,
		"untuk":  true,
		"dengan": true,
		"di":     true,
	}

	var singkatanKhusus = map[string]string{
		"akuntansi": "AK",
	}

	normalisasi := strings.ToLower(strings.TrimSpace(nama))
	if val, ok := singkatanKhusus[normalisasi]; ok {
		return val
	}

	kata := strings.Fields(normalisasi)
	singkatan := ""

	for _, k := range kata {
		if kataYangDiabaikan[k] {
			continue
		}
		for _, r := range k {
			if unicode.IsLetter(r) {
				singkatan += strings.ToUpper(string(r))
				break
			}
		}
	}

	return strings.TrimSpace(singkatan)
}
