package billdropdown

import (
	"encoding/json"
	"fmt"
	"log"
	"medapp/database"
	"net/http"
)

type BillStockStruct struct {
	MedicineName string `json:"medicinename"`
	Quantity     int    `json:"quantity"`
	Brand        string `json:"brand"`
	UnitPrice    int    `json:"unitprice"`
	Medicine_id  int    `json:"medicine_id"`
}
type BillStockRespStruct struct {
	BillStockArr []BillStockStruct `json:"billstockarr"`
	ErrMsg       string            `json:"errmsg"`
	Status       string            `json:"status"`
	Msg          string            `json:"msg"`
}

func BillDropDown(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
	log.Println("BillDropDown(+)")

	if r.Method == "PUT" {
		var lBillStockRespRec BillStockRespStruct

		lBillStockRespRec, err := GetBillStock()
		if err != nil {
			log.Println(err)
			lBillStockRespRec.Status = "E"
			lBillStockRespRec.ErrMsg = "Error :" + err.Error()
		} else {
			lBillStockRespRec.Status = "S"
			lBillStockRespRec.Msg = "Success"
		}
		data, err := json.Marshal(lBillStockRespRec)
		if err != nil {
			fmt.Fprintf(w, "Error taking data"+err.Error())
		} else {
			fmt.Fprint(w, string(data))
			fmt.Printf("%s", data)
		}
		log.Println("BillDropDown(-)")
	}
}

func GetBillStock() (BillStockRespStruct, error) {
	log.Println("GetBillStock(+)")

	var lBillStockRec BillStockStruct
	var lBillStockRespRec BillStockRespStruct

	db, err := database.LocalDBConnect()
	if err != nil {
		log.Println(err)
		lBillStockRespRec.Status = "E"
		lBillStockRespRec.ErrMsg = "Error :" + err.Error()
		return lBillStockRespRec, err
	} else {
		defer db.Close()
		sqlString := `SELECT mdm.medicineName , ms.quantity,mdm.brand,ms.unitPrice,mdm.medicine_master_id
		FROM medapp_medicine_master mdm , medapp_stock ms
		WHERE mdm.medicine_master_id = ms.medicine_master_id ;`
		rows, err := db.Query(sqlString)
		if err != nil {
			log.Println(err)
			lBillStockRespRec.Status = "E"
			lBillStockRespRec.ErrMsg = "Error :" + err.Error()
			return lBillStockRespRec, err
		} else {
			for rows.Next() {
				err := rows.Scan(&lBillStockRec.MedicineName, &lBillStockRec.Quantity, &lBillStockRec.Brand, &lBillStockRec.UnitPrice, &lBillStockRec.Medicine_id)
				if err != nil {
					log.Println(err)
					lBillStockRespRec.Status = "E"
					lBillStockRespRec.ErrMsg = "Error :" + err.Error()
					return lBillStockRespRec, err
				} else {
					lBillStockRespRec.BillStockArr = append(lBillStockRespRec.BillStockArr, lBillStockRec)
				}
			}
		}
		log.Println("GetBillStock(-)")
		return lBillStockRespRec, err
	}

}
