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
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

type SekolahService struct {
	pb.UnimplementedSekolahServiceServer
	// RedisClient    *redis.Client // Tambahkan Redis sebagai field
	sekolahService             repositories.SekolahRepository
	schemaService              SchemaService
	repoSekolahTenant          repositories.GenericRepository[models.SekolahTenant]
	repoKategoriSekolah        repositories.GenericRepository[models.KategoriSekolah]
	repoBentukPendidikan       repositories.GenericRepository[models.BentukPendidikan]
	repoJenjangPendidikan      repositories.GenericRepository[models.JenjangPendidikan]
	repoKategoriMapel          repositories.GenericRepository[models.KategoriMapel]
	repoKelas                  repositories.GenericRepository[models.RombonganBelajar]
	repoPembelajaran           repositories.GenericRepository[models.Pembelajaran]
	repoMataPelajaran          repositories.GenericRepository[models.MataPelajaran]
	repoMataPelajaranKurikulum repositories.GenericRepository[models.MataPelajaranKurikulum]
}

func NewSekolahService() *SekolahService {
	sekolahRepo := repositories.NewSekolahRepository(config.DB)
	schemaRepo := repositories.NewSchemaRepository(config.DB)
	SekolahTenant := repositories.NewsekolahTenantRepository(config.DB)
	schemaService := NewSchemaService(schemaRepo, *SekolahTenant)
	repoKategoriSekolah := repositories.NewKategoriSekolahRepository(config.DB)
	repoBentukPendidikan := repositories.NewBentukPendidikanRepository(config.DB)
	repoJenjangPendidikan := repositories.NewJenjangPendidikanRepository(config.DB)
	repoKelas := repositories.NewrombonganBelajarRepository(config.DB)
	repoPembelajaran := repositories.NewPembelajaranRepository(config.DB)
	repoMapel := repositories.NewMapelRepository(config.DB)
	repoMapelKurikulum := repositories.NewMapelKurikulumRepository(config.DB)
	repoKategoriMapel := repositories.NewKategoriMapelRepository(config.DB)

	return &SekolahService{
		sekolahService:             sekolahRepo,
		schemaService:              schemaService,
		repoSekolahTenant:          *SekolahTenant,
		repoKategoriSekolah:        *repoKategoriSekolah,
		repoBentukPendidikan:       *repoBentukPendidikan,
		repoJenjangPendidikan:      *repoJenjangPendidikan,
		repoKelas:                  *repoKelas,
		repoPembelajaran:           *repoPembelajaran,
		repoMataPelajaran:          *repoMapel,
		repoMataPelajaranKurikulum: *repoMapelKurikulum,
		repoKategoriMapel:          *repoKategoriMapel,
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
	bentukPendidikanModel, err := s.repoBentukPendidikan.FindWithStringMatch(ctx, "ref", sekolah.Nama, "nama")
	if err != nil {
		return nil, err
	}

	var bentukPendidikan uint32
	cek := utils.ClassifyEducationForm(sekolah.Nama, *bentukPendidikanModel)
	if cek == nil {
		bentukPendidikan = 1
	} else {
		bentukPendidikan = cek.BentukPendidikanID
	}
	jenjangPendidikan := utils.ClassifyEducationGrade(cek)
	statusKepemilikan := 4
	lamaPendidikan := 3
	switch jenjangPendidikan {
	case 4:
		lamaPendidikan = 6
	case 2:
		lamaPendidikan = 1
	case 1:
		lamaPendidikan = 1
	default:
		lamaPendidikan = 3
	}
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
		LamaPendidikan:      utils.Uint32ToPointer(uint32(lamaPendidikan)),
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
	preloads := []string{"BentukPendidikan", "JenjangPendidikan"}
	//  Cari sekolah berdasarkan ID dan schema
	// sekolah, err := s.sekolahService.Find(ctx, Schemaname)
	sekolah, err := s.sekolahService.FindWithRelations(ctx, Schemaname, nil, preloads, nil, nil, nil, nil)
	if err != nil {
		log.Printf("Gagal menemukan sekolah: %v", err)
		return nil, fmt.Errorf("gagal menemukan sekolah: %w", err)
	}

	//  Return response dalam format protobuf
	return &pb.GetSekolahResponse{
		Sekolah: &pb.SekolahDapo{
			SekolahId:           sekolah[0].SekolahID.String(),
			Nama:                sekolah[0].Nama,
			Nss:                 sekolah[0].NSS,
			Npsn:                sekolah[0].NPSN,
			Alamat:              utils.SafeString(sekolah[0].Alamat),
			KdPos:               utils.SafeString(sekolah[0].KodePos),
			Telepon:             utils.SafeString(sekolah[0].Telepon),
			Fax:                 utils.SafeString(sekolah[0].Fax),
			Kelurahan:           utils.SafeString(sekolah[0].Kelurahan),
			Kecamatan:           utils.SafeString(sekolah[0].Kecamatan),
			KabKota:             utils.SafeString(sekolah[0].KabKota),
			Propinsi:            utils.SafeString(sekolah[0].Propinsi),
			Website:             utils.SafeString(sekolah[0].Website),
			Email:               utils.SafeString(sekolah[0].Email),
			NmKepsek:            utils.SafeString(sekolah[0].NamaKepsek),
			NipKepsek:           utils.SafeString(sekolah[0].NIPKepsek),
			NiyKepsek:           utils.SafeString(sekolah[0].NIPKepsek),
			StatusKepemilikanId: utils.SafeUint32(sekolah[0].StatusKepemilikanID),
			KodeAktivasi:        utils.SafeString(sekolah[0].KodeAktivasi),
			JenjangPendidikanId: utils.SafeUint32(sekolah[0].JenjangPendidikanID),
			BentukPendidikanId:  utils.SafeUint32(sekolah[0].BentukPendidikanID),
			LamaPendidikan:      utils.SafeUint32(sekolah[0].LamaPendidikan),
		},
		BentukPendidikanStr:  sekolah[0].BentukPendidikan.Nama,
		JenjangPendidikanStr: sekolah[0].JenjangPendidikan.Nama,
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
		NamaKurikulum:       &pbKategoriSekolah.NamaKurikulum,
		NamaJurusan:         &pbKategoriSekolah.NamaJurusan,
		TahunAjaranId:       pbKategoriSekolah.TahunAjaranId,
		JurusanId:           &pbKategoriSekolah.JurusanId,
		KurikulumId:         pbKategoriSekolah.KurikulumId,
		JenjangPendidikanId: &pbKategoriSekolah.JenjangPendidikanId,
		NamaBidangKeahlian:  &pbKategoriSekolah.NamaBidangKeahlian,
		NamaProgramKeahlian: &pbKategoriSekolah.NamaProgramKeahlian,
		TingkatId:           pbKategoriSekolah.TingkatId,
		Jumlah:              &pbKategoriSekolah.Jumlah,
	}
	err = s.repoKategoriSekolah.Save(ctx, kategoriSekolahModel, Schemaname)
	if err != nil {
		return &pb.CreateKategoriSekolahResponse{
			Message: "Gagal menambahkan kategori sekolah",
			Status:  false,
		}, nil
	}
	log.Printf("Sukses eksekusi sekolah_server/CreateKategoriSekolah!")
	return &pb.CreateKategoriSekolahResponse{
		Message: "Berhasil menambahkan kategori sekolah & kelas",
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
	exactConditions := map[string]any{
		"tabel_kategori_sekolah.tahun_ajaran_id": tahunAjaranId,
	}
	// groupByColumns := []string{"tabel_kategori_sekolah.kurikulum_id"}
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
			KategoriSekolahId:   item.KategorisekolahId,
			NamaKurikulum:       utils.SafeString(item.NamaKurikulum),
			NamaJurusan:         utils.SafeString(item.NamaJurusan),
			TahunAjaranId:       item.TahunAjaranId,
			KurikulumId:         item.KurikulumId,
			JurusanId:           utils.SafeString(item.JurusanId),
			NamaBidangKeahlian:  utils.SafeString(item.NamaBidangKeahlian),
			NamaProgramKeahlian: utils.SafeString(item.NamaProgramKeahlian),
			JenjangPendidikanId: utils.SafeInt32(item.JenjangPendidikanId),
			TingkatId:           item.TingkatId,
			Jumlah:              utils.SafeInt32(item.Jumlah),
		}
	})
	return &pb.GetKategoriSekolahResponse{
		Message:         "Berhasil mendapatkan kategori sekolah ",
		Status:          true,
		KategoriSekolah: pbKategoriSekolah,
	}, nil
}

