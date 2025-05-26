package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sekolah/config"
	pb "sekolah/generated"
	"sekolah/models"
	"sekolah/repositories"
	"sekolah/utils"
	"strconv"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SekolahService struct {
	pb.UnimplementedSekolahServiceServer
	// RedisClient    *redis.Client // Tambahkan Redis sebagai field
	sekolahService      repositories.SekolahRepository
	schemaService       SchemaService
	repoSekolahTenant   repositories.GenericRepository[models.SekolahTenant]
	repoKategoriSekolah repositories.GenericRepository[models.KategoriSekolah]
}

func NewSekolahService() *SekolahService {
	sekolahRepo := repositories.NewSekolahRepository(config.DB)
	schemaRepo := repositories.NewSchemaRepository(config.DB)
	SekolahTenant := repositories.NewsekolahTenantRepository(config.DB)
	schemaService := NewSchemaService(schemaRepo, *SekolahTenant)
	repoKategoriSekolah := repositories.NewKategoriSekolahRepository(config.DB)
	return &SekolahService{
		sekolahService:      sekolahRepo,
		schemaService:       schemaService,
		repoSekolahTenant:   *SekolahTenant,
		repoKategoriSekolah: *repoKategoriSekolah,
	}
}

func (s *SekolahService) RegistrasiSekolah(ctx context.Context, req *pb.TabelSekolahRequest) (*pb.TabelSekolahResponse, error) {
	// Debugging: Cek nilai request yang diterima
	log.Printf("Received Sekolah data request: %+v\n", req)
	requiredFields := []string{"Sekolah"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}

	sekolah := req.GetSekolah()
	Schemaname := fmt.Sprintf("tabel_%s", sekolah.EnkripId)
	existingSchema, err := s.schemaService.GetSchemaBySchemaname(Schemaname)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("Lanjutkan pendaftaran: %v", err)
	}

	if existingSchema != nil {
		return &pb.TabelSekolahResponse{
			Message: "Pendaftaran dibatalkan: Sekolah sudah terdaftar",
			Status:  false,
		}, nil
	}

	cek := s.schemaService.RegistrasiSekolah(ctx, Schemaname)
	if cek != nil {
		// Tangani error spesifik jika error adalah gorm.ErrInvalidData
		if errors.Is(cek, gorm.ErrInvalidData) {
			return nil, errors.New("registrasi sekolah gagal: data tidak valid")
		}
		// Tangani error lainnya
		return nil, fmt.Errorf("registrasi sekolah gagal: %w", cek)
	}
	// 2 Simpan informasi schema sekolah
	err = s.schemaService.SimpanSchemaSekolah(&models.SekolahTenant{
		SekolahTenantId: sekolah.SekolahTenantId,
		SchemaName:      Schemaname,
		NamaSekolah:     sekolah.NamaSekolah,
	})
	if err != nil {
		log.Printf("Gagal menyimpan schema sekolah: %v", err)
		return nil, errors.New("gagal menyimpan informasi sekolah")
	}

	// 3 Kirim respon sukses
	return &pb.TabelSekolahResponse{
		Message: "Pembuatan database berhasil",
		Status:  true,
	}, nil
}

