package model

type ReasonMsg struct {
	TraceId string
	Reason  string
}

type MerchantInfo struct {
	TraceId       string
	TopMerchantId string
	MerchantId    string
	OutletId      string
}

type Result struct {
	TraceId       string
	TopMerchantId string
	MerchantId    string
	OutletId      string
	Reason        string
}
