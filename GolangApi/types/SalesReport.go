package types

// Request salesReport Particular Date Using this Structure
type SalesReportReqStruct struct {
	FromDate string `json:"from_date"`
	ToDate   string `json:"to_date"`
}

// Get ALl the salesReport in Particular Date Response Structure
type SalesReportRespStruct struct {
	BillNo       int     `json:"billno"`
	BillDate     string  `json:"billldate"`
	MedicineName string  `json:"medicine_name"`
	Quantity     int     `json:"qty"`
	Amount       float64 `json:"amount"`
}

// Response to Total  salesReport in Response Structure
type TotalSalesReportRespStruct struct {
	SaleseReportArr []SalesReportRespStruct `json:"salesreport"`
	ErrMsg          string                  `json:"errmsg"`
	Status          string                  `json:"status"`
	Msg             bool                    `json:"msg"`
}
