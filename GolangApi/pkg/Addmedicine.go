package method

import (
	"log"
	"medapp/database"
	"medapp/types"
)

// add new Medicine err code AM01
func Addmedicine(addMed types.MedicineMasterReqStruct) types.CommonRespStruct {
	var resp types.CommonRespStruct
	db, err := database.LocalDBConnect()
	resp.Status = "S"
	resp.Msg = true
	if err != nil {
		log.Println("AM01", err)
		resp.ErrMsg = err.Error()
		resp.Status = "E"
		resp.Msg = false
		return resp
	}
	defer db.Close()
	insertmedicine := `INSERT INTO medicine_master (MedicineName, Brand, created_by, created_date, updated_by, updated_date)
	SELECT ?, ?, ?, NOW(),?,NOW()
	WHERE NOT EXISTS ( SELECT medicine_master_id FROM medicine_master m WHERE m.MedicineName = ? AND m.Brand = ? )`
	result, err := db.Exec(insertmedicine, addMed.MedicineName, addMed.Brand, addMed.UserId, addMed.UserId, addMed.MedicineName, addMed.Brand)
	if err != nil {
		log.Println("AM02", err)
		resp.ErrMsg = err.Error()
		resp.Status = "E"
		resp.Msg = false
		return resp
	}
	rowsAffected, err := result.RowsAffected()
	masterid, err1 := result.LastInsertId()
	if err != nil && err1 != nil {
		log.Println("AM03", err)
		resp.ErrMsg = err.Error()
		resp.Status = "E"
		resp.Msg = false
		return resp
	}
	if rowsAffected > 0 {
		insertStock := `INSERT INTO stock (Quantity, medicine_master_id, UnitPrice, created_by, created_date, updated_by, updated_date ) VALUES (?, ?, ?, ?, NOW(),?,NOW());`
		_, err := db.Exec(insertStock, 0, masterid, 0, addMed.UserId, addMed.UserId)
		if err != nil {
			log.Println("AM03", err)
			resp.ErrMsg = err.Error()
			resp.Status = "E"
			resp.Msg = false
			return resp
		}
		return resp
	} else {
		resp.Status = "E"
		resp.Msg = false
		return resp
	}
}
