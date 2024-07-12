package method

import (
	"log"
	"medapp/database"
	"medapp/types"
)

func GetAllBill(resp *types.GetAllBillDetailsRespStruct) []types.BillDetailStruct {
	var lbill types.BillDetailStruct
	var lbills []types.BillDetailStruct
	db, err := database.LocalDBConnect()
	resp.Status = "S"
	resp.Msg = true
	if err != nil {
		log.Println("GA01", err)
		resp.ErrMsg = err.Error()
		resp.Status = "E"
		resp.Msg = false
		return lbills
	}
	defer db.Close()
	bills := `SELECT  NVL(mm.MedicineName,""), NVL(bd.Amount,0) FROM  bill_master bm JOIN  bill_details bd ON bm.BillNo = bd.BillNo
	JOIN  medicine_master mm ON bd.medicine_master_id = mm.medicine_master_id;`
	rows, err := db.Query(bills)
	if err != nil {
		log.Println("GA02", err)
		resp.ErrMsg = err.Error()
		resp.Status = "E"
		resp.Msg = false
		return lbills

	} else {
		for rows.Next() {
			err := rows.Scan(&lbill.MedicineName, &lbill.Amount)
			if err != nil {
				log.Println("GA03", err)
				resp.ErrMsg = err.Error()
				resp.Status = "E"
				resp.Msg = false
				return lbills
			} else {
				lbills = append(lbills, lbill)
			}
		}

		rows.Close()
		return lbills
	}

}