func (s *SekolahService) DeleteKategoriSekolah(ctx context.Context, req *pb.DeleteKategoriSekolahRequest) (*pb.DeleteKategoriSekolahResponse, error) {
	var err error
	// Debugging: Cek nilai request yang diterima
	log.Printf("Received Sekolah data request: %+v\n", req)
	requiredFields := []string{"Schemaname"}
	// Validasi request
	err = utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	// if req.GetKategoriSekolahId() == 0 || req.GetKurikulumId() == "" {
	// 	return nil, err
	// }
	Schemaname := req.GetSchemaname()
	if req.KurikulumId != nil {
		err = s.repoKategoriSekolah.Delete(ctx, req.GetKurikulumId(), Schemaname, "kurikulum_id")
	} else {
		kategoriSekolahId := strconv.Itoa(int(req.GetKategoriSekolahId()))
		err = s.repoKategoriSekolah.Delete(ctx, kategoriSekolahId, Schemaname, "kategori_sekolah_id")
	}
	if err != nil {
		return &pb.DeleteKategoriSekolahResponse{
			Message: fmt.Sprintf("Gagal menyimpan data kategori sekolah: %s", err),
			Status:  false,
		}, nil
	}
	return &pb.DeleteKategoriSekolahResponse{
		Message: "Kategori sekolah berhasil dihapus!",
		Status:  true,
	}, nil

}

