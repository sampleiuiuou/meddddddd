package addmed

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"medapp/database"
	"net/http"
)

type AddMedicineReqStruct struct {
	MedicineName string `json:"medicinename"`
	Brand        string `json:"brand"`
	User_id      string `json:"userid"`
}

type CommonRespStruct struct {
	ErrMsg string `json:"errmsg"`
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

func AddMed(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
	log.Println("AddMed(+)")
	var AddMedicineReqRec AddMedicineReqStruct
	var lCommonRespRec CommonRespStruct
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error : " + err.Error())
		lCommonRespRec.Status = "E"
		lCommonRespRec.ErrMsg = " Error :" + err.Error()
	} else {
		err = json.Unmarshal(body, &AddMedicineReqRec)
		log.Println(string(body))
		if err != nil {
			log.Println("Error :" + err.Error())
			lCommonRespRec.Status = "E"
			lCommonRespRec.ErrMsg = "Error :" + err.Error()
		} else {
			lCommonRespRec, err = AddMedicine(AddMedicineReqRec)
			if err != nil {
				log.Println("Error :" + err.Error())
				lCommonRespRec.Status = "E"
				lCommonRespRec.ErrMsg = "Error :" + err.Error()
			} else {
				lCommonRespRec.Status = "S"
				lCommonRespRec.Status = "Success"
			}
		}
		data, err := json.Marshal(lCommonRespRec)
		if err != nil {
			fmt.Fprintf(w, "Error taking data"+err.Error())
		} else {
			fmt.Fprint(w, string(data))
			fmt.Printf("%s", data)
		}
		log.Println("AddMed(-)")
	}
}
func AddMedicine(PAddMedicineReqRec AddMedicineReqStruct) (CommonRespStruct, error) {
	log.Println("AddMedicine(+)")
	var lCommonRespRec CommonRespStruct
	db, err := database.LocalDBConnect()
	if err != nil {
		lCommonRespRec.Status = "E"
		log.Println(err)
		lCommonRespRec.ErrMsg = "Error :" + err.Error()
		return lCommonRespRec, err
	} else {
		defer db.Close()
		sqlString := `Insert into medapp_medicine_master 
		(medicineName,brand,created_by,created_date, updated_by,updated_date)
		select ?,?,?,now(),?,now()
		where not exists (select medicine_master_id
		from medapp_medicine_master medmas
		where medmas.medicineName=? and medmas.brand=?);`
		rows, err := db.Exec(sqlString, PAddMedicineReqRec.MedicineName, PAddMedicineReqRec.Brand, PAddMedicineReqRec.User_id, PAddMedicineReqRec.User_id, PAddMedicineReqRec.MedicineName, PAddMedicineReqRec.Brand)
		if err != nil {
			lCommonRespRec.Status = "E"
			log.Println(err)
			lCommonRespRec.ErrMsg = "Error :" + err.Error()
			return lCommonRespRec, err
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
					lCommonRespRec.ErrMsg = "Error No Rows Inserted"
				} else {
					log.Println("Inserted Successfully")
					lCommonRespRec.Status = "S"
					lCommonRespRec.Msg = "Success"
				}
			}
		}
	}
	log.Println("AddMedicine(-)")
	return lCommonRespRec, err
}
