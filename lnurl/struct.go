package lnurl

type LNStruct struct {
	MinSendable int      `json:"min_sendable"`
	MaxSendable int      `json:"max_sendable"`
	Metadata    []string `json:"metadata"`
	Tag         string   `json:"tag"`
	Callback    string   `json:"callback"`
}

type LNWStruct struct {
	MinWithdrawable int                    `json:"min_withdrawable"`
	MaxWithdrawable int                    `json:"max_withdrawable"`
	Metadata        map[string]interface{} `json:"metadata"`
	Tag             string                 `json:"tag"`
	Callback        string                 `json:"callback"`
	K1              string                 `json:"k1"`
}