func (s *SekolahService) UpdateKategoriSekolah(ctx context.Context, req *pb.UpdateKategoriSekolahRequest) (*pb.UpdateKategoriSekolahResponse, error) {
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
		// KategorisekolahId: int32(pbKategoriSekolah.KategoriSekolahId),
		KategorisekolahId:   pbKategoriSekolah.KategoriSekolahId,
		NamaKurikulum:       &pbKategoriSekolah.NamaKurikulum,
		NamaJurusan:         &pbKategoriSekolah.NamaJurusan,
		TahunAjaranId:       pbKategoriSekolah.TahunAjaranId,
		JurusanId:           &pbKategoriSekolah.JurusanId,
		KurikulumId:         pbKategoriSekolah.KurikulumId,
		JenjangPendidikanId: &pbKategoriSekolah.JenjangPendidikanId,
		NamaBidangKeahlian:  &pbKategoriSekolah.NamaBidangKeahlian,
		NamaProgramKeahlian: &pbKategoriSekolah.NamaProgramKeahlian,
		TingkatId:           pbKategoriSekolah.TingkatId,
		Jumlah:              &pbKategoriSekolah.Jumlah,
	}
	err = s.repoKategoriSekolah.Update(ctx, kategoriSekolahModel, Schemaname, "kategori_sekolah_id", fmt.Sprintf("%d", req.KategoriSekolah.KategoriSekolahId))
	if err != nil {
		return &pb.UpdateKategoriSekolahResponse{
			Message: "Gagal update kategori sekolah",
			Status:  false,
		}, nil
	}
	return &pb.UpdateKategoriSekolahResponse{
		Message: "Berhasil update kategori sekolah",
		Status:  true,
	}, nil
}

