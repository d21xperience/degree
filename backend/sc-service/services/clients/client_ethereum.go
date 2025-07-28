package clients

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math"
	"math/big"
	"sc-service/config"
	"sc-service/models"
	"sc-service/repositories"
	"sc-service/utils"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type NetworkInfo struct {
	ChainID       *big.Int `json:"chain_id"`
	NetworkID     *big.Int `json:"network_id"`
	LatestBlock   uint64   `json:"latest_block"`
	BlockTime     uint64   `json:"block_time"`
	GasPrice      *big.Int `json:"gas_price"`
	ClientVersion string   `json:"client_version"`
	IsSyncing     bool     `json:"is_syncing"`
	PeerCount     uint64   `json:"peer_count"`
	Status        string   `json:"status"`
	Timestamp     int64    `json:"timestamp"`
}

// Default implementasi EthClient menggunakan go-ethereum dan RPC
type EthereumClient struct {
	rpcClient *rpc.Client
	client    *ethclient.Client
	repo      *repositories.GenericRepository[models.Account]
	// wsURL     string
}

func NewEthereumClient(cfg *config.BCConfig) (BlockchainClient, error) {
	if cfg.RPCURL == "" {
		return nil, fmt.Errorf("ethereum RPC URL tidak boleh kosong")
	}
	// Buat koneksi RPC
	rpcClient, err := rpc.Dial(cfg.RPCURL)
	if err != nil {
		return nil, fmt.Errorf("gagal menghubungkan ke RPC: %v", err)
	}
	repo := repositories.NewAccountRepository(config.DB)
	// Gunakan ethclient sebagai wrapper untuk RPC
	client := ethclient.NewClient(rpcClient)

	return &EthereumClient{
		rpcClient: rpcClient,
		client:    client, // Sekarang client diinisialisasi
		repo:      repo,
	}, nil
}

// Connect menghubungkan ke jaringan Ethereum
func (e *EthereumClient) Connect() error {
	_, err := e.client.NetworkID(context.Background())
	if err != nil {
		return fmt.Errorf("gagal terhubung ke jaringan Ethereum: %w", err)
	}
	// Cek koneksi dengan mendapatkan block terbaru
	header, err := e.client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("gagal mengambil header blok: %w", err)
	}

	log.Printf("Berhasil terhubung ke Ethereum. Block terbaru: %d\n", header.Number.Uint64())
	return nil
}

// func (c *EthereumClient) SendTransaction(ctx context.Context, tx *types.Transaction) error {
// 	return c.client.SendTransaction(ctx, tx)
// }

// Implementasi NetworkID
func (c *EthereumClient) NetworkID(ctx context.Context) (*big.Int, error) {
	var result string
	err := c.rpcClient.CallContext(ctx, &result, "net_version")
	if err != nil {
		return nil, err
	}
	id := new(big.Int)
	id.SetString(result, 10)
	return id, nil
}

func (e *EthereumClient) GetEVMNetworkInfo() (*NetworkInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Get chain ID
	chainID, err := e.client.ChainID(ctx)
	if err != nil {
		return nil, err
	}
	// Get network ID
	networkID, err := e.client.NetworkID(ctx)
	if err != nil {
		return nil, err
	}
	// Get latest block number
	latestBlock, err := e.client.BlockNumber(ctx)
	if err != nil {
		return nil, err
	}
	// Get gas price
	gasPrice, err := e.client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}
	// Get client version
	var clientVersion string
	cekVersion := e.rpcClient.CallContext(ctx, &clientVersion, "web3_clientVersion")
	if cekVersion != nil {
		log.Fatalf("Failed to get client version: %v", err)
		return nil, err
	}

	// Get sync status
	progress, err := e.client.SyncProgress(ctx)
	if err != nil {
		return nil, err
	}
	isSyncing := progress != nil

	// Get latest block for timestamp
	var blockTime uint64
	block, err := e.client.BlockByNumber(ctx, nil)
	if err == nil {
		blockTime = block.Time()
	}

	// Get peer count (via RPC call)
	var peerCountResult string
	err = e.rpcClient.CallContext(ctx, &peerCountResult, "net_peerCount")
	var peerCount uint64 = 0
	if err == nil && peerCountResult != "" {
		if val, err := utils.ParseHexUint64(peerCountResult); err == nil {
			peerCount = val
		}
	}

	return &NetworkInfo{
		ChainID:       chainID,
		NetworkID:     networkID,
		LatestBlock:   latestBlock,
		BlockTime:     blockTime,
		GasPrice:      gasPrice,
		ClientVersion: clientVersion,
		IsSyncing:     isSyncing,
		PeerCount:     uint64(peerCount),
		Status:        "connected",
		Timestamp:     time.Now().Unix(),
	}, nil
}

