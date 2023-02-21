package lnurl

type LNStruct struct {
	MinSendable int      `json:"min_sendable"`
	MaxSendable int      `json:"max_sendable"`
	Metadata    []string `json:"metadata"`
	Tag         string   `json:"tag"`
	Callback    string   `json:"callback"`
}
