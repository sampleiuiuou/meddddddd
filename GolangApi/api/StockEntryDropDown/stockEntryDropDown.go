package stockentrydropdown

import (
	"encoding/json"
	"fmt"
	"log"
	"medapp/database"
	"net/http"
)

type StockArrayStruct struct {
	MedicineName string `json:"medicinename"`
	Brand        string `json:"brand"`
}
type AddMedicineDropRespStruct struct {
	MedicineDetails []StockArrayStruct `json:"medicinedetails"`
	ErrMsg          string             `json:"errmsg"`
	Status          string             `json:"status"`
	Msg             string             `json:"msg"`
}

func StockEntryDropDown(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
	log.Println("StockEntryDropDown(+)")

	if r.Method == "PUT" {
		var lAddMedicineDropRespStruct AddMedicineDropRespStruct

		lAddMedicineDropRespStruct, err := MedicineDropDownFetch()
		if err != nil {
			log.Println(err)
			lAddMedicineDropRespStruct.Status = "E"
			lAddMedicineDropRespStruct.ErrMsg = "Error :" + err.Error()
		} else {
			lAddMedicineDropRespStruct.Status = "S"
			lAddMedicineDropRespStruct.Msg = "Success"
		}
		data, err := json.Marshal(lAddMedicineDropRespStruct)
		if err != nil {
			fmt.Fprintf(w, "Error taking data"+err.Error())
		} else {
			fmt.Fprint(w, string(data))
			fmt.Printf("%s", data)
		}
		log.Println("StockEntryDropDown(-)")
	}

}

func MedicineDropDownFetch() (AddMedicineDropRespStruct, error) {
	log.Println("MedicineDropDownFetch(+)")
	var lAddMedicineDropRespStruct AddMedicineDropRespStruct
	var lStockArrayRec StockArrayStruct

	db, err := database.LocalDBConnect()
	if err != nil {
		log.Println(err)
		lAddMedicineDropRespStruct.Status = "E"
		lAddMedicineDropRespStruct.ErrMsg = "Error :" + err.Error()
		return lAddMedicineDropRespStruct, err
	} else {
		defer db.Close()
		sqlString := `SELECT medicineName , brand
		from medapp_medicine_master;`
		rows, err := db.Query(sqlString)
		if err != nil {
			log.Println(err)
			lAddMedicineDropRespStruct.Status = "E"
			lAddMedicineDropRespStruct.ErrMsg = "Error :" + err.Error()
			return lAddMedicineDropRespStruct, err
		} else {
			for rows.Next() {
				err := rows.Scan(&lStockArrayRec.MedicineName, &lStockArrayRec.Brand)
				if err != nil {
					log.Println(err)
					lAddMedicineDropRespStruct.Status = "S"
					lAddMedicineDropRespStruct.ErrMsg = "Error :" + err.Error()
					return lAddMedicineDropRespStruct, err
				} else {
					lAddMedicineDropRespStruct.MedicineDetails = append(lAddMedicineDropRespStruct.MedicineDetails, lStockArrayRec)
				}
			}
		}
		log.Println("MedicineDropDownFetch(-)")
		return lAddMedicineDropRespStruct, err
	}

}