func (e *EthereumClient) SubscribeNewHeads(ch chan *types.Header) (func(), error) {
	sub, err := e.client.SubscribeNewHead(context.Background(), ch)
	if err != nil {
		return nil, err
	}

	return func() {
		sub.Unsubscribe()
	}, nil
}

// Implementasi SuggestGasPrice
func (c *EthereumClient) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	var result big.Int
	err := c.rpcClient.CallContext(ctx, &result, "eth_gasPrice")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *EthereumClient) GetBalance(address string) (*models.BalanceInfo, error) {
	account := common.HexToAddress(address)
	balance, err := c.client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return nil, err
	}
	// Konversi saldo dari wei ke ETH
	ethValue := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(math.Pow10(18)))
	return &models.BalanceInfo{
		Wei:       balance.String(),
		Formatted: ethValue.Text('f', 18),
	}, nil
}

func (c *EthereumClient) GetLatestBlock() (*big.Int, error) {
	block, err := c.client.BlockByNumber(context.Background(), nil) // `nil` untuk blok terbaru
	if err != nil {
		return nil, err
	}
	return block.Number(), nil
}

func (c *EthereumClient) SendETH(ctx context.Context, privateKeyHex, toAddress string, amount *big.Int) (string, error) {
	privateKey, err := utils.GetECDSAPrivateKey(privateKeyHex)
	// privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return "", err
	}
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	nonce, err := c.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}
	gasPrice, err := c.client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}
	tx := types.NewTransaction(nonce, common.HexToAddress(toAddress), amount, uint64(21000), gasPrice, nil)
	chainID, err := c.client.NetworkID(context.Background())
	if err != nil {
		return "", err
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return "", err
	}
	err = c.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}
	return signedTx.Hash().Hex(), nil
}

func (c *EthereumClient) SubscribeToEvents(contractAddress string) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(contractAddress)},
	}
	logs := make(chan types.Log)
	sub, err := c.client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatalf("Failed to subscribe: %v", err)
	}
	defer sub.Unsubscribe()

	for {
		select {
		case err := <-sub.Err():
			log.Printf("Subscription error: %v", err)
		case vLog := <-logs:
			log.Printf("New log: %+v", vLog)
		}
	}
}

func (c *EthereumClient) DeployContract(ctx context.Context, bytecode string, privateKey string, gasLimit uint64) (string, string, error) {
	if c.client == nil {
		return "", "", errors.New("ethereum client tidak dikonfigurasi")
	}

	// Konversi bytecode dari string ke format bytes
	contractBytecode := common.FromHex(bytecode)

	// Dapatkan nonce untuk transaksi
	fromAddress, _ := utils.GetAddressFromPrivateKey(privateKey) // Buat fungsi ini jika belum ada
	nonce, err := c.client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return "", "", err
	}

	// Dapatkan harga gas saat ini
	gasPrice, err := c.client.SuggestGasPrice(ctx)
	if err != nil {
		return "", "", err
	}

	// Buat transaksi untuk deploy contract
	tx := types.NewContractCreation(nonce, big.NewInt(0), gasLimit, gasPrice, contractBytecode)

	// Tanda tangani transaksi dengan private key
	signedTx, err := utils.SignTransaction(privateKey, tx) // Buat fungsi SignTransaction jika belum ada
	if err != nil {
		return "", "", err
	}

	// Kirim transaksi
	err = c.client.SendTransaction(ctx, signedTx)
	if err != nil {
		return "", "", err
	}

	// Kembalikan alamat contract & tx hash
	contractAddress := crypto.CreateAddress(fromAddress, nonce) // Buat alamat contract dari nonce
	return contractAddress.Hex(), signedTx.Hash().Hex(), nil
}

// getGasInfo mendapatkan informasi gas
func (c *EthereumClient) GetGasInfo() (*models.GasInfo, error) {
	ctx := context.Background()

	// Dapatkan gas price saat ini
	gasPrice, err := c.client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	gasInfo := &models.GasInfo{
		GasPrice: gasPrice.String(),
	}

	// Try to get EIP-1559 gas data
	tipCap, err := c.client.SuggestGasTipCap(ctx)
	if err != nil {
		// Network might not support EIP-1559, return legacy gas price
		return gasInfo, nil
	}

	header, err := c.client.HeaderByNumber(ctx, nil)
	if err != nil {
		// If we can't get header, just use tip cap as both values
		gasInfo.MaxFeePerGas = tipCap.String()
		gasInfo.MaxPriorityFeePerGas = tipCap.String()
		return gasInfo, nil
	}

	// Calculate max fee using more conservative multiplier (1.25x)
	baseFee := header.BaseFee
	if baseFee == nil {
		baseFee = big.NewInt(0)
	}

	// maxFee = (baseFee * 5 / 4) + maxPriorityFee
	baseFeeMultiplied := new(big.Int).Mul(baseFee, big.NewInt(5))
	baseFeeMultiplied.Div(baseFeeMultiplied, big.NewInt(4))
	maxFeePerGas := new(big.Int).Add(baseFeeMultiplied, tipCap)

	gasInfo.MaxFeePerGas = maxFeePerGas.String()
	gasInfo.MaxPriorityFeePerGas = tipCap.String()

	return gasInfo, nil
}

