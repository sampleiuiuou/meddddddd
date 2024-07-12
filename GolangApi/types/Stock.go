package types

// Add MedicineName and Brand Request Structure
type MedicineMasterReqStruct struct {
	MedicineName string `json:"medicine_name"`
	Brand        string `json:"brand"`
	UserId       string `json:"userid"`
}
type MedicineStruct struct {
	MedicineName string `json:"medicine_name"`
	Brand        string `json:"brand"`
	MasterId     string `json:"medicine_master_id"`
}

// add the Stocks Request Structure
type StockReqStruct struct {
	Quantity  int     `json:"qty"`
	UnitPrice float64 `json:"unit_price"`
	MasterId  string  `json:"medicine_master_id"`
}
type MedicineRespStruct struct {
	MedArr []MedicineStruct `json:"medicine"`
	ErrMsg string           `json:"errmsg"`
	Status string           `json:"status"`
	Msg    bool             `json:"msg"`
}

// stock view structure
type StockViewStruct struct {
	MedicineName string  `json:"medicine_name"`
	Brand        string  `json:"brand"`
	Quantity     int     `json:"qty"`
	UnitPrice    float64 `json:"unit_price"`
}
type StockViewRespStruct struct {
	StockArr []StockViewStruct `json:"stockview"`
	ErrMsg   string            `json:"errmsg"`
	Status   string            `json:"status"`
	Msg      bool              `json:"msg"`
}

// Stock Update Reqest Structure
type StockUpdateStruct struct {
	MedicineName string  `json:"medicine_name"`
	Brand        string  `json:"brand"`
	Quantity     int     `json:"qty"`
	UnitPrice    float64 `json:"unit_price"`
	UserId       string  `json:"userid"`
}
