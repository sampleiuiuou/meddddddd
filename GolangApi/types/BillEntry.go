package types

// Insert Bill Details Structure
type BillDetailStruct struct {
	BillNo       string  `json:"billno"`
	MedicineName string  `json:"medicine_name"`
	Quantity     int     `json:"qty"`
	UnitPrice    float64 `json:"unit_price"`
	Amount       float64 `json:"amount"`
	UserId       string  `json:"userid"`
}
type BillDetailReqStruct struct {
	BillDetailsArr []BillDetailStruct `json:"bill_details"`
}
type GetAllBillDetailsRespStruct struct {
	BillDetailsArr []BillDetailStruct `json:"bill"`
	ErrMsg          string           `json:"errmsg"`
    Status          string           `json:"status"`
    Msg             bool             `json:"msg"`
}

//  Get Bill once add bill  **not Mandatory**

// type BillStockRespStruct struct {
//     BilldetailStock BillDetailReqStruct `json:"billdetails"`
//     ErrMsg          string           `json:"errmsg"`
//     Status          string           `json:"status"`
//     Msg             bool             `json:"msg"`
// }

// Insert Bill Master Bill details Request Structure
type BillMasterReqStruct struct {
	BillNo     string  `json:"billno"`
	BillDate   string  `json:"billdate"`
	BillAmount float64 `json:"bill_amount"`
	BillGst    float64 `json:"gst"`
	NetPrice   float64 `json:"net_price"`
	UserId     string  `json:"userid"`
}

// get Total of AMOUNT,Bill_GST, NetPrice structure
type BillDetailsTotalStruct struct {
	Amount   float64 `json:"amount"`
	BillGst  float64 `json:"gst"`
	NetPrice float64 `json:"net_price"`
}

// get Bill Details Response Structure
type BillDetailsTotalRespStruct struct {
	BillTotal BillDetailsTotalStruct `json:"bill_total"`
	ErrMsg    string                 `json:"errmsg"`
	Status    string                 `json:"status"`
	Msg       bool                   `json:"msg"`
}

// Get Request Arrya of UpdateStockStruct structure
type UpdateStockReqStruct struct {
	UpdateStockArr []UpdateStockStruct `json:"update_stock"`
}

// store updated stock MEdicineName and Quantity structure
type UpdateStockStruct struct {
	Quantity     int    `json:"qty"`
	MedicineName string `json:"medicine_name"`
}