// getChainInfo mendapatkan informasi chain berdasarkan client
func (c *EthereumClient) GetChainInfo(rpcURL string) (*models.ChainInfo, error) {
	ctx := context.Background()

	// Dapatkan chain ID
	chainID, err := c.client.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	// Tentukan nama chain berdasarkan chain ID
	chainDetail := c.getChainDetails(chainID.Uint64())

	return &models.ChainInfo{
		ChainId:  chainID.Uint64(),
		Name:     chainDetail.Name,
		RPC:      rpcURL,
		Explorer: chainDetail.ExplorerURL,
		NativeCurrency: models.CurrencyInfo{
			Symbol:   "ETH",
			Decimals: 18,
		},
	}, nil
}

// getChainDetails retrieves chain details based on chain ID following best practices
func (c *EthereumClient) getChainDetails(chainID uint64) models.ChainDetails {
	// Using a map for better maintainability and easier additions
	chains := map[uint64]models.ChainDetails{
		1:        {Name: "Ethereum Mainnet", ExplorerURL: "https://etherscan.io", IsTestnet: false},
		5:        {Name: "Goerli Testnet", ExplorerURL: "https://goerli.etherscan.io", IsTestnet: true},
		11155111: {Name: "Sepolia Testnet", ExplorerURL: "https://sepolia.etherscan.io", IsTestnet: true},
		137:      {Name: "Polygon Mainnet", ExplorerURL: "https://polygonscan.com", IsTestnet: false},
		80001:    {Name: "Mumbai Testnet", ExplorerURL: "https://mumbai.polygonscan.com", IsTestnet: true},
		42161:    {Name: "Arbitrum One", ExplorerURL: "https://arbiscan.io", IsTestnet: false},
		421613:   {Name: "Arbitrum Goerli", ExplorerURL: "https://goerli.arbiscan.io", IsTestnet: true},
		56:       {Name: "BNB Smart Chain", ExplorerURL: "https://bscscan.com", IsTestnet: false},
		97:       {Name: "BNB Smart Chain Testnet", ExplorerURL: "https://testnet.bscscan.com", IsTestnet: true},
		10:       {Name: "Optimism", ExplorerURL: "https://optimistic.etherscan.io", IsTestnet: false},
		420:      {Name: "Optimism Goerli", ExplorerURL: "https://goerli-optimism.etherscan.io", IsTestnet: true},
		43114:    {Name: "Avalanche C-Chain", ExplorerURL: "https://snowtrace.io", IsTestnet: false},
		43113:    {Name: "Avalanche Fuji Testnet", ExplorerURL: "https://testnet.snowtrace.io", IsTestnet: true},
	}

	if details, exists := chains[chainID]; exists {
		return details
	}

	return models.ChainDetails{
		Name:        fmt.Sprintf("Unknown Network (ChainID: %d)", chainID),
		ExplorerURL: "https://etherscan.io",
		IsTestnet:   false,
	}
}

// // SendTransactionToContract mengirim transaksi ke smart contract
// func (c *EthereumClient) SendTransactionToContract(ctx context.Context, contractAddress, abiJSON, method string, params []string, privateKeyHex string, gasLimit uint64) (string, error) {
// 	//  Konversi private key dari hex ke ECDSA
// 	privateKey, err := crypto.HexToECDSA(privateKeyHex)
// 	if err != nil {
// 		return "", errors.New("gagal mengonversi private key: " + err.Error())
// 	}

// 	//  Dapatkan alamat pengirim dari private key
// 	publicKey := privateKey.Public().(*ecdsa.PublicKey)
// 	fromAddress := crypto.PubkeyToAddress(*publicKey)

// 	//  Dapatkan nonce akun pengirim
// 	nonce, err := c.client.PendingNonceAt(ctx, fromAddress)
// 	if err != nil {
// 		return "", errors.New("gagal mendapatkan nonce: " + err.Error())
// 	}

// 	//  Load ABI contract
// 	parsedABI, err := abi.JSON(strings.NewReader(abiJSON))
// 	if err != nil {
// 		return "", errors.New("gagal mem-parsing ABI: " + err.Error())
// 	}

