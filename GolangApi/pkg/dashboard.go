package method

import (
	"log"
	"medapp/database"
	"medapp/types"
)

func Biller(userid string) types.TodaySalesRespStruct {

	var resp types.TodaySalesRespStruct
	db, err := database.LocalDBConnect()
	resp.Status = "S"
	resp.Msg = true
	if err != nil {
		log.Println("B01", err)
		resp.ErrMsg = err.Error()
		resp.Status = "E"
		resp.Msg = false
		return resp
	}
	defer db.Close()
	todaysales := `SELECT COALESCE(SUM(BillAmount), 0) AS 'totalsales' FROM bill_master bm WHERE bm.BillDate = CURDATE() AND login_id=(SELECT login_id FROM login WHERE UserId=?)`
	yesterdaysales := `SELECT COALESCE(SUM(BillAmount), 0) AS 'totalsales'  FROM bill_master bm  WHERE bm.BillDate = DATE_SUB(CURDATE(), INTERVAL 1 DAY)  AND login_id=(SELECT login_id FROM login WHERE UserId=?)`
	rows, err := db.Query(todaysales, userid)
	rows1, err1 := db.Query(yesterdaysales, userid)
	if err != nil {
		log.Println("B02", err)
		resp.ErrMsg = err.Error()
		resp.Status = "E"
		resp.Msg = false
		return resp
	}
	if err1 != nil {
		log.Println("B02", err1)
		resp.ErrMsg = err1.Error()
		resp.Status = "E"
		resp.Msg = false
		return resp
	}
	defer rows.Close()
	defer rows1.Close()
	var totalSales float64
	var ysales float64
	for rows1.Next() {
		err := rows1.Scan(&ysales)
		if err != nil {
			log.Println(err)
			resp.ErrMsg = err.Error()
			resp.Status = "E"
			resp.Msg = false
		}
	}
	for rows.Next() {
		err := rows.Scan(&totalSales)
		if err != nil {
			log.Println(err)
			resp.ErrMsg = err.Error()
			resp.Status = "E"
			resp.Msg = false
		}
	}
	log.Println(totalSales)
	log.Println(ysales)
	if totalSales <= ysales {
		resp.Msg = false
	}
	resp.TodaySales = totalSales
	return resp

}
func Manager(userid string) types.InventryRespStruct {

	var resp types.InventryRespStruct
	db, err := database.LocalDBConnect()
	resp.Status = "S"
	resp.Msg = true
	if err != nil {
		log.Println("B01", err)
		resp.ErrMsg = err.Error()
		resp.Status = "E"
		resp.Msg = false
		return resp
	}
	defer db.Close()
	invent := `SELECT COALESCE(SUM(UnitPrice), 0) AS 'CurrentInventry', COALESCE(SUM(Quantity), 0) AS 'qty' FROM stock s ;`
	rows, err := db.Query(invent)

	if err != nil {
		log.Println("B02", err)
		resp.ErrMsg = err.Error()
		resp.Status = "E"
		resp.Msg = false
		return resp
	}
	defer rows.Close()

	var inventry float64
	var qty int
	for rows.Next() {
		err := rows.Scan(&inventry, &qty)
		if err != nil {
			log.Println(err)
			resp.ErrMsg = err.Error()
			resp.Status = "E"
			resp.Msg = false
		}
	}
	sales := Todaysales()
	log.Println(sales)
	resp.TodaySales = sales.TodaySales

	resp.Inventry = inventry * float64(qty)
	return resp

}

func Todaysales() types.TodaySalesRespStruct {

	var resp types.TodaySalesRespStruct
	db, err := database.LocalDBConnect()
	resp.Status = "S"
	resp.Msg = true
	if err != nil {
		log.Println("B01", err)
		resp.ErrMsg = err.Error()
		resp.Status = "E"
		resp.Msg = false
		return resp
	}
	defer db.Close()
	todaysales := `SELECT COALESCE(SUM(BillAmount), 0) AS 'totalsales' FROM bill_master bm`
	rows, err := db.Query(todaysales)
	if err != nil {
		log.Println("B02", err)
		resp.ErrMsg = err.Error()
		resp.Status = "E"
		resp.Msg = false
		return resp
	}

	defer rows.Close()
	var totalSales float64
	for rows.Next() {
		err := rows.Scan(&totalSales)
		if err != nil {
			log.Println(err)
			resp.ErrMsg = err.Error()
			resp.Status = "E"
			resp.Msg = false
		}
	}
	log.Println(totalSales)
	resp.TodaySales = totalSales
	return resp

}