func (s *SekolahService) ProsesKategoriSekolahDanKelas(ctx context.Context, req *pb.ProsesKategoriSekolahDanKelasRequest) (*pb.ProsesKategoriSekolahDanKelasResponse, error) {
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

	kondisi := map[string]any{
		"tabel_kategori_sekolah.tahun_ajaran_id": tahunAjaranId,
	}
	// Ambil tabel kategori sekolah dan kelas
	modelKategoriSekolah, err := s.repoKategoriSekolah.FindWithRelations(ctx, Schemaname, nil, nil, kondisi, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	// jmlKelas := len(modelKategoriSekolah)
	var nmKelas string
	var kelompokKelas string
	for _, v := range modelKategoriSekolah {
		fmt.Println(v)
		// if v. != nil {
		nmKelas = utils.BuatSingkatan(*v.NamaJurusan)
		// for _, x := range v. {
		for y := 0; y < int(*v.Jumlah); y++ {
			// buatk kelas
			// jika jumlah > 1 buatkan nama
			if int(*v.Jumlah) > 1 {
				kelompokKelas, _ = excelize.ColumnNumberToName(y + 1)
			} else {
				kelompokKelas = ""
			}
			for z := 1; z <= 2; z++ {
				entity := &models.RombonganBelajar{
					RombonganBelajarId:  uuid.New(),
					SemesterId:          fmt.Sprintf("%s%d", tahunAjaranId, z),
					JurusanId:           v.JurusanId,
					NmKelas:             fmt.Sprintf("%s %s %s", utils.AngkaKeRomawi(int(v.TingkatId)), nmKelas, kelompokKelas),
					TingkatPendidikanId: v.TingkatId,
					KurikulumId:         &v.KurikulumId,
					JenisRombel:         utils.Int32ToPointer(1),
					NamaJurusanSp:       v.NamaJurusan,
				}
				cek := s.repoKelas.Save(ctx, entity, Schemaname)
				if cek != nil {
					return nil, err
				}
				// s.createPembelajaran(ctx, entity, Schemaname)
			}
		}
		// }
		// }
	}

	return &pb.ProsesKategoriSekolahDanKelasResponse{
		Status:  true,
		Message: "Berhasil membuat kelas sesusai kurikulum!",
	}, nil

}

// rev 1
func (s *SekolahService) createPembelajaran(ctx context.Context, modelRombel *models.RombonganBelajar, schemaname string) error {
	conditions := map[string]any{
		"ref.mata_pelajaran_kurikulum.kurikulum_id":          modelRombel.KurikulumId,
		"ref.mata_pelajaran_kurikulum.tingkat_pendidikan_id": modelRombel.TingkatPendidikanId,
	}
	modelsMapelKurikulum, err := s.repoMataPelajaranKurikulum.FindAllByConditions(ctx, "ref", conditions, 100, 0, nil)
	if err != nil {
		return fmt.Errorf("gagal mengambil mata pelajaran kurikulum: %w", err)
	}

	var modelsPembelajaran []models.Pembelajaran
	for _, v := range modelsMapelKurikulum {
		modelMapel, err := s.repoMataPelajaran.FindByID(ctx, strconv.FormatInt(int64(v.MataPelajaranId), 10), "ref", "mata_pelajaran_id")
		if err != nil {
			log.Printf("Gagal ambil mapel id %d: %v", v.MataPelajaranId, err)
			continue
		}
		if modelMapel.Nama == "" {
			log.Printf("Nama kosong untuk mapel id %d", v.MataPelajaranId)
			continue
		}
		modelPembelajaran := models.Pembelajaran{
			PembelajaranId:     uuid.New(),
			MataPelajaranId:    int(v.MataPelajaranId),
			SemesterId:         modelRombel.SemesterId,
			RombonganBelajarId: modelRombel.RombonganBelajarId,
			NamaMataPelajaran:  &modelMapel.Nama,
		}
		modelsPembelajaran = append(modelsPembelajaran, modelPembelajaran)
	}

	if len(modelsPembelajaran) == 0 {
		log.Println("Tidak ada pembelajaran yang bisa disimpan.")
		return nil
	}

	err = s.repoPembelajaran.SaveMany(ctx, schemaname, utils.SliceToPointer(modelsPembelajaran), 100)
	if err != nil {
		return fmt.Errorf("gagal menyimpan pembelajaran: %w", err)
	}
	return nil
}

func (s *SekolahService) GetKategoriMapel(ctx context.Context, req *pb.GetKategoriMapelRequest) (*pb.GetKategoriMapelResponse, error) {
	var err error
	log.Printf("sekolah_server/GetKategoriKelas, received data request: %+v\n", req)
	requiredFields := []string{"Schemaname", "TingkatPendidikan", "KurikulumId"}
	// Validasi request
	err = utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	Schemaname := req.GetSchemaname()
	// tahunAjaranId := req.GetTahunAjaranId()
	tingkatPendidikan := req.GetTingkatPendidikan()
	kurikulumId := req.GetKurikulumId()
	conditions := map[string]any{
		// "tabel_kategori_mapel.tahun_ajaran_id":    tahunAjaranId,
		"tabel_kategori_mapel.tingkat_pendidikan": tingkatPendidikan,
		"tabel_kategori_mapel.kurikulum_id":       kurikulumId,
	}
	modelKategoriMapel, err := s.repoKategoriMapel.FindAllByConditions(ctx, Schemaname, conditions, 100, 0, nil)
	if err != nil {
		return nil, err
	}
	pbKategoriMapel := utils.ConvertModelsToPB(modelKategoriMapel, func(item *models.KategoriMapel) *pb.KategoriMapel {
		return &pb.KategoriMapel{
			Id:                item.Id,
			MataPelajaranId:   item.MataPelajaranId,
			KurikulumId:       item.KurikulumId,
			JurusanId:         item.JurusanId,
			NmMapel:           item.NmMapel,
			TingkatPendidikan: item.TingkatPendidikan,
			TahunAjaranId:     item.TahunAjaranId,
		}
	})

	return &pb.GetKategoriMapelResponse{
		KategoriMapel: pbKategoriMapel,
		Status:        true,
		Message:       "Berhasil mengambil data mapel",
	}, nil
}

func (s *SekolahService) DeleteKategoriMapel(ctx context.Context, req *pb.DeleteKategoriMapelRequest) (*pb.DeleteKategoriMapelResponse, error) {
	var err error
	// Debugging: Cek nilai request yang diterima
	// log.Printf("Received Sekolah data request: %+v\n", req)
	requiredFields := []string{"Schemaname", "Id"}
	// Validasi request
	err = utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	Schemaname := req.GetSchemaname()
	kategoriMapelId := strconv.Itoa(int(req.GetId()))
	modelKategoriMapel, err := s.repoKategoriMapel.FindByID(ctx, kategoriMapelId, Schemaname, "id")
	if err != nil {
		return &pb.DeleteKategoriMapelResponse{
			Message: fmt.Sprintf("Gagal menghapus data kategori mapel: %s", err),
			Status:  false,
		}, nil
	}
	err = s.repoKategoriMapel.SoftDelete(ctx, Schemaname, modelKategoriMapel, kategoriMapelId, "id")
	if err != nil {
		return &pb.DeleteKategoriMapelResponse{
			Message: fmt.Sprintf("Gagal menghapus data kategori mapel: %s", err),
			Status:  false,
		}, nil
	}
	return &pb.DeleteKategoriMapelResponse{
		Message: "Kategori mapel berhasil dihapus!",
		Status:  true,
	}, nil

}
func (s *SekolahService) BatchDeleteKategoriMapel(ctx context.Context, req *pb.BatchDeleteKategoriMapelRequest) (*pb.BatchDeleteKategoriMapelResponse, error) {
	var err error
	// Debugging: Cek nilai request yang diterima
	// log.Printf("Received Sekolah data request: %+v\n", req)
	requiredFields := []string{"Schemaname", "Id"}
	// Validasi request
	err = utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	Schemaname := req.GetSchemaname()

	var ids []string
	for _, v := range req.GetId() {
		if v != "" {
			ids = append(ids, v)
		}
	}
	conditions := []struct {
		Query string
		Args  []any
	}{
		{"tabel_kategori_mapel.id IN ?", []any{ids}},
	}

	modelKategoriMapel, err := s.repoKategoriMapel.FindWithRelations(ctx, Schemaname, nil, nil, nil, conditions, nil, nil)
	if err != nil {
		return &pb.BatchDeleteKategoriMapelResponse{
			Message: fmt.Sprintf("Gagal menghapus data kategori mapel: %s", err),
			Status:  false,
		}, nil
	}
	err = s.repoKategoriMapel.SoftDeleteBatch(ctx, Schemaname, utils.SliceToPointer(modelKategoriMapel), ids, "id")
	if err != nil {
		return &pb.BatchDeleteKategoriMapelResponse{
			Message: fmt.Sprintf("Gagal menghapus data kategori mapel: %s", err),
			Status:  false,
		}, nil
	}
	return &pb.BatchDeleteKategoriMapelResponse{
		Message: "Kategori mapel berhasil dihapus!",
		Status:  true,
	}, nil

}