// 	//  Konversi parameter menjadi data transaksi
// 	data, err := parsedABI.Pack(method, convertParams(params)...)
// 	if err != nil {
// 		return "", errors.New("gagal mengkodekan data transaksi: " + err.Error())
// 	}

// 	//  Dapatkan harga gas
// 	gasPrice, err := c.client.SuggestGasPrice(ctx)
// 	if err != nil {
// 		return "", errors.New("gagal mendapatkan harga gas: " + err.Error())
// 	}

// 	//  Buat transaksi
// 	toAddress := common.HexToAddress(contractAddress)
// 	tx := types.NewTransaction(nonce, toAddress, big.NewInt(0), gasLimit, gasPrice, data)

// 	//  Tanda tangani transaksi
// 	chainID, err := c.client.NetworkID(ctx)
// 	if err != nil {
// 		return "", errors.New("gagal mendapatkan chain ID: " + err.Error())
// 	}

// 	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), privateKey)
// 	if err != nil {
// 		return "", errors.New("gagal menandatangani transaksi: " + err.Error())
// 	}

// 	//  Kirim transaksi ke jaringan
// 	err = c.client.SendTransaction(ctx, signedTx)
// 	if err != nil {
// 		return "", errors.New("gagal mengirim transaksi: " + err.Error())
// 	}

// 	//  Kembalikan hash transaksi
// 	return signedTx.Hash().Hex(), nil
// }

// // TransferToken mengirim token ERC-20 ke alamat lain
// func (c *EthereumClient) TransferToken(ctx context.Context, tokenAddress, from, to, amountStr, privateKeyHex string, gasLimit uint64) (string, error) {
// 	//  Konversi private key dari hex ke ECDSA
// 	privateKey, err := crypto.HexToECDSA(privateKeyHex)
// 	if err != nil {
// 		return "", errors.New("gagal mengonversi private key: " + err.Error())
// 	}

// 	//  Dapatkan alamat pengirim dari private key
// 	publicKey := privateKey.Public().(*ecdsa.PublicKey)
// 	fromAddress := crypto.PubkeyToAddress(*publicKey)

// 	//  Pastikan alamat pengirim sesuai dengan private key
// 	if !strings.EqualFold(fromAddress.Hex(), from) {
// 		return "", errors.New("private key tidak cocok dengan alamat pengirim")
// 	}

// 	//  Dapatkan nonce akun pengirim
// 	nonce, err := c.client.PendingNonceAt(ctx, fromAddress)
// 	if err != nil {
// 		return "", errors.New("gagal mendapatkan nonce: " + err.Error())
// 	}

// 	//  Konversi jumlah token dari string ke *big.Int
// 	amount := new(big.Int)
// 	amount, ok := amount.SetString(amountStr, 10)
// 	if !ok {
// 		return "", errors.New("gagal mengonversi amount ke *big.Int")
// 	}

// 	//  ABI fungsi transfer ERC-20: transfer(address,uint256)
// 	erc20ABI := `[{"constant":false,"inputs":[{"name":"recipient","type":"address"},{"name":"amount","type":"uint256"}],"name":"transfer","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"}]`

// 	//  Parse ABI
// 	parsedABI, err := abi.JSON(strings.NewReader(erc20ABI))
// 	if err != nil {
// 		return "", errors.New("gagal mem-parsing ABI: " + err.Error())
// 	}

// 	//  Encode data untuk fungsi transfer ERC-20
// 	data, err := parsedABI.Pack("transfer", common.HexToAddress(to), amount)
// 	if err != nil {
// 		return "", errors.New("gagal mengkodekan data transaksi: " + err.Error())
// 	}

// 	//  Dapatkan harga gas
// 	gasPrice, err := c.client.SuggestGasPrice(ctx)
// 	if err != nil {
// 		return "", errors.New("gagal mendapatkan harga gas: " + err.Error())
// 	}

// 	//  Buat transaksi
// 	tokenContract := common.HexToAddress(tokenAddress)
// 	tx := types.NewTransaction(nonce, tokenContract, big.NewInt(0), gasLimit, gasPrice, data)

// 	//  Dapatkan Chain ID
// 	chainID, err := c.client.NetworkID(ctx)
// 	if err != nil {
// 		return "", errors.New("gagal mendapatkan chain ID: " + err.Error())
// 	}

// 	//  Tandatangani transaksi
// 	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), privateKey)
// 	if err != nil {
// 		return "", errors.New("gagal menandatangani transaksi: " + err.Error())
// 	}

