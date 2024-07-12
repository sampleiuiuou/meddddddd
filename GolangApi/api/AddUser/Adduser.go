package adduser

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"medapp/database"
	"net/http"
)

type AddUserReqStruct struct {
	UserId      string `json:"userid"`
	Password    string `json:"password"`
	Role        string `json:"role"`
	AdminUserId string `json:"adminuserid"`
}
type CommonRespStruct struct {
	ErrMsg string `json:"errmsg"`
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
	log.Println("AddUser(+)")
	if r.Method == "PUT" {
		var lAddUserReqRec AddUserReqStruct
		var lCommonRespRec CommonRespStruct

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("Error : " + err.Error())
			lCommonRespRec.Status = "E"
			lCommonRespRec.ErrMsg = " Error :" + err.Error()
		} else {
			err = json.Unmarshal(body, &lAddUserReqRec)
			log.Println(lAddUserReqRec)
			if err != nil {
				log.Println("Error :" + err.Error())
				lCommonRespRec.Status = "E"
				lCommonRespRec.ErrMsg = "Error :" + err.Error()
			} else {
				lCommonRespRec, err = AddUserMethod(lAddUserReqRec)
				if err != nil {
					log.Println("Error :" + err.Error())
					lCommonRespRec.Status = "E"
					lCommonRespRec.ErrMsg = "Error :" + err.Error()
				} else {
					lCommonRespRec.Status = "S"
					lCommonRespRec.Msg = "Success"
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
		log.Println("AddUser(-)")
	}
}

func AddUserMethod(pAddUserReqRec AddUserReqStruct) (CommonRespStruct, error) {
	log.Println("AddUserMethod(+)")
	var lCommonRespRec CommonRespStruct

	db, err := database.LocalDBConnect()
	if err != nil {
		log.Println(err)
		return lCommonRespRec, err
	} else {
		defer db.Close()
		sqlString := `Insert into medapp_login (userid,password,role,created_by,created_date) 
		select ?,?,?,?,now()
		where not exists (select login_id
		from medapp_login ma
		where ma.userid=?);`
		rows, err := db.Exec(sqlString, pAddUserReqRec.UserId, pAddUserReqRec.Password, pAddUserReqRec.Role, pAddUserReqRec.AdminUserId, pAddUserReqRec.UserId)
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
					lCommonRespRec.ErrMsg = "Error User Already Exists"
				} else {
					log.Println("Inserted Successfully")
					lCommonRespRec.Status = "S"
					lCommonRespRec.Msg = "Success"
					log.Println("AddUserMethod(-)")
				}
			}
		}
		return lCommonRespRec, err

	}

}
