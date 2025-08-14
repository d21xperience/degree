package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NetworkType string

const (
	Mainnet NetworkType = "mainnet"
	Testnet NetworkType = "testnet"
	Private NetworkType = "private"
	Local   NetworkType = "local"
)

type AccountType string

const (
	ImportAccount NetworkType = "imported"
	Keystore      NetworkType = "keystore"
)

type BCPlatform struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	NmBlockchain string    `gorm:"type:varchar(50);not null;unique"`
	Active       bool      `gorm:"default:false"`
	CreatedAt    time.Time `gorm:"default:now()"`
	UpdatedAt    time.Time `gorm:"default:now()"`
}

func (BCPlatform) TableName() string {
	return "blockchain_platform"
}

// WalletTransaction menyimpan riwayat transaksi pengguna
type WalletTransaction struct {
	ID          uint      `gorm:"primaryKey"`
	AccountID   uint      `gorm:"not null;index"`       // Relasi ke akun
	Hash        string    `gorm:"uniqueIndex;not null"` // Hash transaksi Ethereum
	Amount      float64   // Jumlah ETH atau token lain
	TokenSymbol string    // Simbol token jika bukan ETH
	Timestamp   time.Time // Waktu transaksi
}

// Wallet model untuk database
type Wallet struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Address    string    `gorm:"uniqueIndex;not null" json:"address"`
	PrivateKey string    `gorm:"not null" json:"privateKey"`
	Keystore   string    `gorm:"type:text" json:"keystore"`
	Filename   string    `json:"filename"`
	Password   string    `gorm:"-" json:"-"` // tidak disimpan di database
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

// Network menyimpan informasi jaringan blockchain
type Network struct {
	ID          uint32      `gorm:"primaryKey"`
	Name        string      `gorm:"size:100;not null;unique"`            // Nama jaringan (Ethereum, Polygon, BSC)
	ChainID     int64       `gorm:"not null;unique"`                     // Chain ID jaringan
	RPCURL      string      `gorm:"size:255;not null"`                   // URL RPC jaringan
	ExplorerURL string      `gorm:"size:255"`                            // URL block explorer
	Symbol      string      `gorm:"size:10;not null"`                    // Simbol token utama (ETH, MATIC, BNB)
	Type        NetworkType `gorm:"type:network_type;default:'mainnet'"` // ENUM di PostgreSQL
	CreatedAt   time.Time
	UpdatedAt   time.Time
	// kolom Activate digunakan untuk menampilkan jaringan pada saat pemilihan jaringan
	Activate bool `gorm:"default:false"`
	// kolom NetworkAvailable digunakan jika logic bisnis sudah dibuat, saat ini baru tersedia ethereum, quorum dan hyperledger fabric
	Available    bool   `gorm:"default:false"`
	Architecture string `gorm:"size:100;not null"` // Nama jaringan (Ethereum, Polygon, BSC)
}

func (Network) TableName() string {
	return "ref.networks"
}

// Account menyimpan alamat Ethereum pengguna
type Account struct {
	Id         uint32 `gorm:"primaryKey"`
	Address    string `gorm:"column:address"`     // Alamat Ethereum unik
	Username   string `gorm:"column:username"`    // Nama pengguna opsional
	PrivateKey string `gorm:"column:private_key"` // Nama pengguna opsional
	Keystroke  string `gorm:"column:keystroke"`
	Filename   string `gorm:"column:filename"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	// NetworkID    uint32  `gorm:"not null;index"` // Relasi ke jaringan blockchain
	// Network      Network `gorm:"foreignKey:NetworkID"`
	// Organization string  `json:"organization,omitempty"` // Untuk Hyperledger Fabric
	// IsActive     bool    `json:"isActive,omitempty"`
	// Password     string `gorm:"size:100"`             // digunakan untuk keystore file
	// Type         AccountType `gorm:"type:account_type;default:'IMPORTED'"`
	// UserID       uint32 // data dari admin sekolah
}



// Buat ENUM secara manual sebelum AutoMigrate()
func Migrate(db *gorm.DB) error {
	// Buat ENUM jika belum ada
	query1 := `DO $$ BEGIN IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'network_type') THEN CREATE TYPE network_type AS ENUM ('MAINNET', 'TESTNET','PRIVATE'); END IF; END $$;`
	query2 := `DO $$ BEGIN IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'account_type') THEN CREATE TYPE account_type AS ENUM ('IMPORTED', 'KEYSTORE'); END IF; END $$;`
	err := db.Exec(query1 + query2).Error
	if err != nil {
		return err
	}
	return db.AutoMigrate(&SekolahTenant{}, &SchemaLog{})
}
