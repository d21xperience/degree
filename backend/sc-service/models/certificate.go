package models

import (
	"time"

	"github.com/google/uuid"
)

type Certificate struct {
	CertificateID   uint      `gorm:"primaryKey" json:"certificate_id"`
	StudentID       uint      `gorm:"index" json:"student_id"`
	CertificateHash string    `gorm:"size:64;unique" json:"certificate_hash"`
	IPFSCID         string    `gorm:"size:64" json:"ipfs_cid"`
	IssueDate       time.Time `json:"issue_date"`
	IssuerName      string    `gorm:"size:100" json:"issuer_name"`
	NilaiRata       string    `gorm:"size:100" json:"nilai_rata"`
	BlockchainType  string    `gorm:"size:50" json:"blockchain_type"`
	TransactionHash string    `gorm:"size:64" json:"transaction_hash"`
	ContractAddress string    `gorm:"size:64" json:"contract_address"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	Verifications []Verification `gorm:"foreignKey:CertificateID" json:"verifications"`
	IPFSMetadata  IPFSMetadata   `gorm:"foreignKey:IPFSCID;references:IPFSCID" json:"ipfs_metadata"`
}

type Verification struct {
	VerificationID   uint      `gorm:"primaryKey" json:"verification_id"`
	CertificateID    uint      `gorm:"index" json:"certificate_id"`
	VerifiedBy       string    `gorm:"size:100" json:"verified_by"`
	VerificationDate time.Time `gorm:"autoCreateTime" json:"verification_date"`
	Result           bool      `json:"result"`
	Remarks          string    `gorm:"type:text" json:"remarks"`
}

type IjazahBc struct {
	ID                          uuid.UUID  `gorm:"primaryKey;type:uuid"` // atau varchar sesuai kebutuhan
	PesertaDidikID              uuid.UUID  `gorm:"not null"`
	Nama                        string     `gorm:"column:nama"`
	NIS                         string     `gorm:"column:nis"`
	NISN                        string     `gorm:"column:nisn"`
	NPSN                        string     `gorm:"column:npsn"`
	NomorIjazah                 string     `gorm:"column:nomor_ijazah"`
	TempatLahir                 string     `gorm:"column:tempat_lahir"`
	TanggalLahir                *time.Time `gorm:"column:tanggal_lahir"`
	NamaOrtuwali                string     `gorm:"column:nama_ortuwali"`
	PaketKeahlian               string     `gorm:"column:paket_keahlian"`
	KabupatenKota               string     `gorm:"column:kabupatenkota"`
	Provinsi                    string     `gorm:"column:provinsi"`
	ProgramKeahlian             string     `gorm:"column:program_keahlian"`
	SekolahPenyelenggaraUjianUS string     `gorm:"column:sekolah_penyelenggara_ujian_us"`
	SekolahPenyelenggaraUjianUN string     `gorm:"column:sekolah_penyelenggara_ujian_un"`
	AsalSekolah                 string     `gorm:"column:asal_sekolah"`
	TempatIjazah                string     `gorm:"column:tempat_ijazah"`
	TanggalIjazah               *time.Time `gorm:"column:tanggal_ijazah"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (IjazahBc) TableName() string {
	return "ijazah_bc"
}

type DegreeData struct {
	ID             uint      `gorm:"primaryKey;autoIncrement"`
	IjazahID       uuid.UUID `gorm:"not null"`
	Ijazah         IjazahBc  `gorm:"foreignKey:IjazahID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	DegreeHash     string
	TxHash         string
	IpfsURL        string
	BcType         string
	LinkBcExplorer string
	TahunAjaranId  string     `gorm:"column:tahun_ajaran_id"`
	SekolahId      *uuid.UUID `gorm:"column:sekolah_id"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	// UrlBCExplorerEther  string `gorm:"column:url_bc_explorerEther"`
	// UrlBCExplorerQuorum string `gorm:"column:url_bc_explorerQuorum"`
	// UrlBCExplorerFabric string `gorm:"column:url_bc_explorerFabric"`
}

func (DegreeData) TableName() string {
	return "degree_data"
}
