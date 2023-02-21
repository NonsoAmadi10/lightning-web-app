package models

import "database/sql/driver"

type lnType string

const (
	PAY      lnType = "pay"
	WITHDRAW lnType = "withdraw"
)

func (ct *lnType) Scan(value interface{}) error {
	*ct = lnType(value.([]byte))
	return nil
}

func (ct lnType) Value() (driver.Value, error) {
	return string(ct), nil
}

type LNEntity struct {
	Model
	K1                 string `json:"k1"`
	Identifier         string `json:"identifier"`
	Url                string `json:"url" gorm:"type:varchar(250)"`
	Qr                 string `json:"qr" gorm:"type:varchar(250)"`
	Description        string `json:"description"`
	SatMaxSendable     int    `json:"sat_max_sendable" gorm:"default:10000"`
	SatMinSendable     int    `json:"sat_min_sendable" gorm:"default:1"`
	SatMinWithdrawable int    `json:"sat_min_withdrawable"`
	SatMaxWithdrawable int    `json:"sat_max_withdrawable"`
	LnurlTag           lnType `gorm:"column:ln_type;type:enum('PAY','WITHDRAW')" json:"ln_type"`
}

func (lnurl LNEntity) String() string {
	return lnurl.Url
}
