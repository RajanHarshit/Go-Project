package models

import "gorm.io/gorm"

type LoanApplication struct {
	gorm.Model
	Name        string  `json:"name"`
	Income      float64 `json:"income"`
	CreditScore int     `json:"credit_score"`
	Amount      float64 `json:"amount"`
	Status      string  `json:"status"`
}
