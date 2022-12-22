package service

type DepositRequest struct {
	WalletId string  `json:"wallet_id"`
	Amount   float64 `json:"amount"`
}

type CheckResponse struct {
	WalletId       string  `json:"wallet_id"`
	Balance        float64 `json:"balance"`
	AboveThreshold bool    `json:"above_threshold"`
}
