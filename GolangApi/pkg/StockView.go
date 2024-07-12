package method

import (
	"log"
	"medapp/database"
	"medapp/types"
)

func GetAllStock(resp *types.StockViewRespStruct) []types.StockViewStruct {

	var lstock types.StockViewStruct
	var lstockArr []types.StockViewStruct
	db, err := database.LocalDBConnect()
	resp.Status = "S"
	resp.Msg = true
	if err != nil {
		log.Println("SR01", err)
		resp.ErrMsg = err.Error()
		resp.Status = "E"
		resp.Msg = false
		return lstockArr
	}
	defer db.Close()
	stockview := `SELECT medicineName , brand , quantity , unitPrice
	from medapp_medicine_master mdm , medapp_stock ms
	where mdm.medicine_master_id=ms.medicine_master_id`
	rows, err := db.Query(stockview)
	if err != nil {
		log.Println("SR02", err)
		resp.ErrMsg = err.Error()
		resp.Status = "E"
		resp.Msg = false
		return lstockArr

	} else {
		for rows.Next() {
			err := rows.Scan(&lstock.MedicineName, &lstock.Brand, &lstock.Quantity, &lstock.UnitPrice)
			if err != nil {
				log.Println("SR03", err)
				resp.ErrMsg = err.Error()
				resp.Status = "E"
				resp.Msg = false
				return lstockArr
			} else {
				lstockArr = append(lstockArr, lstock)
			}
		}

		rows.Close()
		return lstockArr
	}

}
