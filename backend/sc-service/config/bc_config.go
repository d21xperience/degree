package config

import (
	"errors"
)

// Config untuk koneksi blockchain
type BCConfig struct {
	NetworkId      uint32 // Untuk Ethereum & Quorum
	BlockchainType string // "ethereum", "quorum", atau "hyperledger"
	RPCURL         string // URL RPC untuk Ethereum/Quorum

	// WalletPath        string // Path untuk wallet Hyperledger Fabric
	// ConnectionProfile string // Connection profile untuk Hyperledger Fabric
	// // Untuk Hyperledger Fabric
	// FabricConfigPath string
	// FabricWallet     string
	// FabricIdentity   string
	CertPath         string
	KeyPath          string
	MSPID            string
	PeerHostOverride string
	Channel          string
}

// LoadConfig membaca environment variables
func LoadBCConfig(bcConfig *BCConfig) (*BCConfig, error) {
	// Validasi parameter berdasarkan jenis blockchain
	switch bcConfig.BlockchainType {
	case "ethereum", "quorum":
		if bcConfig.RPCURL == "" {
			return nil, errors.New("RPC_URL harus diisi untuk Ethereum/Quorum")
		}
	case "hyperledger fabric":
		if bcConfig.CertPath == "" || bcConfig.PeerHostOverride == "" || bcConfig.MSPID == "" {
			return nil, errors.New("FABRIC_CONFIG_PATH, FABRIC_WALLET, dan FABRIC_IDENTITY harus diisi untuk Hyperledger Fabric")
		}

	default:
		return nil, errors.New("BLOCKCHAIN_TYPE tidak valid: gunakan 'ethereum', 'quorum', atau 'hyperledger'")
	}

	return &BCConfig{
		NetworkId:        bcConfig.NetworkId,
		BlockchainType:   bcConfig.BlockchainType,
		RPCURL:           bcConfig.RPCURL,
		CertPath:         bcConfig.CertPath,
		PeerHostOverride: bcConfig.PeerHostOverride,
		// FabricWallet:     fabricWallet,
		// FabricIdentity:   fabricIdentity,
	}, nil
}
