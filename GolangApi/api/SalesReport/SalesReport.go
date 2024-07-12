package salesreport

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"medapp/database"
	"net/http"
)

type SalesReportReqStruct struct {
	FromDate string `json:"fromdate"`
	ToDate   string `json:"todate"`
}
type SalesReportStruct struct {
	BillNo       string  `json:"billno"`
	BillDate     string  `json:"billdate"`
	MedicineName string  `json:"medicinename"`
	Quantity     int     `json:"quantity"`
	Amount       float64 `json:"amount"`
}
type SalesReportRespStruct struct {
	SalesArr []SalesReportStruct `json:"salesarr"`
	ErrMsg   string              `json:"errmsg"`
	Status   string              `json:"status"`
	Msg      string              `json:"msg"`
}

func SalesReport(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
	log.Println("SalesReport(+)")
	if r.Method == "PUT" {
		var lSalesReportReqRec SalesReportReqStruct
		var lSalesReportRespRec SalesReportRespStruct
		lSalesReportRespRec.Status = "S"
		lSalesReportRespRec.Msg = "SUCCESS"

		body, err := io.ReadAll(r.Body)
		log.Println(r.Body)
		if err != nil {
			log.Println(err)
			lSalesReportRespRec.Status = "E"
			lSalesReportRespRec.ErrMsg = "Error :" + err.Error()
		} else {
			err := json.Unmarshal(body, &lSalesReportReqRec)
			log.Println(lSalesReportReqRec)
			if err != nil {
				log.Println(err)
				lSalesReportRespRec.Status = "E"
				lSalesReportRespRec.ErrMsg = "Error :" + err.Error()
			} else {

				lSalesReportRespRec, err = SalesReportFetch(lSalesReportReqRec)

				if err != nil {
					log.Println(err)
					lSalesReportRespRec.Status = "E"
					lSalesReportRespRec.ErrMsg = "Error :" + err.Error()
				}

			}
		}
		data, err := json.Marshal(lSalesReportRespRec)
		if err != nil {
			fmt.Fprintf(w, "Error taking data"+err.Error())
		} else {
			fmt.Fprint(w, string(data))
			fmt.Printf("%s", data)
		}
		log.Println("SalesReport(-)")

	}
}

func SalesReportFetch(PSalesReportReqRec SalesReportReqStruct) (SalesReportRespStruct, error) {
	var lSalesReportRespRec SalesReportRespStruct
	var lSalesReportStruct SalesReportStruct
	db, err := database.LocalDBConnect()
	if err != nil {
		log.Println(err)
		return lSalesReportRespRec, err
	} else {
		defer db.Close()
		sqlString := `SELECT mbm.billNo , mbm.billDate , mdm.medicineName , mbd.quantity , mbd.amount 

		from medapp_bill_details mbd , medapp_medicine_master  mdm , medapp_bill_master mbm
	   
		where mbd.medicine_master_id=mdm.medicine_master_id AND mbm.billNo=mbd.billNo AND 
		mbm.billDate BETWEEN ?  AND  ?;`
		rows, err := db.Query(sqlString, PSalesReportReqRec.FromDate, PSalesReportReqRec.ToDate)
		if err != nil {
			log.Println(err)
			lSalesReportRespRec.Status = "E"
			lSalesReportRespRec.ErrMsg = "Error :" + err.Error()
			return lSalesReportRespRec, err
		} else {
			for rows.Next() {
				err := rows.Scan(&lSalesReportStruct.BillNo, &lSalesReportStruct.BillDate, &lSalesReportStruct.MedicineName, &lSalesReportStruct.Quantity, &lSalesReportStruct.Amount)
				if err != nil {
					log.Println(err)
					lSalesReportRespRec.Status = "E"
					lSalesReportRespRec.ErrMsg = "Error :" + err.Error()
					return lSalesReportRespRec, err
				} else {
					lSalesReportRespRec.Status = "S"
					lSalesReportRespRec.Msg = "SUCCESS"
					lSalesReportRespRec.SalesArr = append(lSalesReportRespRec.SalesArr, lSalesReportStruct)
				}
			}

		}
		return lSalesReportRespRec, err
	}
}
