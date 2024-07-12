package updatemed

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"medapp/database"
	"net/http"
)

type UpdateMedReqStruct struct {
	MedicineName string `json:"medicinename"`
	Brand        string `json:"brand"`
	Quantity     string `json:"quantity"`
	UnitPrice    string `json:"unitprice"`
	User_id      string `json:"userid"`
}

type MedicineDetails struct {
}
type CommonRespStruct struct {
	ErrMsg string `json:"errmsg"`
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

func UpdateMed(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
	log.Println("UpdateMed(+)")

	if r.Method == "PUT" {
		var lUpdateMedReqRec UpdateMedReqStruct
		var lCommonRespRec CommonRespStruct

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("Error : " + err.Error())
			lCommonRespRec.Status = "E"
			lCommonRespRec.ErrMsg = " Error :" + err.Error()
		} else {
			err = json.Unmarshal(body, &lUpdateMedReqRec)
			log.Println(string(body))
			if err != nil {
				log.Println("Error :" + err.Error())
				lCommonRespRec.Status = "E"
				lCommonRespRec.ErrMsg = "Error :" + err.Error()
			} else {
				lCommonRespRec, err = UpdateMedMethod(lUpdateMedReqRec)
				if err != nil {
					log.Println("Error :" + err.Error())
					lCommonRespRec.Status = "E"
					lCommonRespRec.ErrMsg = "Error :" + err.Error()
				} else {
					lCommonRespRec.Status = "S"
					lCommonRespRec.Msg = "Success"
				}

			}
		}
		data, err := json.Marshal(lCommonRespRec)
		if err != nil {
			fmt.Fprintf(w, "Error taking data"+err.Error())
		} else {
			fmt.Fprint(w, string(data))
			fmt.Printf("%s", data)
		}
		log.Println("UpdateMed(-)")
	}

}

func UpdateMedMethod(pUpdateMedReqRec UpdateMedReqStruct) (CommonRespStruct, error) {
	log.Println("UpdateMedMethod(+)")
	var lCommonRespRec CommonRespStruct
	db, err := database.LocalDBConnect()
	if err != nil {
		log.Println(err)
		return lCommonRespRec, err
	} else {
		defer db.Close()
		sqlString := `UPDATE medapp_stock 
		SET quantity = quantity + ?, 
			unitPrice = ?, 
			updated_by =?, 
			updated_date = NOW() 
		WHERE medicine_master_id IN (
			SELECT medicine_master_id
			FROM medapp_medicine_master mdm
			WHERE mdm.medicineName = ? AND mdm.brand = ?
		);
		`
		rows, err := db.Exec(sqlString, pUpdateMedReqRec.Quantity, pUpdateMedReqRec.UnitPrice, pUpdateMedReqRec.User_id, pUpdateMedReqRec.MedicineName, pUpdateMedReqRec.Brand)
		if err != nil {
			lCommonRespRec.Status = "E"
			log.Println(err)
			lCommonRespRec.ErrMsg = "Error :" + err.Error()
		} else {
			affected, err := rows.RowsAffected()
			if err != nil {
				lCommonRespRec.Status = "E"
				log.Println(err)
				lCommonRespRec.ErrMsg = "Error :" + err.Error()
			} else {
				if affected == 0 {
					lCommonRespRec.Status = "E"
					log.Println(err)
					lCommonRespRec.ErrMsg = "Error No Rows Updated"
				} else {
					log.Println("Updated Successfully")
					lCommonRespRec.Status = "S"
					lCommonRespRec.Msg = "Success"
					log.Println("UpdateMedMethod(-)")
				}
			}
		}
		return lCommonRespRec, err

	}

}
