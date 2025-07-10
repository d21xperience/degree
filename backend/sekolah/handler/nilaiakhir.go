package handler

import (
	"context"
	"sekolah/config"
	"sekolah/models"
	"sekolah/repositories"
	"sekolah/utils"

	"github.com/google/uuid"
)

type NilaiSiswa struct {
	PesertaDidikId      uuid.UUID
	NmSiswa             string
	RombonganBelajarId  uuid.UUID
	NmKelas             string
	SemesterId          string
	TingkatPendidikanId int32
	NilaiMapel          *[]NilaiMapel
}
type NilaiMapel struct {
	MataPelajaranId uint32
	NmMapel         string
	NilaiPeng       uint32
	SemesterId      string
	NmSingkatan     string
}
type NilaiSiswaHandler struct {
	repoNilaiAkhir    repositories.GenericRepository[models.NilaiAkhir]
	repoRombelAnggota repositories.GenericRepository[models.RombelAnggota]
}

func NewNilaiSiswaHandler() *NilaiSiswaHandler {
	repoNilaiAkhir := repositories.NewNilaiAkhirRepository(config.DB)
	repoRombelAnggota := repositories.NewRombelAnggotaRepository(config.DB)
	return &NilaiSiswaHandler{
		repoNilaiAkhir:    *repoNilaiAkhir,
		repoRombelAnggota: *repoRombelAnggota,
	}
}

func (s *NilaiSiswaHandler) GetNilaiSiswa(ctx context.Context, schemaName, semesterId, rombelId string) ([]NilaiSiswa, error) {
	conditions := map[string]any{
		"tabel_anggotakelas.semester_id":          semesterId,
		"tabel_anggotakelas.rombongan_belajar_id": rombelId,
	}
	joins := []string{
		"JOIN tabel_kelas ON tabel_kelas.rombongan_belajar_id = tabel_anggotakelas.rombongan_belajar_id",
		"JOIN tabel_siswa ON tabel_siswa.peserta_didik_id = tabel_anggotakelas.peserta_didik_id",
		// "JOIN tabel_ptk ON tabel_ptk.ptk_id = tabel_kelas.ptk_id",
	}
	selectFields := []string{
		"tabel_siswa.peserta_didik_id",
		"tabel_siswa.nm_siswa",
		"tabel_kelas.rombongan_belajar_id",
		"tabel_kelas.nm_kelas",
		"tabel_anggotakelas.semester_id",
		"tabel_kelas.tingkat_pendidikan_id",
	}
	type Hasil struct {
		PesertaDidikId      uuid.UUID
		NmSiswa             string
		RombonganBelajarId  uuid.UUID
		NmKelas             string
		SemesterId          string
		TingkatPendidikanId int32
	}
	var result []Hasil
	err := s.repoRombelAnggota.FindWithJoinAndScan(ctx, schemaName, joins, conditions, nil, selectFields, &result)
	if err != nil {
		return nil, err
	}
	var nilaiSiswa = NilaiSiswa{}
	var nilaiSeluruhSiswa = []NilaiSiswa{}
	for _, v := range result {
		nilaiAkhir, err := s.SearchNilaiSiswa(ctx, schemaName, semesterId, v.PesertaDidikId.String())
		if err != nil {
			return nil, err
		}
		nilaiSiswa = NilaiSiswa{
			NmSiswa:             v.NmSiswa,
			NmKelas:             v.NmKelas,
			SemesterId:          v.SemesterId,
			TingkatPendidikanId: v.TingkatPendidikanId,
			NilaiMapel:          &nilaiAkhir,
		}
		nilaiSeluruhSiswa = append(nilaiSeluruhSiswa, nilaiSiswa)
	}
	return nilaiSeluruhSiswa, nil
}

func (s *NilaiSiswaHandler) SearchNilaiSiswa(ctx context.Context, schemaName, semesterId, pesertaDidikId string) ([]NilaiMapel, error) {
	conditions := map[string]any{
		"tabel_nilaiakhir.semester_id":      semesterId,
		"tabel_nilaiakhir.peserta_didik_id": pesertaDidikId,
	}
	// joins := []string{
	// 	"JOIN tabel_mapel tm ON tm.mata_pelajaran_id = tn.mata_pelajaran_id",
	// }
	preloads := []string{"MataPelajaran"}
	// orderBy := []string{"tabel_kelas.nm_kelas ASC", "tabel_siswa.nm_siswa ASC"}

	nilaiAkhirModel, err := s.repoNilaiAkhir.FindWithPreloadAndJoinsOrigin(ctx, schemaName, nil, preloads, conditions, nil)
	if err != nil {
		return nil, err
	}
	nilaiSeluruhMapel := []NilaiMapel{}
	for _, v := range nilaiAkhirModel {
		nilaiMapel := NilaiMapel{
			MataPelajaranId: utils.PointerToUint32(v.MataPelajaranId),
			NmMapel:         v.MataPelajaran.Nama,
			NilaiPeng:       utils.SafeUint32(v.NilaiPeng),
			SemesterId:      v.SemesterId,
			NmSingkatan:     utils.SafeString(v.MataPelajaran.NmSingkatan),
		}
		nilaiSeluruhMapel = append(nilaiSeluruhMapel, nilaiMapel)
	}

	return nilaiSeluruhMapel, nil
}
