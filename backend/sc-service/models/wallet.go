package models

import "time"

type WalletData struct {
	Address string      `json:"address"`
	Chain   ChainInfo   `json:"chain"`
	Balance BalanceInfo `json:"balance"`
	Tokens  []TokenInfo `json:"tokens"`
	Gas     GasInfo     `json:"gas"`
	Meta    MetaInfo    `json:"meta"`
}

type ChainInfo struct {
	ChainId        uint64       `json:"chainId"`
	Name           string       `json:"name"`
	RPC            string       `json:"rpc"`
	Explorer       string       `json:"explorer"`
	NativeCurrency CurrencyInfo `json:"nativeCurrency"`
}

type CurrencyInfo struct {
	Symbol   string `json:"symbol"`
	Decimals uint8  `json:"decimals"`
}

type BalanceInfo struct {
	Wei       string `json:"wei"`
	Formatted string `json:"formatted"`
}

type TokenInfo struct {
	Contract string `json:"contract"`
	Symbol   string `json:"symbol"`
	Decimals uint8  `json:"decimals"`
	Balance  string `json:"balance"`
	LogoURI  string `json:"logoURI,omitempty"`
}

type GasInfo struct {
	GasPrice             string `json:"gasPrice"`
	MaxFeePerGas         string `json:"maxFeePerGas,omitempty"`
	MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas,omitempty"`
}

type MetaInfo struct {
	CreatedAt  time.Time `json:"createdAt"`
	Label      string    `json:"label"`
	IsContract bool      `json:"isContract"`
}

// ChainDetails represents the details of a blockchain network
type ChainDetails struct {
	Name        string
	ExplorerURL string
	IsTestnet   bool
}
