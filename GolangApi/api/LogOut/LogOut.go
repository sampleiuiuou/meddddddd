package logout

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"medapp/database"
	"net/http"
)

type LogoutReqStruct struct {
	UserId  string `json:"userid"`
	LoginId int    `json:"login_id"`
}
type CommonRespStruct struct {
	ErrMsg string `json:"errmsg"`
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

func Logout(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
	log.Println("Logout(+)")

	var lLogoutReqRec LogoutReqStruct
	var lCommonRespRec CommonRespStruct

	body, err := io.ReadAll(r.Body)
	log.Println(string(body))
	if err != nil {
		log.Println("Error : " + err.Error())
		lCommonRespRec.Status = "E"
		lCommonRespRec.ErrMsg = " Error :" + err.Error()
	} else {
		err = json.Unmarshal(body, &lLogoutReqRec)
		if err != nil {
			log.Println("Error :" + err.Error())
			lCommonRespRec.Status = "E"
			lCommonRespRec.ErrMsg = "Error :" + err.Error()
		} else {
			lCommonRespRec, err = LogOutInsert(lLogoutReqRec)
			if err != nil {
				log.Println("Error :" + err.Error())
				lCommonRespRec.Status = "E"
				lCommonRespRec.ErrMsg = "Error :" + err.Error()
			} else {
				lCommonRespRec.Msg = "Success"
				lCommonRespRec.Status = "S"
			}
		}
	}
	data, err := json.Marshal(lCommonRespRec)
	if err != nil {
		fmt.Fprintf(w, "Error taking data"+err.Error())
	} else {
		fmt.Fprint(w, string(data))
		fmt.Printf("%s", data)
	}
	log.Println("Logout(-)")
}

func LogOutInsert(lLogoutReqRec LogoutReqStruct) (CommonRespStruct, error) {
	log.Println("LogOutInsert(+)")
	var lCommonRespRec CommonRespStruct
	db, err := database.LocalDBConnect()
	if err != nil {
		log.Println(err)
		return lCommonRespRec, err
	} else {
		defer db.Close()
		sqlString := `INSERT into medapp_login_history (login_id,Entry,Date,Time,created_by,created_date,updated_by,updated_date)
		VALUES (?,?,now(),now(),?,now(),?,now());`
		rows, err := db.Exec(sqlString, lLogoutReqRec.LoginId, "Logout", lLogoutReqRec.UserId, lLogoutReqRec.UserId)
		if err != nil {
			lCommonRespRec.Status = "E"
			log.Println(err)
			lCommonRespRec.ErrMsg = "Error :" + err.Error()
		} else {
			affected, err := rows.RowsAffected()
			if err != nil {
				lCommonRespRec.Status = "E"
				log.Println(err)
				lCommonRespRec.ErrMsg = "Error :" + err.Error()
			} else {
				if affected == 0 {
					lCommonRespRec.Status = "E"
					log.Println(err)
					lCommonRespRec.ErrMsg = "Error No Rows Inserted"
				} else {
					log.Println("Inserted Successfully")
					lCommonRespRec.Status = "S"
					lCommonRespRec.Msg = "Success"
					log.Println("logoutInsert(-)")
				}
			}
		}
		return lCommonRespRec, err

	}

}