// 	//  Kirim transaksi ke jaringan
// 	err = c.client.SendTransaction(ctx, signedTx)
// 	if err != nil {
// 		return "", errors.New("gagal mengirim transaksi: " + err.Error())
// 	}

// 	//  Kembalikan hash transaksi
// 	return signedTx.Hash().Hex(), nil
// }

// func (c *EthereumClient) CallContractMethod(ctx context.Context, contractAddress, abiStr, method string, params []string) (string, error) {
// 	// Parse ABI
// 	parsedABI, err := abi.JSON(strings.NewReader(abiStr))
// 	if err != nil {
// 		return "", fmt.Errorf("failed to parse ABI: %w", err)
// 	}

// 	// Convert parameters
// 	args := make([]interface{}, len(params))
// 	for i, param := range params {
// 		args[i] = param
// 	}

// 	// Pack the method call
// 	data, err := parsedABI.Pack(method, args...)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to pack method call: %w", err)
// 	}

// 	// Prepare contract call message
// 	contractAddr := common.HexToAddress(contractAddress) // Buat variabel terlebih dahulu
// 	msg := ethereum.CallMsg{
// 		To:   &contractAddr,
// 		Data: data,
// 	}

// 	// Call contract method
// 	result, err := c.client.CallContract(ctx, msg, nil)
// 	if err != nil {
// 		return "", fmt.Errorf("contract call failed: %w", err)
// 	}

// 	// Decode return value
// 	return hexutil.Encode(result), nil
// }

// func (c *EthereumClient) GetTokenBalance(ctx context.Context, tokenAddress, ownerAddress string) (*big.Int, error) {
// 	// ABI ERC20 standar untuk balanceOf
// 	erc20ABI := `[{"constant":true,"inputs":[{"name":"owner","type":"address"}],"name":"balanceOf","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"}]`

// 	// Load ABI
// 	parsedABI, err := abi.JSON(strings.NewReader(erc20ABI))
// 	if err != nil {
// 		return nil, fmt.Errorf("gagal memparsing ABI: %v", err)
// 	}

// 	// Encode data untuk memanggil balanceOf(owner)
// 	data, err := parsedABI.Pack("balanceOf", common.HexToAddress(ownerAddress))
// 	if err != nil {
// 		return nil, fmt.Errorf("gagal mengkodekan data: %v", err)
// 	}

// 	// Panggil kontrak ERC20
// 	callMsg := ethereum.CallMsg{
// 		To:   &common.Address{},
// 		Data: data,
// 	}

// 	if callMsg.To == nil || callMsg.To.Hex() == "0x0000000000000000000000000000000000000000" {
// 		return nil, fmt.Errorf("alamat kontrak tidak valid")
// 	}
// 	if len(callMsg.Data) == 0 {
// 		return nil, fmt.Errorf("data transaksi kosong, pastikan ABI dan parameter benar")
// 	}

// 	copy(callMsg.To[:], common.HexToAddress(tokenAddress).Bytes())
// 	if c.client == nil {
// 		return nil, fmt.Errorf("ethereum client belum dikonfigurasi")
// 	}
// 	// Eksekusi call ke kontrak
// 	result, err := c.client.CallContract(ctx, callMsg, nil)
// 	if err != nil {
// 		return nil, fmt.Errorf("gagal memanggil kontrak: %v", err)
// 	}
// 	log.Printf("🔍 Debug: CallContract result = %x", result) // Log hasil CallContract

// 	if len(result) == 0 {
// 		return nil, fmt.Errorf("gagal mendapatkan saldo: hasil kosong, pastikan kontrak valid dan alamat benar")
// 	}
// 	// Decode hasil
// 	outputs, err := parsedABI.Unpack("balanceOf", result)
// 	if err != nil {
// 		return nil, fmt.Errorf("gagal mendekode hasil: %v", err)
// 	}

// 	// Konversi hasil ke *big.Int
// 	balance := outputs[0].(*big.Int)
// 	return balance, nil
// }

// func (c *EthereumClient) GetContractEvents(ctx context.Context, contractAddress, abiStr, eventName string, fromBlock, toBlock uint64) ([]string, error) {
// 	// Parse ABI
// 	parsedABI, err := abi.JSON(strings.NewReader(abiStr))
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to parse ABI: %w", err)
// 	}

// 	// Ambil event berdasarkan nama
// 	event, exists := parsedABI.Events[eventName]
// 	if !exists {
// 		return nil, fmt.Errorf("event %s not found in ABI", eventName)
// 	}

// 	// Buat filter query
// 	query := ethereum.FilterQuery{
// 		FromBlock: big.NewInt(int64(fromBlock)),
// 		ToBlock:   big.NewInt(int64(toBlock)),
// 		Addresses: []common.Address{common.HexToAddress(contractAddress)},
// 		Topics:    [][]common.Hash{{event.ID}}, // Event signature hash
// 	}

