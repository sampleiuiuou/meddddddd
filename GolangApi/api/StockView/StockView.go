package stockview

import (
	"encoding/json"
	"fmt"
	"log"
	"medapp/database"
	"net/http"
)

type StockViewStruct struct {
	MedicineName string `json:"medicinename"`
	Brand        string `json:"brand"`
	Quantity     int    `json:"quantity"`
	UnitPrice    int    `json:"unitprice"`
}

type StockViewRespStruct struct {
	StockViewArr []StockViewStruct `json:"stockviewarr"`
	ErrMsg       string            `json:"errmsg"`
	Status       string            `json:"status"`
	Msg          string            `json:"msg"`
}

func StockView(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
	log.Println("StockView(+)")
	if r.Method == "PUT" {
		var lStockViewRespRec StockViewRespStruct
		lStockViewRespRec.Status = "S"
		lStockViewRespRec.Msg = "SUCCESS"
		lStockViewRespRec, err := StockViewFetch()
		if err != nil {
			log.Println("Error :" + err.Error())
			lStockViewRespRec.Status = "E"
			lStockViewRespRec.ErrMsg = "Error :" + err.Error()
		} else {
			lStockViewRespRec.Status = "S"
			lStockViewRespRec.Msg = "Success"
		}

		data, err := json.Marshal(lStockViewRespRec)
		if err != nil {
			fmt.Fprintf(w, "Error taking data"+err.Error())
		} else {
			fmt.Fprint(w, string(data))
			fmt.Printf("%s", data)
		}
		log.Println("StockView(-)")
	}
}
func StockViewFetch() (StockViewRespStruct, error) {
	log.Println("StockViewFetch(+)")
	var lStockViewRec StockViewStruct
	var lStockViewRespRec StockViewRespStruct

	db, err := database.LocalDBConnect()
	if err != nil {
		log.Println(err)
		return lStockViewRespRec, err
	} else {
		defer db.Close()
		sqlString := `SELECT medicineName , brand , quantity , unitPrice
		from medapp_medicine_master mdm , medapp_stock ms
		where mdm.medicine_master_id=ms.medicine_master_id`
		rows, err := db.Query(sqlString)
		if err != nil {
			log.Println(err)
			return lStockViewRespRec, err
		} else {
			for rows.Next() {
				err := rows.Scan(&lStockViewRec.MedicineName, &lStockViewRec.Brand, &lStockViewRec.Quantity, &lStockViewRec.UnitPrice)
				if err != nil {
					log.Println(err)
					return lStockViewRespRec, err
				} else {
					lStockViewRespRec.StockViewArr = append(lStockViewRespRec.StockViewArr, lStockViewRec)
				}
			}
			return lStockViewRespRec, err
		}
	}
}
