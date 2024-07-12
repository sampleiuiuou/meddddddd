package billerdashboard

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"medapp/database"
	"net/http"
)

type BillerDashboardReqStruct struct {
	Login_id int `json:"login_id"`
}

type BillerDashboardRespStruct struct {
	TodaySales     float64 `json:"todaysales"`
	YesterdaySales float64 `json:"yesterdaysales"`
	ErrMsg         string  `json:"errmsg"`
	Status         string  `json:"status"`
	Msg            string  `json:"msg"`
}

func BillerDashBoard(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
	log.Println("BillerDashBoard(+)")

	if r.Method == "PUT" {
		var lbillerDashBoardReqRec BillerDashboardReqStruct
		var lbillerDashBoardRespRec BillerDashboardRespStruct

		body, err := io.ReadAll(r.Body)
		log.Println(string(body))
		if err != nil {
			log.Println(err)
			lbillerDashBoardRespRec.Status = "E"
			lbillerDashBoardRespRec.ErrMsg = "Error :" + err.Error()
			lbillerDashBoardRespRec.Msg = "400 Error"
		} else {
			err := json.Unmarshal(body, &lbillerDashBoardReqRec)
			if err != nil {
				log.Println(err)
				lbillerDashBoardRespRec.Status = "E"
				lbillerDashBoardRespRec.ErrMsg = "Error :" + err.Error()
				lbillerDashBoardRespRec.Msg = "400 Error"
			} else {
				var methodErr error
				lbillerDashBoardRespRec, methodErr = BillerDashBoardMethod(lbillerDashBoardReqRec)

				if methodErr != nil {
					lbillerDashBoardRespRec.Status = "E"
					lbillerDashBoardRespRec.ErrMsg = "Error :" + methodErr.Error()
					lbillerDashBoardRespRec.Msg = "400 Error"
				} else {
					lbillerDashBoardRespRec.Status = "S"
					lbillerDashBoardRespRec.Msg = "Success"
				}
			}
		}
		data, err := json.Marshal(lbillerDashBoardRespRec)
		if err != nil {
			fmt.Fprintf(w, "Error taking data"+err.Error())
		} else {
			fmt.Fprint(w, string(data))
			fmt.Printf("%s", data)
		}
		log.Println("BillerDashBoard(-)")
	}
}
func BillerDashBoardMethod(PbillerDashBoardReqRec BillerDashboardReqStruct) (BillerDashboardRespStruct, error) {
	log.Println("BillerDashBoardMethod(+)")
	var pbillerDashBoardRespRec BillerDashboardRespStruct
	db, err := database.LocalDBConnect()
	if err != nil {
		log.Println(err)
		return pbillerDashBoardRespRec, err
	}
	defer db.Close()

	var ysal, tsal float64
	sqlString := `SELECT COALESCE(SUM(mbm.billAmount), 0) AS total_sales
	FROM medapp_bill_master mbm
	WHERE mbm.login_id = ? AND mbm.billDate = CURDATE();`

	rows, err := db.Query(sqlString, PbillerDashBoardReqRec.Login_id)
	if err != nil {
		log.Println(err)
		pbillerDashBoardRespRec.Status = "E"
		pbillerDashBoardRespRec.ErrMsg = "Error: " + err.Error()
		return pbillerDashBoardRespRec, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&tsal); err != nil {
			log.Println(err)
			pbillerDashBoardRespRec.Status = "E"
			pbillerDashBoardRespRec.ErrMsg = "Error: " + err.Error()
			return pbillerDashBoardRespRec, err
		}
	}
	sqlString = `SELECT COALESCE(SUM(mbm.billAmount), 0)
		          FROM medapp_bill_master mbm
		          WHERE billDate = CAST(CURDATE() - INTERVAL 1 DAY AS DATE) AND login_id = ?`
	rows1, err1 := db.Query(sqlString, PbillerDashBoardReqRec.Login_id)
	if err1 != nil {
		log.Println(err1)
		pbillerDashBoardRespRec.Status = "E"
		pbillerDashBoardRespRec.ErrMsg = "Error: " + err1.Error()
		return pbillerDashBoardRespRec, err1
	}
	defer rows1.Close()

	for rows1.Next() {
		if err := rows1.Scan(&ysal); err != nil {
			log.Println(err)
			pbillerDashBoardRespRec.Status = "E"
			pbillerDashBoardRespRec.ErrMsg = "Error: " + err.Error()
			return pbillerDashBoardRespRec, err
		}
	}
	pbillerDashBoardRespRec.TodaySales = tsal
	pbillerDashBoardRespRec.YesterdaySales = ysal
	pbillerDashBoardRespRec.Status = "S"
	log.Println("BillerDashBoardMethod(-)")
	return pbillerDashBoardRespRec, nil
}