// 	// Ambil log dari blockchain
// 	logs, err := c.client.FilterLogs(ctx, query)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to get logs: %w", err)
// 	}

// 	// Decode event logs
// 	var events []string
// 	for _, vLog := range logs {
// 		data, err := parsedABI.Unpack(event.Name, vLog.Data)
// 		if err != nil {
// 			log.Printf("failed to decode log data: %v", err)
// 			continue
// 		}
// 		events = append(events, fmt.Sprintf("%v", data))
// 	}

// 	return events, nil
// }

// func (c *EthereumClient) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
// 	return c.client.PendingNonceAt(ctx, account)
// }

// // Belum jalan
// func (c *EthereumClient) GetContract(ctx context.Context, contractAddress string) (string, string, error) {
// 	// Ambil bytecode dari contract
// 	bytecode, err := c.client.CodeAt(ctx, common.HexToAddress(contractAddress), nil)
// 	if err != nil {
// 		return "", "", err
// 	}

// 	wd, err := os.Getwd()
// 	if err != nil {
// 		log.Fatalf("Failed to get working directory: %v", err)
// 	}

// 	// Menggunakan filepath.Join agar sesuai dengan OS
// 	abiFile := filepath.Join(wd, "smartcontract", "ethbc", "build", "VervalIjazah.abi")
// 	//  Load ABI contract
// 	abi, err := GetABIFromFile(abiFile)
// 	if err != nil {
// 		return "", "", errors.New("gagal mem-parsing ABI: " + err.Error())
// 	}

// 	return hex.EncodeToString(bytecode), abi, nil
// }

// //	func (c *EthereumClient) CallSmartContract(client *ethclient, contractAddress, dataID string) (string, error) {
// //		contractAddr := common.HexToAddress(contractAddress)
// //		// Replace with your contract binding
// //		instance, err := verval_ijazah.NewVervalIjazah(contractAddr, client)
// //		if err != nil {
// //			return "", err
// //		}
// //		result, err := instance.SomeFunction(&bind.CallOpts{
// //			From: contractAddr,
// //		}, dataID)
// //		if err != nil {
// //			return "", err
// //		}
// //		return result, nil
// //	}
// //
// //	func (c *EthereumClient) DeploySmartContract(client *ethclient, privateKeyHex string) (common.Address, string, error) {
// //		res, add, err := DeploySmartContract(client, privateKeyHex)
// //		return res, add, err
// //	}
// //
// // =============================
// // =============Akun============
// func (c *EthereumClient) GenerateNewAccount(ctx context.Context, userId int32, password string) (map[string]interface{}, error) {
// 	key := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)

// 	a, err := key.NewAccount(password)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// simpan ke database
// 	pass, err := utils.EncryptPassword(password)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var results = map[string]interface{}{
// 		"Password":          pass,
// 		"KeystrokeFilename": filepath.Base(a.URL.Path),
// 		"Address":           a.Address.Hex(),
// 	}
// 	return results, nil
// }
// func (c *EthereumClient) DeployIjazahContract(ctx context.Context, privateKeyHex string) (contracAddress string, txHash string, err error) {

// 	pvKey, pubKey, err := ConvertStringPrivateKey(privateKeyHex)
// 	if err != nil {
// 		return "", "", err
// 	}

// 	// akun, err := c.GetAccounts(ctx, userId)
// 	// if err != nil {
// 	// 	return "","", err
// 	// }
// 	// // baca file utc
// 	// wd, err := os.Getwd()
// 	// if err != nil {
// 	// 	log.Fatalf("Failed to get working directory: %v", err)
// 	// }

// 	// // Menggunakan filepath.Join agar sesuai dengan OS
// 	// path := filepath.Join(wd, "wallet", akun[0].WalletFilename)
// 	// b, err := os.ReadFile(path)
// 	// if err != nil {
// 	// 	return "","", err
// 	// }

// 	// pas := utils.VerifyPassword(password, akun[0].Password)
// 	// if pas {
// 	// key, err := keystore.DecryptKey(b, password)
// 	// if err != nil {
// 	// 	return "","", err
// 	// }
// 	// add := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)

