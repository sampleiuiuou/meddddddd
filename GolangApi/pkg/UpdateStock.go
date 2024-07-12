package method

import (
	"log"
	"medapp/database"
	"medapp/types"
)

// Update stock
func UpdateStock(upStock types.StockUpdateStruct) types.CommonRespStruct {
	var resp types.CommonRespStruct
	db, err := database.LocalDBConnect()
	resp.Status = "S"
	resp.Msg = true
	if err != nil {
		log.Println("US01", err)
		resp.ErrMsg = err.Error()
		resp.Status = "E"
		resp.Msg = false
		return resp
	}
	defer db.Close()
	// var MasterId int
	// findId := `SELECT medicine_master_id from medicine_master m WHERE m.MedicineName=? and m.Brand=?`
	// rows, err := db.Query(findId, upStock.MedicineName, upStock.Brand)
	// if !rows.Next() {
	// 	resp.ErrMsg = err.Error()
	// 	resp.Status = "E"
	// 	resp.Msg = false
	// }
	// err1 := rows.Scan(&MasterId)
	// if err1 != nil {
	// 	log.Println("US02", err1)
	// 	resp.ErrMsg = err1.Error()
	// 	resp.Status = "E"
	// 	resp.Msg = false
	// 	return resp
	// }
	insertmedicine := `UPDATE  stock
	SET  Quantity=NVL(Quantity,0) +?,UnitPrice =?,updated_by =?,updated_date =now()
		WHERE medicine_master_id in(SELECT medicine_master_id  FROM medicine_master m
	WHERE  m.MedicineName=? and m.Brand=?)`
	result, err := db.Exec(insertmedicine, upStock.Quantity, upStock.UnitPrice, upStock.UserId, upStock.MedicineName, upStock.Brand)
	if err != nil {
		log.Println("US02", err)
		resp.ErrMsg = err.Error()
		resp.Status = "E"
		resp.Msg = false
		return resp
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("US03", err)
		resp.ErrMsg = err.Error()
		resp.Status = "E"
		resp.Msg = false
		return resp
	}
	if rowsAffected > 0 {
		return resp
	} else {
		resp.Status = "E"
		resp.Msg = false
		return resp
	}
}
