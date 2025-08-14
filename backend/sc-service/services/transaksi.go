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
	repoIjazahBc      repositories.GenericRepository[models.IjazahBc]
	repoDegreeData    repositories.GenericRepository[models.DegreeData]
	repoContract      repositories.GenericRepository[models.Contract]
	repoBCTransaction repositories.GenericRepository[models.BCTransaction]
	// repoTransaksiTenant repositories.GenericRepository[models.TransaksiTabelTenant]
}

func NewTransaksiService() *TransaksiService {
	repoIjazahBc := repositories.NewIjazahBcRepository(config.DB)
	repoDegreeData := repositories.NewDegreeDataRepository(config.DB)
	repoContract := repositories.NewContractDataRepository(config.DB)
	repoBCTransaction := repositories.NewContractBCTransactionRepository(config.DB)
	return &TransaksiService{
		repoIjazahBc:      *repoIjazahBc,
		repoDegreeData:    *repoDegreeData,
		repoContract:      *repoContract,
		repoBCTransaction: *repoBCTransaction,
	}
}

// =====================================
// IJAZAH BLOCKCHAIN
// =====================================
func (s *TransaksiService) CreateIjazahBlockchain(ctx context.Context, req *pb.CreateIjazahBlockchainRequest) (*pb.CreateIjazahBlockchainResponse, error) {
	log.Printf("Received Sekolah data request: %+v\n", req)
	// Daftar field yang wajib diisi
	requiredFields := []string{"Schemaname", "DegreeData"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	degreeData := req.GetDegreeData()

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
		KabupatenKota:               degreeData.Ijazah.Kabupatenkota,
		Provinsi:                    degreeData.Ijazah.Provinsi,
		SekolahPenyelenggaraUjianUS: degreeData.Ijazah.SekolahPenyelenggaraUjianUs,
		SekolahPenyelenggaraUjianUN: degreeData.Ijazah.SekolahPenyelenggaraUjianUn,
		NPSN:                        degreeData.Ijazah.Npsn,
		TanggalLahir:                utils.TimeToPointer(degreeData.Ijazah.TanggalLahir),
	}
	degreeDataModels := &models.DegreeData{
		IjazahID:       ijazahBcModels.ID,
		DegreeHash:     degreeData.DegreeHash,
		TxHash:         degreeData.TxHash,
		IpfsURL:        degreeData.IpfsUrl,
		BcType:         degreeData.BcType,
		LinkBcExplorer: degreeData.LinkBcExplorer,
		TahunAjaranId:  degreeData.TahunAjaranId,
		SekolahId:      utils.UUIDToPointer(utils.StringToUUID(degreeData.SekolahId)),
	}
	bcTransactionModels := &models.BCTransaction{
		FromAddress: degreeData.BcTransaction.FromAddress,
		ToAddress:   degreeData.BcTransaction.ToAddress,
		Value:       degreeData.BcTransaction.Value,
		GasLimit:    degreeData.BcTransaction.GasLimit,
		GasPrice:    degreeData.BcTransaction.GasPrice,
		Nonce:       degreeData.BcTransaction.Nonce,
		Data:        degreeData.BcTransaction.Data,
		ChainId:     degreeData.BcTransaction.ChainId,
		BlockNumber: degreeData.BcTransaction.BlockNumber,
		Status:      degreeData.BcTransaction.Status,
	}

	ijazahBc := s.repoIjazahBc.Save(ctx, ijazahBcModels, "public")
	if ijazahBc != nil {
		return &pb.CreateIjazahBlockchainResponse{
			Status:  false,
			Message: "Gagal menyimpan Ijazah",
		}, nil
	}
	degreData := s.repoDegreeData.Save(ctx, degreeDataModels, "public")
	if degreData != nil {
		return &pb.CreateIjazahBlockchainResponse{
			Status:  false,
			Message: "Gagal menyimpan Hash",
		}, nil
	}
	bcTransaction := s.repoBCTransaction.Save(ctx, bcTransactionModels, req.GetSchemaname())
	if bcTransaction != nil {
		return &pb.CreateIjazahBlockchainResponse{
			Status:  false,
			Message: "Gagal menyimpan Transaksi blockchain!",
		}, nil
	}

	return &pb.CreateIjazahBlockchainResponse{
		Status:  true,
		Message: "Berhasil menyimpan seluluruh data",
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
		"tahun_ajaran_id": fmt.Sprintf("%d", tahunAjaranId),
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
			TglTransaksi:   item.CreatedAt.Format("2006-01-26"),
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

// func (s *TransaksiService) SaveContractAddress(ctx context.Context, req *pb.SaveContractAddressRequest) (*pb.SaveContractAddressResponse, error) {

// 	pbContract := models.ContractData{
// 		ContractAddres: &req.Contract.ContractAddress,
// 		ContractOwner:  &req.Contract.Owner,
// 	}
// 	err := s.repoContract.Save(ctx, &pbContract, "public")
// 	if err != nil {
// 		return &pb.SaveContractAddressResponse{
// 			Status:  false,
// 			Message: "Gagal menyimpan kontrak",
// 		}, nil
// 	}

// 	return &pb.SaveContractAddressResponse{
// 		Status:  true,
// 		Message: "Berhasil menyimpan kontrak",
// 	}, nil
// }
// func (s *TransaksiService) GetContractAddress(ctx context.Context, req *pb.GetContractAddressRequest) (*pb.GetContractAddressResponse, error) {
// 	modelContract, err := s.repoContract.FindAll(ctx, "public", 100, 0)
// 	if err != nil {
// 		return &pb.GetContractAddressResponse{
// 			Status:   true,
// 			Message:  "Gagal mengambil data",
// 			Contract: nil,
// 		}, nil
// 	}
// 	if err == nil {
// 		return nil, err
// 	}
// 	pbModelContract := &pb.Contract{
// 		ContractAddress: utils.SafeString(modelContract[0].ContractAddres),
// 		Owner:           utils.SafeString(modelContract[0].ContractOwner),
// 	}

// 	return &pb.GetContractAddressResponse{
// 		Status:   true,
// 		Message:  "Berhasil mengambil data",
// 		Contract: pbModelContract,
// 	}, nil
// }

func (s *TransaksiService) GetBCTransaction(ctx context.Context, req *pb.GetBCTransactionRequest) (*pb.GetBCTransactionResponse, error) {
	log.Printf("transaksi/GetBCTransactionRequest received data from request: %+v\n", req)
	// Daftar field yang wajib diisi
	requiredFields := []string{"Schemaname"}
	// Validasi request
	requiredFieldsResponse := utils.ValidateFields(req, requiredFields)
	if requiredFieldsResponse != nil {
		return nil, requiredFieldsResponse
	}
	schemaname := req.Schemaname
	// conditions := map[string]any{"conditions"}
	bcTransactionsModel, err := s.repoBCTransaction.FindAllByConditions(ctx, schemaname, nil, 100, 0)
	if err != nil {
		return &pb.GetBCTransactionResponse{
			Status:        true,
			Message:       "Gagal mendapatkan transaksi",
			BcTransaction: nil,
		}, nil
	}
	pbBCTransactions := utils.ConvertModelsToPB(bcTransactionsModel, func(item *models.BCTransaction) *pb.BCTransactions {
		return &pb.BCTransactions{
			FromAddress: item.FromAddress,
			ToAddress:   item.ToAddress,
			Value:       item.Value,
			GasLimit:    item.GasLimit,
			GasPrice:    item.GasPrice,
			Nonce:       item.Nonce,
			Data:        item.Data,
			ChainId:     item.ChainId,
			BlockNumber: item.BlockNumber,
			Status:      item.Status,
		}
	})
	return &pb.GetBCTransactionResponse{
		Status:        true,
		Message:       "Berhasil mendapatkan transaksi",
		BcTransaction: pbBCTransactions,
	}, nil

}

func (s *TransaksiService) DeployContract(ctx context.Context, req *pb.DeployContractRequest) (*pb.DeployContractResponse, error) {
	log.Printf("transaksi/DeployContract received data from request: %+v\n", req)
	// Daftar field yang wajib diisi
	requiredFields := []string{"ContractAddress"}
	// Validasi request
	requiredFieldsResponse := utils.ValidateFields(req, requiredFields)
	if requiredFieldsResponse != nil {
		return nil, requiredFieldsResponse
	}

	return &pb.DeployContractResponse{
		Status:  true,
		Message: "Behasil membuat smartcontract",
	}, nil
}

func (s *TransaksiService) GetSolcVersion(ctx context.Context, req *pb.Empty) (*pb.GetSolcVersionResponse, error) {
	cek, err := utils.GetSolcVersion()
	if err != nil {
		return &pb.GetSolcVersionResponse{
			Status:  false,
			Message: fmt.Sprintf("error get version:%v", err),
		}, nil
	}
	return &pb.GetSolcVersionResponse{
		Status:  true,
		Message: cek,
	}, nil
}
