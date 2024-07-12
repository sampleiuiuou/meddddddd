package method

import (
	"log"
	"medapp/database"
	"medapp/types"
)

// add new Medicine err code AM01
func SaveBill(lbill types.BillDetailReqStruct) types.CommonRespStruct {
	var resp types.CommonRespStruct
	db, err := database.LocalDBConnect()
	resp.Status = "S"
	resp.Msg = true
	if err != nil {
		log.Println("SB01", err)
		resp.ErrMsg = err.Error()
		resp.Status = "E"
		resp.Msg = false
		return resp
	}
	defer db.Close()
	insertmedicine := `INSERT INTO medapp_bill_details (billNo, medicine_master_id, quantity, unitPrice, amount, created_by, created_date, updated_by, updated_date)
	VALUES (?, (SELECT medicine_master_id FROM medapp_medicine_master m WHERE m.medicineName = ?), ?, ?, ?, ?, NOW(), ?, NOW())`
	for _, bill := range lbill.BillDetailsArr {
		result, err := db.Exec(insertmedicine, bill.BillNo, bill.MedicineName, bill.Quantity, bill.UnitPrice, bill.Amount, bill.UserId, bill.UserId)
		if err != nil {
			log.Println("SB02", err)
			resp.ErrMsg = err.Error()
			resp.Status = "E"
			resp.Msg = false
			return resp
		}

		rowsAffected, err := result.RowsAffected()
		log.Println(rowsAffected)
		if err != nil {
			log.Println("SB03", err)
			resp.ErrMsg = err.Error()
			resp.Status = "E"
			resp.Msg = false
			return resp
		}

		if rowsAffected == 0 {
			log.Println("no row inserted ")
			resp.Status = "E"
			resp.Msg = false
			return resp

		}

	}
	return resp
}
