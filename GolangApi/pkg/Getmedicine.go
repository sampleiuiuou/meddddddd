package method

import (
	"log"
	"medapp/database"
	"medapp/types"
)

// get all medicine in medicinemasters
func Getmed(resp *types.MedicineRespStruct) []types.MedicineStruct {

	var lstock types.MedicineStruct
	var med []types.MedicineStruct
	db, err := database.LocalDBConnect()
	resp.Status = "S"
	resp.Msg = true
	if err != nil {
		log.Println("GM01", err)
		resp.ErrMsg = err.Error()
		resp.Status = "E"
		resp.Msg = false
		return med
	}
	defer db.Close()
	medicine := `SELECT NVL(m.MedicineName,""), NVL(m.Brand,""), NVL(m.medicine_master_id,0) FROM medicine_master m ;`
	rows, err := db.Query(medicine)
	if err != nil {
		log.Println("GM02", err)
		resp.ErrMsg = err.Error()
		resp.Status = "E"
		resp.Msg = false
		return med

	} else {
		for rows.Next() {
			err := rows.Scan(&lstock.MedicineName, &lstock.Brand, &lstock.MasterId)
			if err != nil {
				log.Println("GM03", err)
				resp.ErrMsg = err.Error()
				resp.Status = "E"
				resp.Msg = false
				return med
			} else {
				med = append(med, lstock)
			}
		}

		rows.Close()
		return med
	}

}
