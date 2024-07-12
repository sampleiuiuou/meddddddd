package loginhistory

import (
	"encoding/json"
	"fmt"
	"log"
	"medapp/database"
	"net/http"
)

type LoginHistoryStruct struct {
	Userid string `json:"userid"`
	Entry  string `json:"entry"`
	Time   string `json:"time"`
	Date   string `json:"date"`
}

type LoginHistoryRespStruct struct {
	LoginHistoryArr []LoginHistoryStruct `json:"loginhistory"`
	ErrMsg          string               `json:"errmsg"`
	Status          string               `json:"status"`
	Msg             string               `json:"msg"`
}

func LoginHistory(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
	log.Println("LoginHistory(+)")

	if r.Method == "PUT" {
		// var lLoginHistoryRec LoginHistoryStruct
		var lLoginHistoryRespRec LoginHistoryRespStruct

		lLoginHistoryRespRec, err := LoginHistoryFetch()
		if err != nil {
			log.Println(err)
			lLoginHistoryRespRec.Status = "E"
			lLoginHistoryRespRec.ErrMsg = "Error : " + err.Error()
		}
		data, err := json.Marshal(lLoginHistoryRespRec)
		if err != nil {
			fmt.Fprintf(w, "Error taking data"+err.Error())
		} else {
			fmt.Fprint(w, string(data))
			fmt.Printf("%s", data)
		}
		log.Println("LoginHistory(-)")
	}

}
func LoginHistoryFetch() (LoginHistoryRespStruct, error) {
	log.Println("LoginHistoryFetch(+)")
	var lLoginHistoryRespRec LoginHistoryRespStruct
	var lLoginHistoryRec LoginHistoryStruct
	db, err := database.LocalDBConnect()
	if err != nil {
		log.Println(err)
		return lLoginHistoryRespRec, err
	} else {
		defer db.Close()
		sqlString := `SELECT ml.userid, 
		NVL(mlh.entry, 'working') AS entry, 
		NVL(mlh.time, 'Working') AS time, 
		NVL(mlh.date, 'NULL') AS date
 FROM medapp_login ml
 Right JOIN medapp_login_history mlh ON ml.login_id = mlh.login_id;
 
 `
		rows, err := db.Query(sqlString)
		if err != nil {
			log.Println(err)
			return lLoginHistoryRespRec, err
		} else {
			for rows.Next() {
				err := rows.Scan(&lLoginHistoryRec.Userid, &lLoginHistoryRec.Entry, &lLoginHistoryRec.Time, &lLoginHistoryRec.Date)
				if err != nil {
					log.Println(err)
					return lLoginHistoryRespRec, err
				} else {
					lLoginHistoryRespRec.LoginHistoryArr = append(lLoginHistoryRespRec.LoginHistoryArr, lLoginHistoryRec)
					lLoginHistoryRespRec.Status = "S"
					lLoginHistoryRespRec.Msg = "Success"
				}
			}
			log.Println("LoginHistoryFetch(-)")
			return lLoginHistoryRespRec, err
		}
	}
}