func (s *SekolahService) GetSekolahTenant(ctx context.Context, req *pb.SekolahTenantRequest) (*pb.SekolahTenantResponse, error) {
	// Daftar field yang wajib diisi
	requiredFields := []string{"SekolahTenantId"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	sekolahID := req.GetSekolahTenantId()
	sekolahTerdaftar, err := s.repoSekolahTenant.FindByID(ctx, strconv.Itoa(int(sekolahID)), "public", "sekolah_tenant_id")
	if err != nil {
		return nil, err
	}
	if sekolahTerdaftar == nil {
		return nil, err
	}
	return &pb.SekolahTenantResponse{
		SekolahTenantId: sekolahTerdaftar.SekolahTenantId,
		NamaSekolah:     sekolahTerdaftar.NamaSekolah,
		Schemaname:      sekolahTerdaftar.SchemaName, // nama schema
	}, err

}

// SCHEMA TABLE SEKOLAH---------------------input informasi sekolah yang telah terdaftar
// ================================================================================//
func (s *SekolahService) CreateSekolah(ctx context.Context, req *pb.CreateSekolahRequest) (*pb.CreateSekolahResponse, error) {
	// Debugging: Cek nilai request yang diterima
	// log.Printf("Received Sekolah data request: %+v\n", req)
	requiredFields := []string{"Schemaname", "Sekolah"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	Schemaname := req.GetSchemaname()
	sekolah := req.GetSekolah()
	// sekolahID, _ := uuid.Parse(sekolah.SekolahId)
	// // Contoh: Konversi UUID ke string
	// uuidObj := uuid.New() // Generate UUID baru
	// strValue, err := utils.ConvertUUIDToStringViceVersa(uuidObj)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Println("UUID ke String:", strValue)
	// }
	// bentukPendidikan, err := s.sekolahService.GetKategoriSekolah(ctx, sekolah.Nama)
	// if err != nil {
	// 	return nil, err
	// }

	// sekolah.StatusKepemilikanId = 4
	// sekolah.BentukPendidikanId = 39
	// sekolah.JenjangPendidikanId = 98
	bentukPendidikan := 39
	jenjangPendidikan := 6
	statusKepemilikan := 4
	sekolahModel := &models.Sekolah{
		SekolahID: uuid.New(),
		Nama:      sekolah.Nama,
		NPSN:      sekolah.Npsn,
		Alamat:    &sekolah.Alamat,
		KodePos:   &sekolah.KdPos,
		Telepon:   &sekolah.Telepon,
		Fax:       &sekolah.Fax,
		Kelurahan: &sekolah.Kelurahan,
		Kecamatan: &sekolah.Kecamatan,
		KabKota:   &sekolah.KabKota,
		Propinsi:  &sekolah.Propinsi,
		Website:   &sekolah.Website,

		Email:        &sekolah.Email,
		NamaKepsek:   &sekolah.NmKepsek,
		NIPKepsek:    &sekolah.NipKepsek,
		NIYKepsek:    &sekolah.NipKepsek,
		KodeAktivasi: &sekolah.KodeAktivasi,

		StatusKepemilikanID: utils.Uint32ToPointer(uint32(statusKepemilikan)),
		JenjangPendidikanID: utils.Uint32ToPointer(uint32(jenjangPendidikan)),
		BentukPendidikanID:  utils.Uint32ToPointer(uint32(bentukPendidikan)),
	}

	sekolahTerdaftar := s.sekolahService.Save(ctx, sekolahModel, Schemaname)
	if sekolahTerdaftar != nil {
		log.Printf("Gagal menyimpan sekolah: %v", sekolahTerdaftar.Error())
		return nil, errors.New("gagal menyimpan informasi sekolah")
	}

	return &pb.CreateSekolahResponse{
		Message: "sekolah berhasil ditambahkan",
		Status:  true,
	}, nil

}

func (s *SekolahService) GetSekolah(ctx context.Context, req *pb.GetSekolahRequest) (*pb.GetSekolahResponse, error) {
	//  Ambil schema dari request
	Schemaname := req.GetSchemaname()
	// sekolahID := req.GetSekolahId()

	//  Cari sekolah berdasarkan ID dan schema
	sekolah, err := s.sekolahService.Find(ctx, Schemaname)
	if err != nil {
		log.Printf("Gagal menemukan sekolah: %v", err)
		return nil, fmt.Errorf("gagal menemukan sekolah: %w", err)
	}

	//  Return response dalam format protobuf
	return &pb.GetSekolahResponse{
		Sekolah: &pb.SekolahDapo{
			SekolahId:           sekolah.SekolahID.String(),
			Nama:                sekolah.Nama,
			Nss:                 sekolah.NSS,
			Npsn:                sekolah.NPSN,
			Alamat:              utils.SafeString(sekolah.Alamat),
			KdPos:               utils.SafeString(sekolah.KodePos),
			Telepon:             utils.SafeString(sekolah.Telepon),
			Fax:                 utils.SafeString(sekolah.Fax),
			Kelurahan:           utils.SafeString(sekolah.Kelurahan),
			Kecamatan:           utils.SafeString(sekolah.Kecamatan),
			KabKota:             utils.SafeString(sekolah.KabKota),
			Propinsi:            utils.SafeString(sekolah.Propinsi),
			Website:             utils.SafeString(sekolah.Website),
			Email:               utils.SafeString(sekolah.Email),
			NmKepsek:            utils.SafeString(sekolah.NamaKepsek),
			NipKepsek:           utils.SafeString(sekolah.NIPKepsek),
			NiyKepsek:           utils.SafeString(sekolah.NIPKepsek),
			StatusKepemilikanId: utils.SafeUint32(sekolah.StatusKepemilikanID),
			KodeAktivasi:        utils.SafeString(sekolah.KodeAktivasi),
			JenjangPendidikanId: utils.SafeUint32(sekolah.JenjangPendidikanID),
			BentukPendidikanId:  utils.SafeUint32(sekolah.BentukPendidikanID),
		},
	}, nil
}

// Tambahkan fitur tambahan DELET, UPDATE , dan LIST digunakan untuk SUPER ADMIN
func (s *SekolahService) UpdateSekolah(ctx context.Context, req *pb.UpdateSekolahRequest) (*pb.UpdateSekolahResponse, error) {
	// Debugging: Cek nilai request yang diterima
	// log.Printf("Received Sekolah data request: %+v\n", req)
	requiredFields := []string{"Schemaname", "Sekolah"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	Schemaname := req.GetSchemaname()
	sekolah := req.GetSekolah()

	modSekolah, err := s.sekolahService.FindByID(ctx, sekolah.SekolahId, Schemaname)
	if err != nil {
		return nil, err
	}
	if modSekolah == nil {
		return nil, fmt.Errorf("sekolah dengan ID %s tidak ditemukan", sekolah.SekolahId)
	}
	// Update fields
	modSekolah.NSS = sekolah.Nss
	modSekolah.Website = &sekolah.Website
	modSekolah.Email = &sekolah.Email
	modSekolah.Fax = &sekolah.Fax
	modSekolah.Telepon = &sekolah.Telepon
	modSekolah.BentukPendidikanID = &sekolah.BentukPendidikanId
	modSekolah.JenjangPendidikanID = &sekolah.JenjangPendidikanId
	modSekolah.KodePos = &sekolah.KdPos
	modSekolah.NamaKepsek = &sekolah.NmKepsek
	modSekolah.NIPKepsek = &sekolah.NipKepsek

	err = s.sekolahService.Update(ctx, modSekolah, Schemaname)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateSekolahResponse{
		Message: "Sekolah berhasil diupdate",
		Status:  true,
	}, nil
}

func (s *SekolahService) CreateKategoriSekolah(ctx context.Context, req *pb.CreateKategoriSekolahRequest) (*pb.CreateKategoriSekolahResponse, error) {
	var err error
	log.Printf("Received Sekolah data request: %+v\n", req)
	requiredFields := []string{"Schemaname", "KategoriSekolah"}
	// Validasi request
	err = utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	Schemaname := req.GetSchemaname()
	pbKategoriSekolah := req.KategoriSekolah
	kategoriSekolahModel := &models.KategoriSekolah{
		NamaKurikulum: &pbKategoriSekolah.NamaKurikulum,
		NamaJurusan:   &pbKategoriSekolah.NamaJurusan,
		TahunAjaranId: pbKategoriSekolah.TahunAjaranId,
	}
	err = s.repoKategoriSekolah.Save(ctx, kategoriSekolahModel, Schemaname)
	if err != nil {
		return &pb.CreateKategoriSekolahResponse{
			Message: "Gagal menambahkan kategori sekolah",
			Status:  false,
		}, nil
	}
	return &pb.CreateKategoriSekolahResponse{
		Message: "Berhasil menambahkan kategori sekolah",
		Status:  true,
	}, nil
}

func (s *SekolahService) GetKategoriSekolah(ctx context.Context, req *pb.GetKategoriSekolahRequest) (*pb.GetKategoriSekolahResponse, error) {
	var err error
	log.Printf("Received Sekolah data request: %+v\n", req)
	requiredFields := []string{"Schemaname", "TahunAjaranId"}
	// Validasi request
	err = utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	Schemaname := req.GetSchemaname()
	tahunAjaranId := req.GetTahunAjaranId()
	exactConditions := map[string]interface{}{
		"tabel_kategori_sekolah.tahun_ajaran_id": tahunAjaranId,
	}
	modelKategoriSekolah, err := s.repoKategoriSekolah.FindWithRelations(ctx, Schemaname, nil, nil, exactConditions, nil, nil, nil)
	if err != nil {
		return &pb.GetKategoriSekolahResponse{
			Message:         "Gagal mendapatkan kategori sekolah",
			Status:          false,
			KategoriSekolah: nil,
		}, nil
	}

	pbKategoriSekolah := utils.ConvertModelsToPB(modelKategoriSekolah, func(item models.KategoriSekolah) *pb.KategoriSekolah {
		return &pb.KategoriSekolah{
			KategoriSekolahId: uint32(item.KategorisekolahId),
			NamaKurikulum:     utils.SafeString(item.NamaKurikulum),
			NamaJurusan:       utils.SafeString(item.NamaJurusan),
			TahunAjaranId:     item.TahunAjaranId,
		}
	})
	return &pb.GetKategoriSekolahResponse{
		Message:         "Berhasil mendapatkan kategori sekolah",
		Status:          true,
		KategoriSekolah: pbKategoriSekolah,
	}, nil
}
