package services

import (
	"context"
	"fmt"
	"log"
	"path/filepath"
	"time"

	"sc-service/config"
	pb "sc-service/generated"
	"sc-service/models"
	"sc-service/repositories"
	"sc-service/utils"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

type BlockchainAccountService struct {
	pb.UnimplementedBlockchainAccountServiceServer
	// config   *Config          // Konfigurasi runtime
	// client   clients.BlockchainClient // Client yang digunakan (Ethereum/Quorum)
	schema   SchemaService
	repoAkun *repositories.GenericRepository[models.Account]
}

// Constructor untuk BlockchainAccountService
func NewBlockchainAccountService() *BlockchainAccountService {
	schemaRepository := repositories.NewSchemaRepository(config.DB)
	sekolahTenantRepository := repositories.NewsekolahTenantRepository(config.DB)
	schema := NewSchemaService(schemaRepository, sekolahTenantRepository)
	akunRepository := repositories.NewAccountRepository(config.DB)
	// client := NewBlockchainService()
	return &BlockchainAccountService{
		// config:   &Config{},
		schema:   schema,
		repoAkun: akunRepository,
	}
}

func (s *BlockchainAccountService) CreateBlockchainAccount(ctx context.Context, req *pb.CreateBlockchainAccountRequest) (*pb.CreateBlockchainAccountResponse, error) {
	log.Printf("blockchain_account_service/CreateBlockchainAccount received data from request: %+v\n", req)
	// Daftar field yang wajib diisi
	requiredFields := []string{"Password"}
	// Validasi request
	requiredFieldsResponse := utils.ValidateFields(req, requiredFields)
	if requiredFieldsResponse != nil {
		return nil, requiredFieldsResponse
	}

	address, err := createKeystore(s, req.Password)
	if err != nil {
		return &pb.CreateBlockchainAccountResponse{
			Status:  true,
			Message: fmt.Sprintf("%v", err),
			Account: &pb.BlockchainAccount{
				Address: address,
			},
		}, nil
	}

	return &pb.CreateBlockchainAccountResponse{
		Status:  true,
		Message: "Berhasil membuat akun",
		Account: &pb.BlockchainAccount{
			Address: address,
		},
	}, nil
}

func (s *BlockchainAccountService) GetBlockchainAccounts(ctx context.Context, req *pb.GetBlockchainAccountsRequest) (*pb.GetBlockchainAccountsResponse, error) {
	log.Printf("bc_account_service/GetBlochainAccounts received data from request: %+v\n", req)
	// Daftar field yang wajib diisi
	requiredFields := []string{"Username"}
	// Validasi request
	requiredFieldsResponse := utils.ValidateFields(req, requiredFields)
	if requiredFieldsResponse != nil {
		return nil, requiredFieldsResponse
	}
	conditions := map[string]any{
		"username": req.Username,
	}
	accounts, err := s.repoAkun.FindAllByConditions(ctx, "ref", conditions, 50, 0)
	if err != nil {
		log.Printf("Gagal mendapatkan akun: %v", err)
		return &pb.GetBlockchainAccountsResponse{
			Status:   false,
			Message:  fmt.Sprintf("gagal mendapatkan akun: %s", err),
			Accounts: nil,
		}, nil
	}

	results := utils.ConvertModelsToPB(accounts, func(item *models.Account) *pb.BlockchainAccount {
		return &pb.BlockchainAccount{
			Username:  item.Username,
			Address:   item.Address,
			Keystroke: item.Keystroke,
			Filename:  item.Filename,
			CreatedAt: item.CreatedAt.String(),
		}
	})

	return &pb.GetBlockchainAccountsResponse{
		Status:   true,
		Message:  "Berhasil mendapatkan akun",
		Accounts: results,
	}, nil
}

func (s *BlockchainAccountService) DeleteBlockchainAccount(ctx context.Context, req *pb.DeleteBlockchainAccountRequest) (*pb.DeleteBlockchainAccountResponse, error) {

	return &pb.DeleteBlockchainAccountResponse{
		Status:  true,
		Message: "Berhasil menghapus wallet",
	}, nil
}

// func (s *BlockchainAccountService) CreateWhiteListAccount(ctx context.Context, req *pb.CreateWhiteListAccountRequest) (*pb.CreateWhiteListAccountResponse, error) {
// 	log.Printf("bc_account_service/CreateWhiteListAccount received data from request: %+v\n", req)
// 	// Daftar field yang wajib diisi
// 	requiredFields := []string{"Username", "WalletAddress"}
// 	// Validasi request
// 	requiredFieldsResponse := utils.ValidateFields(req, requiredFields)
// 	if requiredFieldsResponse != nil {
// 		return nil, requiredFieldsResponse
// 	}

// 	rpcUrl := "http://localhost:8545"
// 	client, err := ethclient.Dial(rpcUrl)
// 	if err != nil {
// 		return &pb.CreateWhiteListAccountResponse{
// 			Status:  false,
// 			Message: fmt.Sprintf("Gagal koneksi ke jaringan. err: %w", err),
// 		}, nil
// 	}
// 	privateKey, err := crypto.HexToECDSA(os.Getenv("SUPERADMIN_PRIVATE_KEY"))

// 	publicKey := privateKey.Public().(*ecdsa.PublicKey)
// 	from := crypto.PubkeyToAddress(*publicKey)

// 	nonce, _ := client.PendingNonceAt(context.Background(), from)
// 	gasPrice, _ := client.SuggestGasPrice(context.Background())
// 	chainID, _ := client.NetworkID(context.Background())

// 	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
// 	auth.Nonce = big.NewInt(int64(nonce))
// 	auth.Value = big.NewInt(0)
// 	auth.GasLimit = uint64(300000)
// 	auth.GasPrice = gasPrice

// 	// Ambil address dan binding kontrak
// 	contractAddress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
// 	instance, err := abi.NewVerifikasiIjazah(contractAddress, client)
// 	if err != nil {
// 		return &pb.CreateWhiteListAccountResponse{
// 			Status:  false,
// 			Message: fmt.Sprintf("Gagal koneksi ke jaringan. err: %w", err),
// 		}, nil
// 	}

//		// Panggil daftarSekolah
//		sekolahAddr := common.HexToAddress(req.WalletAddress)
//		tx, err := instance.DaftarSekolah(auth, sekolahAddr)
//		if err != nil {
//			return &pb.CreateWhiteListAccountResponse{
//				Status:  false,
//				Message: fmt.Sprintf("Gagal koneksi ke jaringan. err: %w", err),
//			}, nil
//		}
//		log.Println("Whitelist berhasil:", tx.Hash().Hex())
//		return &pb.CreateWhiteListAccountResponse{
//			Status:  true,
//			Message: fmt.Sprintf("Berhasil mendeploy smartcontract dengan tx_hash: %s", tx.Hash().Hex()),
//		}, nil
//	}
func (s *BlockchainAccountService) GetWhiteListAccount(ctx context.Context, req *pb.GetWhiteListAccountRequest) (*pb.GetWhiteListAccountResponse, error) {
	return &pb.GetWhiteListAccountResponse{
		Status:  true,
		Message: "Berhasil",
	}, nil
}
func (s *BlockchainAccountService) DeleteWhiteListAccount(ctx context.Context, req *pb.DeleteWhiteListAccountRequest) (*pb.DeleteWhiteListAccountResponse, error) {
	return &pb.DeleteWhiteListAccountResponse{
		Status:  true,
		Message: "Berhasil",
	}, nil
}

func (s *BlockchainAccountService) ImportBlockchainAccount(ctx context.Context, req *pb.ImportBlockchainAccountRequest) (*pb.ImportBlockchainAccountResponse, error) {
	log.Printf("bc_account_service/ImportBlockhainAccount received data from request: %+v\n", req)
	// Daftar field yang wajib diisi
	requiredFields := []string{"PrivateKey", "Username"}
	// Validasi request
	requiredFieldsResponse := utils.ValidateFields(req, requiredFields)
	if requiredFieldsResponse != nil {
		return nil, requiredFieldsResponse
	}
	pubAdd, err := utils.PrivateKeyToPubAddress(req.PrivateKey)
	if err != nil {
		return &pb.ImportBlockchainAccountResponse{
			Status:    false,
			Message:   "Gagal mengimpor akun",
			BcAccount: nil,
		}, nil
	}
	// simpan ke database
	modelAkun := &models.Account{
		Address:  pubAdd.Hex(),
		Username: req.Username,
	}
	cek := s.repoAkun.Save(ctx, modelAkun, "ref")
	if cek != nil {

	}

	return &pb.ImportBlockchainAccountResponse{
		Status:  true,
		Message: "Behasil mengimpor akun",
		BcAccount: &pb.BlockchainAccount{
			Address:  fmt.Sprint(pubAdd.Hex()),
			Username: req.Username,
			// CreatedAt: time.Now().GoString(),
		},
	}, nil
}

func createKeystore(c *BlockchainAccountService, password string) (string, error) {
	key := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)

	a, err := key.NewAccount(password)
	if err != nil {
		return "", err
	}
	// simpan ke database
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// pass, err := utils.EncryptPassword(password)
	// if err != nil {
	// 	return "", err
	// }
	simpan := c.repoAkun.Save(ctx, &models.Account{
		Address: a.Address.Hex(),
		// Password:          pass,
		Filename: filepath.Base(a.URL.Path),
	}, "ref")
	if simpan != nil {
		return "", simpan
	}

	return string(a.Address.Hex()), nil

}