// 	nonce, err := c.client.PendingNonceAt(ctx, pubKey)
// 	if err != nil {
// 		return "", "", err
// 	}
// 	gasPrice, err := c.client.SuggestGasPrice(ctx)
// 	if err != nil {
// 		return "", "", err
// 	}
// 	chainId, err := c.client.ChainID(ctx)
// 	if err != nil {
// 		return "", "", err
// 	}
// 	gasLimit := uint64(3000000)
// 	// auth := TransactOptsAuth(key, chainId, gasPrice, nonce, gasLimit)
// 	auth, err := bind.NewKeyedTransactorWithChainID(pvKey, chainId)
// 	if err != nil {
// 		return "", "", err
// 	}
// 	auth.Nonce = big.NewInt(int64(nonce))
// 	auth.Value = big.NewInt(0)
// 	auth.GasLimit = gasLimit
// 	auth.GasPrice = gasPrice
// 	a, tx, _, err := verifikasiIjazah.DeployVerifikasiIjazah(auth, c.client)
// 	if err != nil {
// 		return "", "", err
// 	}
// 	fmt.Println(tx.Hash().Hex())
// 	return a.Hex(), tx.Hash().Hex(), nil

// 	// }
// 	// return "", "",nil
// }

// // func (c *EthereumClient) GetAccounts(ctx context.Context, userId int32, schemaname string) ([]*models.Account, error) {
// // 	var modelAccount []*models.Account
// // 	var err error
// // 	if userId == 0 {
// // 		modelAccount, err = c.repo.FindAll(ctx, schemaname, 100, 0)
// // 		if err != nil {
// // 			return nil, err
// // 		}
// // 		return modelAccount, nil
// // 	}

// // 	var condition = map[string]interface{}{
// // 		"user_id": userId,
// // 		"network_id":0,
// // 	}
// // 	modelAccount, err = c.repo.FindAllByConditions(ctx, schemaname, condition, 100, 0)
// // 	if err != nil {
// // 		return nil, err
// // 	}
// // 	return modelAccount, nil
// // }
// // // Fungsi untuk mengimpor akun dari private key
// // func (c *EthereumClient) ImportEthereumAccount(ctx context.Context, privateKeyHex string) (common.Address, *ecdsa.PrivateKey, error) {
// // 	// Hapus "0x" jika ada di awal
// // 	privateKeyHex = common.HexToHash(privateKeyHex).Hex()[2:]

// // 	// Konversi ke *ecdsa.PrivateKey
// // 	privateKey, err := crypto.HexToECDSA(privateKeyHex)
// // 	if err != nil {
// // 		return common.Address{}, nil, fmt.Errorf("gagal mengonversi private key: %w", err)
// // 	}

// // 	// Ambil public key dari private key
// // 	publicKey := privateKey.Public().(*ecdsa.PublicKey)

// // 	// Dapatkan alamat Ethereum dari public key
// // 	address := crypto.PubkeyToAddress(*publicKey)

// //		return address, privateKey, nil
// //	}
// //
// // =============================
// // IssueDegree mengeluarkan ijazah di Ethereum
// func (e *EthereumClient) IssueDegree(ctx context.Context, contractAddress string, degreeHash [32]byte, sekolah string, issueDate uint64, privateKey string, gasLimit uint64) (string, error) {
// 	//  Load ABI
// 	parsedABI, err := abi.JSON(strings.NewReader(abiJSON))
// 	if err != nil {
// 		return "", fmt.Errorf("error parsing ABI: %v", err)
// 	}

// 	//  Encode data untuk fungsi `issueDegree`
// 	data, err := parsedABI.Pack("issueDegree", degreeHash, sekolah, big.NewInt(int64(issueDate)))
// 	if err != nil {
// 		return "", fmt.Errorf("error packing data: %v", err)
// 	}

// 	//  Kirim transaksi menggunakan SendTransactionToContract
// 	txHash, err := SendTransactionToContract(ctx, e.client, contractAddress, data, privateKey, gasLimit)
// 	if err != nil {
// 		return "", fmt.Errorf("transaction failed: %v", err)
// 	}

// 	return txHash, nil
// }

// func (e *EthereumClient) IssueTranscript(
//     ctx context.Context,
//     contractAddress string,
//     degreeHash [32]byte,
//     transkrip map[string]uint8, // Mata pelajaran -> Nilai
//     privateKey string,
//     gasLimit uint64,
// ) (string, error) {
//     // Load ABI
//     parsedABI, err := abi.JSON(strings.NewReader(abiJSON))
//     if err != nil {
//         return "", fmt.Errorf("error parsing ABI: %v", err)
//     }

//     // Convert transcript map to Solidity-compatible types
//     subjects := make([]string, 0, len(transkrip))
//     grades := make([]*big.Int, 0, len(transkrip))

//     for subject, grade := range transkrip {
//         subjects = append(subjects, subject)
//         grades = append(grades, big.NewInt(int64(grade)))
//     }

//     // Encode data untuk fungsi issueTranscript
//     data, err := parsedABI.Pack("issueTranscript", degreeHash, subjects, grades)
//     if err != nil {
//         return "", fmt.Errorf("error packing data: %v", err)
//     }

