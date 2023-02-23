package models

import "github.com/google/uuid"

type LNEntity struct {
	Model
	K1                 string      `json:"k1"`
	Identifier         string      `json:"identifier"`
	Url                string      `json:"url" gorm:"type:varchar(250)"`
	Qr                 string      `json:"qr" gorm:"type:varchar(250)"`
	Description        string      `json:"description"`
	SatMaxSendable     int         `json:"sat_max_sendable" gorm:"default:10000"`
	SatMinSendable     int         `json:"sat_min_sendable" gorm:"default:100"`
	SatMinWithdrawable int         `json:"sat_min_withdrawable" gorm:"default:100"`
	SatMaxWithdrawable int         `json:"sat_max_withdrawable" gorm:"default:1000"`
	LnurlTag           string      `json:"ln_type"`
	PaymentRequest     []LNInvoice `json:"payment_request" gorm:"foreignKey:PaymentID"`
}

func (lnurl LNEntity) String() string {
	return lnurl.Url
}

type LNInvoice struct {
	Model
	PaymentID uuid.UUID `json:"payment_id"`
	Pr        string    `json:"pr"`
	Status    string    `json:"status" gorm:"default:unsettled"`
}
