package method

import (
	"log"
	"medapp/database"
	"medapp/types"
)

// Update stock
func BillUpdateStock(upStock []types.UpdateStockStruct) types.CommonRespStruct {
	var resp types.CommonRespStruct
	log.Println(upStock)
	db, err := database.LocalDBConnect()
	resp.Status = "S"
	resp.Msg = true
	if err != nil {
		log.Println("BUS01", err)
		resp.ErrMsg = err.Error()
		resp.Status = "E"
		resp.Msg = false
		return resp
	}
	defer db.Close()
	updateStock := `UPDATE stock
	JOIN medicine_master ON stock.medicine_master_id = medicine_master.medicine_master_id
	SET stock.Quantity = NVL(stock.Quantity,0) - ?
	WHERE medicine_master.MedicineName = ?`

	for _, update := range upStock {
		result, err := db.Exec(updateStock, update.Quantity, update.MedicineName)
		if err != nil {
			log.Println("BUS02", err)
			resp.ErrMsg = err.Error()
			resp.Status = "E"
			resp.Msg = false
			return resp
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			log.Println("BUS03", err)
			resp.ErrMsg = err.Error()
			resp.Status = "E"
			resp.Msg = false
			return resp
		}

		if rowsAffected == 0 {
			log.Println("no row", update.MedicineName)
		}
	}
	return resp
}