//     // Kirim transaksi
//     txHash, err := SendTransactionToContract(ctx, e.client, contractAddress, data, privateKey, gasLimit)
//     if err != nil {
//         return "", fmt.Errorf("transaction failed: %v", err)
//     }

//     return txHash, nil
// }
// ContractService adalah service untuk interaksi dengan smart contract
// type ContractService struct {
// 	client EthClient
// }

// type SenderInfo struct {
// }

// // Constructor untuk ContractService
// func NewContractService(client EthClient) *ContractService {
// 	return &ContractService{client: client}
// }

// Fungsi untuk menambahkan transkrip nilai
// func (s *ContractService) AddTranscript(degreeHash [32]byte, mataPelajaran []string, nilai []uint8) {
// 	parsedABI, err := abi.JSON(strings.NewReader(abiJSON))
// 	if err != nil {
// 		log.Fatalf("Error parsing ABI: %v", err)
// 	}

// 	data, err := parsedABI.Pack("addTranscript", degreeHash, mataPelajaran, nilai)
// 	if err != nil {
// 		log.Fatalf("Error packing data: %v", err)
// 	}

// 	txHash, err := sendTransaction(s.client, data)
// 	if err != nil {
// 		log.Fatalf("Transaction failed: %v", err)
// 	}

// 	fmt.Printf("Transkrip berhasil ditambahkan! TxHash: %s\n", txHash.Hex())
// }

// func DeployContract(auth *bind.TransactOpts, client EthClient) (common.Address, string, error) {
// 	contractAddress, tx, _, err := verval_ijazah.DeployVervalIjazah(auth, client)
// 	if err != nil {
// 		return common.Address{}, "", err
// 	}
// 	return contractAddress, tx.Hash().Hex(), nil
// }

// func DeploySmartContract(client EthClient, privateKeyHex string, chainID *big.Int) (common.Address, string, error) {
// 	// Buat transactor
// 	auth, err := CreateTransactor(privateKeyHex, chainID)
// 	if err != nil {
// 		return common.Address{}, "", err
// 	}

// 	// Deploy smart contract
// 	contractAddress, txHash, err := DeployContract(auth, client)
// 	if err != nil {
// 		return common.Address{}, "", err
// 	}

// 	return contractAddress, txHash, nil
// }

// Fungsi untuk membuat transaksi dan menandatangani
// func sendTransaction(client EthClient, data []byte) (common.Hash, error) {
// 	// privateKeyHex := client.
// 	privateKey, err := crypto.HexToECDSA(privateKeyHex)
// 	if err != nil {
// 		return common.Hash{}, err
// 	}

// 	publicKey := privateKey.Public().(*ecdsa.PublicKey)
// 	fromAddress := crypto.PubkeyToAddress(*publicKey)

// 	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
// 	if err != nil {
// 		return common.Hash{}, err
// 	}

// 	gasPrice, err := client.SuggestGasPrice(context.Background())
// 	if err != nil {
// 		return common.Hash{}, err
// 	}

// 	tx := types.NewTransaction(nonce, common.HexToAddress(contractAddress), big.NewInt(0), 3000000, gasPrice, data)
// 	chainID, _ := client.NetworkID(context.Background())
// 	signedTx, _ := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)

// 	err = client.SendTransaction(context.Background(), signedTx)
// 	if err != nil {
// 		return common.Hash{}, err
// 	}

// 	return signedTx.Hash(), nil
// }

// func DeploySmartContract(client *ethclient.Client, privateKeyHex string) (common.Address, string, error) {
// 	privateKey, err := crypto.HexToECDSA(privateKeyHex)
// 	if err != nil {
// 		return common.Address{}, "", err
// 	}
// 	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337)) // Ganti Chain ID sesuai kebutuhan
// 	if err != nil {
// 		return common.Address{}, "", err
// 	}
// 	contractAddress, tx, _, err := verval_ijazah.DeployVervalIjazah(auth, client)
// 	if err != nil {
// 		return common.Address{}, "", err
// 	}
// 	return contractAddress, tx.Hash().Hex(), nil
// }

//	func CallSmartContract(client *ethclient.Client, contractAddress, dataID string) (string, error) {
//		contractAddr := common.HexToAddress(contractAddress)
//		// Replace with your contract binding
//		instance, err := verval_ijazah.NewVervalIjazah(contractAddr, client)
//		if err != nil {
//			return "", err
//		}
//		result, err := instance.Get(&bind.CallOpts{
//			From: contractAddr,
//		}, dataID)
//		if err != nil {
//			return "", err
//		}
//		return result, nil
//	}
