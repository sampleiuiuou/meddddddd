package method

import (
	"log"
	"medapp/database"
	"medapp/types"
)

// find the Sales in Perticular Date Using Query
func SalesReport(lreport types.SalesReportReqStruct) types.TotalSalesReportRespStruct {
	var report types.SalesReportRespStruct

	var lrep []types.SalesReportRespStruct
	var resp types.TotalSalesReportRespStruct
	db, err := database.LocalDBConnect()
	resp.Status = "S"
	resp.Msg = true
	if err != nil {
		log.Println("SR01", err)
		resp.ErrMsg = err.Error()
		resp.Status = "E"
		resp.Msg = false
		return resp
	}
	defer db.Close()
	salesreport := `SELECT bm.BillNo, NVL(bm.BillDate,""), NVL(mm.MedicineName,""), NVL(bd.Quantity,""), NVL(bd.Amount,"")
	FROM bill_master bm
	JOIN bill_details bd ON bm.BillNo = bd.BillNo
	JOIN medicine_master mm ON bd.medicine_master_id = mm.medicine_master_id
	WHERE bm.BillDate BETWEEN ? AND ?`
	rows, err := db.Query(salesreport, lreport.FromDate, lreport.ToDate)
	if err != nil {
		log.Println("SR02", err)
		resp.ErrMsg = err.Error()
		resp.Status = "E"
		resp.Msg = false
		return resp

	} else {

		for rows.Next() {
			err := rows.Scan(&report.BillNo, &report.BillDate, &report.MedicineName, &report.Quantity, &report.Amount)

			if err != nil {
				log.Println("SR03", err)
				resp.ErrMsg = err.Error()
				resp.Status = "E"
				resp.Msg = false
				return resp
			} else {
				lrep = append(lrep, report)

			}
		}
		resp.SaleseReportArr = lrep
		rows.Close()
		return resp
	}

}
