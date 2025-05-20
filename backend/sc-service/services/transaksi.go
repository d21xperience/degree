package services

import (
	"context"
	"fmt"
	"log"
	"sc-service/config"
	pb "sc-service/generated"
	"sc-service/models"
	"sc-service/repositories"
	"sc-service/utils"

	"github.com/google/uuid"
)

type TransaksiService struct {
	pb.UnimplementedTransaksiServiceServer
	repoIjazahBc   repositories.GenericRepository[models.IjazahBc]
	repoDegreeData repositories.GenericRepository[models.DegreeData]
	// repoTransaksiTenant repositories.GenericRepository[models.TransaksiTabelTenant]
}

func NewTransaksiService() *TransaksiService {
	repoIjazahBc := repositories.NewIjazahBcRepository(config.DB)
	repoDegreeData := repositories.NewDegreeDataRepository(config.DB)
	return &TransaksiService{
		repoIjazahBc:   *repoIjazahBc,
		repoDegreeData: *repoDegreeData,
	}
}

// =====================================
// IJAZAH BLOCKCHAIN
// =====================================
func (s *TransaksiService) CreateIjazahBlockchain(ctx context.Context, req *pb.CreateIjazahBlockchainRequest) (*pb.CreateIjazahBlockchainResponse, error) {
	log.Printf("Received Sekolah data request: %+v\n", req)
	// Daftar field yang wajib diisi
	requiredFields := []string{"DegreeData"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	// dataJSON, err := json.MarshalIndent(req.DegreeData, "", "  ")
	// if err != nil {
	// 	log.Printf("Gagal serialisasi DegreeData: %v", err)
	// } else {
	// 	log.Printf("Isi lengkap DegreeData:\n%s", dataJSON)
	// }
	degreeData := req.GetDegreeData()
	// log.Printf("%s", degreeData)
	// =================
	// proto sekolah
	sekolahClient, err := NewSekolahServiceClient()
	if err != nil {
		return nil, err
	}
	// =================
	sekolah := sekolahClient.SearchSekolah(req.Schemaname)

	var ijazahBcModels = &models.IjazahBc{
		ID:                          uuid.New(),
		PesertaDidikID:              utils.StringToUUID(degreeData.Ijazah.PesertaDidikId),
		Nama:                        degreeData.Ijazah.Nama,
		NIS:                         degreeData.Ijazah.Nis,
		NISN:                        degreeData.Ijazah.Nisn,
		TempatLahir:                 degreeData.Ijazah.TempatLahir,
		AsalSekolah:                 degreeData.Ijazah.AsalSekolah,
		NomorIjazah:                 degreeData.Ijazah.NomorIjazah,
		ProgramKeahlian:             degreeData.Ijazah.ProgramKeahlian,
		NamaOrtuwali:                degreeData.Ijazah.NamaOrtuwali,
		TempatIjazah:                degreeData.Ijazah.TempatIjazah,
		TanggalIjazah:               utils.TimeToPointer(degreeData.Ijazah.TanggalIjazah),
		PaketKeahlian:               degreeData.Ijazah.PaketKeahlian,
		KabupatenKota:               sekolah.KabKota,
		Provinsi:                    sekolah.Propinsi,
		SekolahPenyelenggaraUjianUS: sekolah.Nama,
		SekolahPenyelenggaraUjianUN: sekolah.Nama,
		NPSN:                        sekolah.Npsn,
	}

	// ijazahBcModels = append(ijazahBcModels, cek)
	degreeDataModels := &models.DegreeData{
		IjazahID:       ijazahBcModels.ID,
		DegreeHash:     degreeData.DegreeHash,
		TxHash:         degreeData.TxHash,
		IpfsURL:        degreeData.IpfsUrl,
		BcType:         degreeData.BcType,
		LinkBcExplorer: degreeData.LinkBcExplorer,
	}
	ijazahBc := s.repoIjazahBc.Save(ctx, ijazahBcModels, "public")
	if ijazahBc != nil {
		return nil, err
	}
	degreData := s.repoDegreeData.Save(ctx, degreeDataModels, "public")
	if degreData != nil {
		return nil, err
	}

	return &pb.CreateIjazahBlockchainResponse{
		Status:  true,
		Message: "Berhasil",
	}, nil
}

func (s *TransaksiService) GetIjazahBlockchain(ctx context.Context, req *pb.GetIjazahBlockchainRequest) (*pb.GetIjazahBlockchainResponse, error) {
	var err error
	// Daftar field yang wajib diisi
	requiredFields := []string{"SekolahId", "TahunAjaranId"}
	// Validasi request
	err = utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	sekolahId := req.GetSekolahId()
	tahunAjaranId := req.GetTahunAjaranId()
	if sekolahId == "\"\"" && tahunAjaranId == 0 {
		return nil, fmt.Errorf("sekolah_id is required")
	}
	conditions := map[string]any{
		"tahun_ajaran_id": tahunAjaranId,
		"sekolah_id":      sekolahId,
	}

	// joins := []string{
	// 	// "JOIN tabel_ptk ON tabel_kelas.ptk_id = tabel_ptk.ptk_id",
	// 	// "JOIN tabel_pembelajaran ON tabel_kelas.rombongan_belajar_id = tabel_pembelajaran.rombongan_belajar_id",
	// 	// fmt.Sprintf("JOIN ref.jurusan ON %s.tabel_kelas.jurusan_id = ref.jurusan.jurusan_id", schemaName),
	// 	// fmt.Sprintf("JOIN ref.kurikulum ON %s.tabel_kelas.kurikulum_id = ref.kurikulum.kurikulum_id", schemaName),
	// 	// fmt.Sprintf("JOIN ref.tingkat_pendidikan ON %s.tabel_kelas.tingkat_pendidikan_id = ref.tingkat_pendidikan.tingkat_pendidikan_id", schemaName),
	// }
	preloads := []string{"Ijazah"}
	// orderBy := []string{""}
	// groupByColumns := []string{} // Hindari duplikasi
	results, err := s.repoDegreeData.FindWithRelations(ctx, "public", nil, preloads, conditions, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	pbLulusan := utils.ConvertModelsToPB(utils.SliceToPointer(results), func(item *models.DegreeData) *pb.DegreeData {
		return &pb.DegreeData{
			DegreeHash:     item.DegreeHash,
			TxHash:         item.TxHash,
			IpfsUrl:        item.IpfsURL,
			BcType:         item.BcType,
			LinkBcExplorer: item.LinkBcExplorer,
			TahunAjaranId:  item.TahunAjaranId,
			Ijazah: &pb.Ijazah{
				PesertaDidikId: item.Ijazah.PesertaDidikID.String(),
				Nama:           item.Ijazah.Nama,
				Nisn:           item.Ijazah.NISN,
				NomorIjazah:    item.Ijazah.NomorIjazah,
				TempatLahir:    item.Ijazah.TempatIjazah,
				TanggalLahir:   item.Ijazah.TanggalLahir.Format("2006-01-26"),
				AsalSekolah:    item.Ijazah.AsalSekolah,
				NamaOrtuwali:   item.Ijazah.NamaOrtuwali,
			},
		}
	})
	return &pb.GetIjazahBlockchainResponse{
		Status:     true,
		Message:    "Sukses",
		DegreeData: pbLulusan,
		// AnggotaKelas: nil,
	}, nil
}

func (s *TransaksiService) SearchIjazahBlockchain(ctx context.Context, req *pb.SearchIjazahBlockchainRequest) (*pb.SearchIjazahBlockchainResponse, error) {
	// Debugging: Cek nilai request yang diterima
	log.Printf("Received Sekolah data request: %+v\n", req)
	// Daftar field yang wajib diisi
	// requiredFields := []string{"nisn"}
	// // Validasi request
	// err := utils.ValidateFields(req, requiredFields)
	// if err != nil {
	// 	return nil, err
	// }

	// schemaName := req.GetSchemaname()
	ijazahModel, err := s.repoIjazahBc.FindByID(ctx, req.GetNisn(), "public", "nisn")
	if err != nil {
		return nil, err
	}
	degreeModel, err := s.repoDegreeData.FindByID(ctx, ijazahModel.ID.String(), "public", "ijazah_id")
	if err != nil {
		return nil, err
	}
	respon := utils.ConvertModelToPB(degreeModel, func(item *models.DegreeData) *pb.DegreeData {
		return &pb.DegreeData{
			DegreeHash:     item.DegreeHash,
			TxHash:         item.TxHash,
			IpfsUrl:        item.IpfsURL,
			BcType:         item.BcType,
			LinkBcExplorer: item.LinkBcExplorer,
			TahunAjaranId:  item.TahunAjaranId,
			Ijazah: &pb.Ijazah{
				Nama:            ijazahModel.Nama,
				NomorIjazah:     ijazahModel.NomorIjazah,
				Nis:             ijazahModel.NIS,
				Nisn:            ijazahModel.NISN,
				ProgramKeahlian: ijazahModel.ProgramKeahlian,
				Kabupatenkota:   ijazahModel.KabupatenKota,
				Provinsi:        ijazahModel.Provinsi,
				TempatIjazah:    ijazahModel.TempatIjazah,
				AsalSekolah:     ijazahModel.AsalSekolah,
			},
		}
	})
	return &pb.SearchIjazahBlockchainResponse{
		Status:   true,
		Message:  "Berahail",
		IjazahBc: respon,
	}, nil
}
