package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"sc-service/config"
	pb "sc-service/generated"
	"sc-service/models"
	"sc-service/repositories"
	"sc-service/services/clients"

	"sc-service/utils"
	"sc-service/websocket"
)

type BlockchainService struct {
	pb.UnimplementedBlockchainServiceServer
	config       *config.BCConfig // Konfigurasi runtime
	client       clients.BlockchainClient
	wsServer     *websocket.WebSocketServer
	repoContract *repositories.GenericRepository[models.Contract]
}

// Constructor untuk BlockchainService
func NewBlockchainService() *BlockchainService {
	repoContract := repositories.NewContractDataRepository(config.DB)
	return &BlockchainService{
		config:       &config.BCConfig{},
		client:       nil,
		wsServer:     nil,
		repoContract: repoContract,
	}
}

var networkConected bool

// SetConfig: Mengatur konfigurasi blockchain
func (s *BlockchainService) SetConfig(ctx context.Context, req *pb.SetConfigRequest) (*pb.SetConfigResponse, error) {
	if networkConected {
		return nil, nil
	}
	// Daftar field yang wajib diisi
	requiredFields := []string{"BcConfig"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	environment := req.BcConfig.Environment
	platform := req.BcConfig.Platform
	bcConfig := &config.BCConfig{
		NetworkId:      uint32(environment.ChainId),
		BlockchainType: strings.ToLower(platform.Name),
		RPCURL:         environment.RPCURL,
	}
	// Load konfigurasi dari environment variables
	cfg, err := config.LoadBCConfig(bcConfig)
	if err != nil {
		log.Fatalf("Gagal memuat konfigurasi: %v", err)
		return &pb.SetConfigResponse{
			Message: fmt.Sprintf("gagal memuat konfigurasi: %s ", err),
			Status:  false,
		}, nil
	}

	// Buat blockchain client sesuai config
	client, err := clients.CreateClientFactory(cfg)
	if err != nil {
		log.Fatalf("Gagal membuat klien: %v", err)
		return &pb.SetConfigResponse{
			Message: fmt.Sprintf("gagal membuat klien: %s ", err),
			Status:  false,
		}, nil
	}
	// Connect ke blockchain
	if err := client.Connect(); err != nil {
		// return nil, errors.New("gagal terhubung ke blockhain")
		return &pb.SetConfigResponse{
			Message: fmt.Sprintf("gagal terhubung ke blockchain: %s ", err),
			Status:  false,
		}, nil
	}
	s.client = client
	// Inisialisasi dan start WebSocket server jika belum ada
	if s.wsServer == nil {
		s.wsServer = websocket.NewWebSocketServer()

		// Start WebSocket server di port 8080 (bisa disesuaikan)
		go func() {
			if err := s.wsServer.Start("8080"); err != nil {
				log.Printf("Failed to start WebSocket server: %v", err)
			} else {
				log.Println("WebSocket server started successfully")
			}
		}()

		// Tunggu sebentar untuk memastikan server sudah start
		time.Sleep(100 * time.Millisecond)
	}
	networkConected = true
	// Kirim notifikasi bahwa koneksi berhasil
	if s.wsServer != nil && s.wsServer.IsRunning() {
		s.wsServer.BroadcastMessage(map[string]any{
			"type": "blockchain_connected",
			"data": map[string]any{
				"blockchain_type": cfg.BlockchainType,
				"rpc_url":         cfg.RPCURL,
				"network_id":      cfg.NetworkId,
				"timestamp":       time.Now().Unix(),
				"status":          "success",
			},
		})
	}

	// Mulai monitoring blockchain jika diperlukan
	go s.startBlockchainMonitoring()

	return &pb.SetConfigResponse{
		Message: fmt.Sprintf("berhasil terhubung ke blockchain: %s ", cfg.BlockchainType),
		Status:  true,
	}, nil
}

// GetNetworkID: Mendapatkan Network ID dari blockchain
func (s *BlockchainService) GetNetworkID(ctx context.Context, _ *pb.Empty) (*pb.NetworkIDResponse, error) {
	if s.client == nil {
		return nil, errors.New("client belum dikonfigurasi")
	}
	networkID, err := s.client.NetworkID(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.NetworkIDResponse{
		NetworkId: uint32(networkID.Uint64()),
	}, nil
}

func (s *BlockchainService) GetWalletInfo(ctx context.Context, req *pb.GetWalletInfoRequest) (*pb.GetWalletInfoResponse, error) {
	log.Printf("bc_account_service/GetWalletInfo received data from request: %+v\n", req)
	// Daftar field yang wajib diisi
	requiredFields := []string{"PublicAddress"}
	// Validasi request
	requiredFieldsResponse := utils.ValidateFields(req, requiredFields)
	if requiredFieldsResponse != nil {
		// return nil, requiredFieldsResponse
		return &pb.GetWalletInfoResponse{
			Status:     false,
			Message:    fmt.Sprintf("Error: %v", requiredFieldsResponse),
			WalletData: nil,
		}, nil
	}

	if s.client == nil {
		return &pb.GetWalletInfoResponse{
			Status:     false,
			Message:    "klien belum dikonfigurasi",
			WalletData: nil,
		}, nil
	}

	// Dapatkan informasi chain
	chainInfo, err := s.client.GetChainInfo(s.config.RPCURL)
	if err != nil {
		return nil, err
	}

	// Dapatkan informasi gas
	gasInfo, err := s.client.GetGasInfo()
	if err != nil {
		return nil, err
	}
	// Dapatkan balance ETH
	balance, err := s.client.GetBalance(req.PublicAddress)
	if err != nil {
		return &pb.GetWalletInfoResponse{
			Status:     false,
			Message:    fmt.Sprintf("Gagal mendapatkan balance: %v", err),
			WalletData: nil,
		}, nil
	}

	results := &pb.WalletData{
		Address: req.PublicAddress,
		Balance: &pb.BalanceInfo{
			Wei:       balance.Wei,
			Formatted: balance.Formatted,
		},
		Gas: &pb.GasInfo{
			GasPrice:             gasInfo.GasPrice,
			MaxFeePerGas:         gasInfo.MaxFeePerGas,
			MaxPriorityFeePerGas: gasInfo.MaxPriorityFeePerGas,
		},
		Chain: &pb.ChainInfo{
			ChainId:  chainInfo.ChainId,
			Name:     chainInfo.Name,
			Rpc:      chainInfo.RPC,
			Explorer: chainInfo.Explorer,
			// NativeCurrency: &pb.CurrencyInfo{
			// 	Symbol: ,
			// },
		},
		// Meta: &pb.MetaInfo{
		// 	CreatedAt: ,
		// },
	}
	return &pb.GetWalletInfoResponse{
		Status:     true,
		Message:    "Sukses mengakses wallet",
		WalletData: results,
	}, nil
}

// Fungsi untuk monitoring blockchain secara real-time
func (s *BlockchainService) startBlockchainMonitoring() {
	// Tunggu sebentar sebelum mulai monitoring
	time.Sleep(2 * time.Second)

	ticker := time.NewTicker(10 * time.Second) // Update setiap 10 detik
	defer ticker.Stop()

	// Gunakan for range pada ticker.C
	for range ticker.C {
		if s.client != nil && s.wsServer != nil && s.wsServer.IsRunning() {
			// Kirim informasi status koneksi
			networkInfo := s.getBlockchainNetworkInfo()
			s.wsServer.BroadcastMessage(map[string]any{
				"type": "network_info",
				"data": networkInfo,
			})
		}
	}
}

// Fungsi helper untuk mendapatkan informasi network
func (s *BlockchainService) getBlockchainNetworkInfo() map[string]any {

	info := map[string]any{
		"timestamp":    time.Now().Unix(),
		"client_count": 0,
	}

	if s.wsServer != nil {
		info["client_count"] = s.wsServer.GetClientCount()
	}

	// Tambahkan informasi spesifik blockchain jika tersedia
	if s.client != nil {
		networkInfo, err := s.client.GetEVMNetworkInfo()
		if err != nil {
			return nil
		}
		info["blockchain_status"] = "connected"
		info["network_id"] = networkInfo.NetworkID
		info["latest_block"] = networkInfo.LatestBlock
		info["block_time"] = networkInfo.BlockTime
		info["gas_price"] = networkInfo.GasPrice
		info["client_version"] = networkInfo.ClientVersion
		info["is_syncing"] = networkInfo.IsSyncing
		info["peer_count"] = networkInfo.PeerCount
		info["status"] = networkInfo.Status
		// info["status"] = networkInfo.Status
	}

	return info
}

// Tambahkan method untuk menghentikan WebSocket
func (s *BlockchainService) StopWebSocket() {
	if s.wsServer != nil {
		// Implement stop logic jika diperlukan
		s.wsServer.BroadcastMessage(map[string]interface{}{
			"type": "server_shutdown",
			"data": map[string]interface{}{
				"message":   "Server is shutting down",
				"timestamp": time.Now().Unix(),
			},
		})
	}
}

// GetContractEvents: Mendapatkan daftar event dari smart contract
// func (s *BlockchainService) GetContractEvents(ctx context.Context, req *pb.GetContractEventsRequest) (*pb.GetContractEventsResponse, error) {
// 	if s.client == nil {
// 		return nil, errors.New("client belum dikonfigurasi")
// 	}

// 	// Panggil client untuk mendapatkan event logs
// 	events, err := s.client.GetContractEvents(ctx, req.ContractAddress, req.Abi, req.EventName, req.FromBlock, req.ToBlock)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.GetContractEventsResponse{
// 		Events: events,
// 	}, nil
// }

// // TransferToken: Mengirim token ERC20 dari satu alamat ke alamat lain
// func (s *BlockchainService) TransferToken(ctx context.Context, req *pb.TransferTokenRequest) (*pb.TransferTokenResponse, error) {
// 	if s.client.client == nil {
// 		return nil, errors.New("client belum dikonfigurasi")
// 	}

// 	// Panggil fungsi transfer pada smart contract ERC20
// 	txHash, err := s.client.TransferToken(ctx, req.TokenAddress, req.From, req.To, req.Amount, req.PrivateKey, req.GasLimit)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.TransferTokenResponse{
// 		TxHash: txHash,
// 	}, nil
// }

// // SendETH: Mengirim ETH dari satu alamat ke alamat lain
// func (s *BlockchainService) SendETH(ctx context.Context, req *pb.SendETHRequest) (*pb.SendETHResponse, error) {
// 	if s.client.client == nil {
// 		return nil, errors.New("client belum dikonfigurasi")
// 	}
// 	// Daftar field yang wajib diisi
// 	requiredFields := []string{"From", "To", "Amount", "PrivateKey"}
// 	// Validasi request
// 	err := utils.ValidateFields(req, requiredFields)
// 	if err != nil {
// 		return nil, err
// 	}

// 	amount := new(big.Int)
// 	amount, ok := amount.SetString(req.Amount, 10)
// 	if !ok {
// 		return nil, errors.New("gagal mengonversi amount ke *big.Int")
// 	}
// 	// Kirim transaksi ETH
// 	txHash, err := s.client.SendETH(ctx, req.PrivateKey, req.To, amount)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.SendETHResponse{
// 		TxHash: txHash,
// 	}, nil
// }

// // GetTokenBalance: Mendapatkan saldo token ERC20 dari smart contract
// func (s *BlockchainService) GetTokenBalance(ctx context.Context, req *pb.GetTokenBalanceRequest) (*pb.GetTokenBalanceResponse, error) {
// 	if s.client.client == nil {
// 		return nil, errors.New("client belum dikonfigurasi")
// 	}
// 	// Daftar field yang wajib diisi
// 	requiredFields := []string{"TokenAddress", "OwnerAddress"}
// 	// Validasi request
// 	err := utils.ValidateFields(req, requiredFields)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if req.TokenAddress == "\"\"" || req.OwnerAddress == "\"\"" {
// 		return nil, errors.New("token dan owner address tidak boleh kosong")
// 	}

// 	// Panggil fungsi "balanceOf" dari kontrak ERC20
// 	balance, err := s.client.GetTokenBalance(ctx, req.TokenAddress, req.OwnerAddress)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.GetTokenBalanceResponse{
// 		Balance: balance.String(),
// 	}, nil
// }

// // CallContractMethod: Memanggil fungsi read-only pada smart contract
// func (s *BlockchainService) CallContractMethod(ctx context.Context, req *pb.CallContractMethodRequest) (*pb.CallContractMethodResponse, error) {
// 	if s.client.client == nil {
// 		return nil, errors.New("client belum dikonfigurasi")
// 	}

// 	// Panggil client untuk membaca data dari smart contract
// 	result, err := s.client.CallContractMethod(ctx, req.ContractAddress, req.Abi, req.Method, req.Params)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.CallContractMethodResponse{
// 		Result: result,
// 	}, nil
// }

// // GetContractOwner: Mendapatkan alamat pemilik dari smart contract (jika ada)
// func (s *BlockchainService) GetContractOwner(ctx context.Context, req *pb.GetContractOwnerRequest) (*pb.GetContractOwnerResponse, error) {
// 	if s.client.client == nil {
// 		return nil, errors.New("client belum dikonfigurasi")
// 	}

// 	// Panggil fungsi "owner()" pada smart contract
// 	owner, err := s.client.CallContractMethod(ctx, req.ContractAddress, req.Abi, "owner", nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.GetContractOwnerResponse{
// 		OwnerAddress: owner,
// 	}, nil
// }

// // GetContract: Mendapatkan informasi contract dari blockchain
// func (s *BlockchainService) GetContract(ctx context.Context, req *pb.GetContractRequest) (*pb.GetContractResponse, error) {
// 	if s.client.client == nil {
// 		return nil, errors.New("client belum dikonfigurasi")
// 	}
// 	// Daftar field yang wajib diisi
// 	requiredFields := []string{"ContractAddress"}
// 	// Validasi request
// 	err := utils.ValidateFields(req, requiredFields)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// Panggil client untuk mendapatkan informasi contract
// 	bytecode, abi, err := s.client.GetContract(ctx, req.ContractAddress)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.GetContractResponse{
// 		ContractAddress: req.ContractAddress,
// 		Bytecode:        bytecode,
// 		Abi:             abi,
// 	}, nil
// }

func (s *BlockchainService) DeployContract(ctx context.Context, req *pb.DeployContractRequest) (*pb.DeployContractResponse, error) {
	if s.client == nil {
		return nil, errors.New("client belum dikonfigurasi")
	}
	// Daftar field yang wajib diisi
	requiredFields := []string{"ContractRequest", "Password", "AbiName", "BinName"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	bcAccount := models.Account{
		Address:  req.ContractRequest.ContractAddress,
		Username: req.Username,
	}

	contractAddress, txHash, err := s.client.DeployContract(ctx, req.Password, req.AbiName, req.BinName, &bcAccount)
	if err != nil {
		return &pb.DeployContractResponse{
			Status:  false,
			Message: fmt.Sprintf("%v", err),
		}, nil
	}
	// contractAddress := "0x700b6A60ce7EaaEA56F065753d8dcB9653dbAD35"
	// txHash := "0x1ff08a7281ade814e914a333b22e0f682b7a1c94153c48a00a5b842198aa256f"
	// Simpan ke DB
	contractModel := models.Contract{
		ContractName:    req.ContractRequest.ContractName,
		ContractAddress: contractAddress,
		TxHash:          txHash,
		ContractOwner:   &req.ContractRequest.ContractOwner,
		OwnerAddress:    bcAccount.Address,
		NetworkID:       req.ContractRequest.Id,
	}
	simpan := s.repoContract.Save(ctx, &contractModel, "public")
	if simpan != nil {
		return &pb.DeployContractResponse{
			Status:  false,
			Message: fmt.Sprintf("%v", simpan),
		}, nil
	}
	return &pb.DeployContractResponse{
		ContractAddress: contractAddress,
		TxHash:          txHash,
		Status:          true,
		Message:         "behasil membuat kontrak",
	}, nil

}

func (s *BlockchainService) GetContract(ctx context.Context, req *pb.GetContractRequest) (*pb.GetContractResponse, error) {
	log.Printf("bc_service/GetContract received data from request: %+v\n", req)
	// Daftar field yang wajib diisi
	requiredFields := []string{"OwnerAddress"}
	// Validasi request
	requiredFieldsResponse := utils.ValidateFields(req, requiredFields)
	if requiredFieldsResponse != nil {
		return nil, requiredFieldsResponse
	}
	conditions := map[string]any{
		"owner_address": req.OwnerAddress,
	}
	contracts, err := s.repoContract.FindAllByConditions(ctx, "public", conditions, 100, 0)
	if err != nil {
		return nil, err
	}
	results := utils.ConvertModelsToPB(contracts, func(item *models.Contract) *pb.Contract {
		return &pb.Contract{
			Id:              item.ID,
			ContractName:    item.ContractName,
			OwnerAddress:    item.OwnerAddress,
			ContractAddress: item.ContractAddress,
			TxHash:          item.TxHash,
			NetworkId:       item.NetworkID,
			ContractOwner:   utils.SafeString(item.ContractOwner),
			IsActive:        item.IsActive,
		}
	})
	return &pb.GetContractResponse{
		Status:   true,
		Message:  "Berhasil mengambil kontrak",
		Contract: results,
	}, nil
}

func (s *BlockchainService) ActiveContract(ctx context.Context, req *pb.Empty) (*pb.ActiveContractResponse, error) {

	conditions := map[string]any{
		"is_active": true,
	}
	contracts, err := s.repoContract.FindAllByConditions(ctx, "public", conditions, 1, 0)
	if err != nil {
		return nil, err
	}

	return &pb.ActiveContractResponse{
		Status:  true,
		Message: "Berhasil mendapatkan contract aktif",
		Contract: &pb.Contract{
			Id:              contracts[0].ID,
			ContractName:    contracts[0].ContractName,
			ContractAddress: contracts[0].ContractAddress,
			ContractOwner:   utils.SafeString(contracts[0].ContractOwner),
			OwnerAddress:    utils.SafeString(contracts[0].ContractOwner),
			TxHash:          contracts[0].TxHash,
			NetworkId:       contracts[0].NetworkID,
			IsActive:        contracts[0].IsActive,
		},
	}, nil
}

// SendTransactionToContract: Mengirim data ke smart contract dengan memanggil fungsi tertentu
// func (s *BlockchainService) SendTransactionToContract(ctx context.Context, req *pb.SendTransactionToContractRequest) (*pb.SendTransactionToContractResponse, error) {
// 	if s.client.client == nil {
// 		return nil, errors.New("client belum dikonfigurasi")
// 	}

// 	// Kirim transaksi ke smart contract
// 	txHash, err := s.client.SendTransactionToContract(ctx, req.ContractAddress, req.Abi, req.Method, req.Params, req.PrivateKey, req.GasLimit)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.SendTransactionToContractResponse{
// 		TxHash: txHash,
// 	}, nil
// }

// GetConsensusAlgorithm: Mendapatkan algoritma konsensus (hanya untuk Quorum)
// func (s *BlockchainService) GetConsensusAlgorithm(ctx context.Context, _ *pb.Empty) (*pb.ConsensusAlgorithmResponse, error) {
// 	// Periksa apakah client adalah QuorumClient
// 	quorumClient, ok := s.client.(QuorumClient)
// 	if !ok {
// 		return nil, errors.New("fitur ini hanya tersedia untuk Quorum")
// 	}

// 	consensus, err := quorumClient.GetConsensusAlgorithm(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.ConsensusAlgorithmResponse{
// 		ConsensusAlgorithm: consensus,
// 	}, nil
// }

// ApproveToken: Memberikan izin kepada smart contract lain untuk menggunakan token ERC20
// func (s *BlockchainService) ApproveToken(ctx context.Context, req *pb.ApproveTokenRequest) (*pb.ApproveTokenResponse, error) {
// 	if s.client.client == nil {
// 		return nil, errors.New("client belum dikonfigurasi")
// 	}

// 	// Panggil fungsi "approve" dari kontrak ERC20
// 	txHash, err := s.client.SendTransactionToContract(ctx, req.TokenAddress, req.Abi, "approve", []string{req.Spender, req.Amount}, req.PrivateKey, req.GasLimit)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.ApproveTokenResponse{
// 		TxHash: txHash,
// 	}, nil
// }
// GetTokenAllowance: Mengecek jumlah token ERC20 yang telah diizinkan untuk digunakan oleh smart contract lain
// func (s *BlockchainService) GetTokenAllowance(ctx context.Context, req *pb.GetTokenAllowanceRequest) (*pb.GetTokenAllowanceResponse, error) {
// 	if s.client.client == nil {
// 		return nil, errors.New("client belum dikonfigurasi")
// 	}

// 	// Panggil fungsi "allowance" dari kontrak ERC20
// 	allowance, err := s.client.CallContractMethod(ctx, req.TokenAddress, req.Abi, "allowance", []string{req.Owner, req.Spender})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.GetTokenAllowanceResponse{
// 		Allowance: allowance,
// 	}, nil
// }

// // GetGasPrice: Mendapatkan harga gas saat ini di jaringan blockchain
// func (s *BlockchainService) GetGasPrice(ctx context.Context, req *pb.GetGasPriceRequest) (*pb.GetGasPriceResponse, error) {
// 	if s.client.client == nil {
// 		return nil, errors.New("client belum dikonfigurasi")
// 	}

// 	// Ambil harga gas dari client
// 	gasPrice, err := s.client.GetGasPrice(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.GetGasPriceResponse{
// 		GasPrice: gasPrice,
// 	}, nil
// }

// // GetNonce: Mendapatkan nonce dari alamat tertentu
// func (s *BlockchainService) GetNonce(ctx context.Context, req *pb.GetNonceRequest) (*pb.GetNonceResponse, error) {
// 	if s.client.client == nil {
// 		return nil, errors.New("client belum dikonfigurasi")
// 	}

// 	nonce, err := s.GetNonce(ctx, req.Address)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.GetNonceResponse{
// 		Nonce: nonce,
// 	}, nil
// }

// func (s *BlockchainService) validateRequest(req any, requiredFields []string, checkEmptyFields map[string]func() string) error {
// 	if s.client.client == nil {
// 		return errors.New("client belum dikonfigurasi")
// 	}

// 	// Validasi apakah field-field wajib ada
// 	if err := utils.ValidateFields(req, requiredFields); err != nil {
// 		return err
// 	}

// 	// Validasi apakah field wajib kosong ("" atau nilai lain yang dianggap kosong)
// 	for field, getter := range checkEmptyFields {
// 		if getter() == "" || getter() == "\"\"" { // Sesuaikan dengan format yang mungkin terjadi
// 			return fmt.Errorf("%s tidak boleh kosong", field)
// 		}
// 	}

//		return nil
//	}
//

// func (s *BlockchainService) GetEVMNetworkInfo(ctx context.Context, _ *pb.Empty) (*pb.GetEVMNetworkInfoResponse, error) {
// 	if s.client == nil {
// 		// return nil, errors.New("client belum dikonfigurasi")
// 		return &pb.GetEVMNetworkInfoResponse{
// 			Status:      true,
// 			Message:     "client belum dikonfigurasi",
// 			NetworkInfo: nil,
// 		}, nil
// 	}

// 	networkInfo, err := s.client.GetEVMNetworkInfo()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &pb.GetEVMNetworkInfoResponse{
// 		Status:  true,
// 		Message: "Berhasi mendapatkan informasi jaringan",
// 		NetworkInfo: &pb.NetworkInfo{
// 			ChainID:       networkInfo.ChainID.Int64(),
// 			NetworkID:     networkInfo.NetworkID.Int64(),
// 			LatestBlock:   networkInfo.LatestBlock,
// 			BlockTime:     networkInfo.BlockTime,
// 			GasPrice:      networkInfo.GasPrice.Int64(),
// 			ClientVersion: networkInfo.ClientVersion,
// 			IsSyncing:     networkInfo.IsSyncing,
// 			PeerCount:     networkInfo.PeerCount,
// 			Status:        networkInfo.Status,
// 			Timestamp:     networkInfo.Timestamp,
// 		},
// 	}, nil

// }
