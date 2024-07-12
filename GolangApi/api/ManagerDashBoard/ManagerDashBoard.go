package managerdashboard

import (
	"encoding/json"
	"fmt"
	"log"
	"medapp/database"
	"net/http"
)

type ManagerDashRespStruct struct {
	Todaysales     float64 `json:"todaysales"`
	Inventoryvalue float64 `json:"inventoryvalue"`
	ErrMsg         string  `json:"errmsg"`
	Status         string  `json:"status"`
}

func ManagerDashBoard(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
	log.Println("SalesReport(+)")
	if r.Method == "PUT" {
		var lManagerDashboardRespRec ManagerDashRespStruct
		lManagerDashboardRespRec, err := ManagerDashBoardMethod()
		if err != nil {
			lManagerDashboardRespRec.Status = "E"
			lManagerDashboardRespRec.ErrMsg = "Error :" + err.Error()
		} else {
			lManagerDashboardRespRec.Status = "S"
		}

		data, err := json.Marshal(lManagerDashboardRespRec)
		if err != nil {
			fmt.Fprintf(w, "Error taking data"+err.Error())
		} else {
			fmt.Fprint(w, string(data))
			fmt.Printf("%s", data)
		}
		log.Println("BillerDashBoard(-)")
	}
}
func ManagerDashBoardMethod() (ManagerDashRespStruct, error) {
	var lManagerDashboardRespRec ManagerDashRespStruct
	db, err := database.LocalDBConnect()
	if err != nil {
		log.Println(err)
		lManagerDashboardRespRec.ErrMsg = err.Error()
		lManagerDashboardRespRec.Status = "E"
		return lManagerDashboardRespRec, err
	} else {
		defer db.Close()
		sqlString := `Select nvl(sum(billAmount),0)
		from medapp_bill_master
		where billDate=current_date()`
		rows, err := db.Query(sqlString)
		if err != nil {
			log.Println(err)
			return lManagerDashboardRespRec, err
		} else {
			for rows.Next() {
				value := rows.Scan(&lManagerDashboardRespRec.Todaysales)
				if value != nil {
					log.Println(err)

					lManagerDashboardRespRec.Status = "E"
					return lManagerDashboardRespRec, err
				} else {

					lManagerDashboardRespRec.Status = "S"

				}
			}
		}
	}
	sqlString := `Select nvl(sum(quantity*unitPrice),0)
	              from medapp_stock`
	rows, err := db.Query(sqlString)
	if err != nil {
		log.Println(err)
		return lManagerDashboardRespRec, err
	} else {
		for rows.Next() {
			value := rows.Scan(&lManagerDashboardRespRec.Inventoryvalue)
			if value != nil {
				log.Println(value)
				lManagerDashboardRespRec.ErrMsg = "Error : " + value.Error()
				lManagerDashboardRespRec.Status = "E"
				return lManagerDashboardRespRec, err
			} else {
				lManagerDashboardRespRec.Status = "S"
			}
		}
	}
	return lManagerDashboardRespRec, err
}
