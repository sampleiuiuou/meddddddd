package method

import (
	"log"
	"medapp/database"
	"medapp/types"
)

// add new Medicine err code AM01
func SaveBillMaster(lbm types.BillMasterReqStruct) types.CommonRespStruct {
	var resp types.CommonRespStruct
	db, err := database.LocalDBConnect()
	resp.Status = "S"
	resp.Msg = true
	if err != nil {
		log.Println("SBM01", err)
		resp.ErrMsg = err.Error()
		resp.Status = "E"
		resp.Msg = false
		return resp
	}
	defer db.Close()
	billmaster := `INSERT INTO medapp_bill_master (billNo, billDate, billAmount, billGst, netPrice, login_id, created_by, created_date, updated_by, updated_date) 
	VALUES (?, NOW(), ?, ?, ?, (SELECT login_id FROM medapp_login WHERE userid=?), ?, NOW(), ?, NOW())`
	result, err := db.Exec(billmaster, lbm.BillNo, lbm.BillAmount, lbm.BillGst, lbm.NetPrice, lbm.UserId, lbm.UserId, lbm.UserId)
	if err != nil {
		log.Println("SBM02", err)
		resp.ErrMsg = err.Error()
		resp.Status = "E"
		resp.Msg = false
		return resp
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("SBM03", err)
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
