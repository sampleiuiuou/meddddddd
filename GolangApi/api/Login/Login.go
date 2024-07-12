package login

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"medapp/database"
	"net/http"
)

type LoginReqStruct struct {
	UserId   string `json:"userid"`
	PassWord string `json:"password"`
}

type UserStruct struct {
	LoginId        int    `json:"login_id"`
	Role           string `json:"role"`
	UserId         string `json:"userid"`
	LoginHistoryId int    `json:"loginhistoryid"`
}

type LoginRespStruct struct {
	UserDetails UserStruct `json:"userdetails"`
	ErrMsg      string     `json:"errmsg"`
	Status      string     `json:"status"`
	Msg         string     `json:"msg"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
	log.Println("Login(+)")
	if r.Method == "PUT" {

		var lLoginDetailsRec LoginReqStruct
		var lUserDetailRec UserStruct
		var lLoginRespRec LoginRespStruct
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("Error : " + err.Error())
			lLoginRespRec.Status = "E"
			lLoginRespRec.ErrMsg = " Error :" + err.Error()
		} else {
			err = json.Unmarshal(body, &lLoginDetailsRec)
			if err != nil {
				log.Println("Error :" + err.Error())
				lLoginRespRec.Status = "E"
				lLoginRespRec.ErrMsg = "Error :" + err.Error()
			} else {
				lUserDetailRec, err = LoginCheckValidation(lLoginDetailsRec)
				if err != nil {
					log.Println("Error :" + err.Error())
					lLoginRespRec.Status = "E"
					lLoginRespRec.ErrMsg = "Error :" + err.Error()
				} else {
					lLoginRespRec.UserDetails = lUserDetailRec
					lLoginRespRec.Status = "S"
					lLoginRespRec.Msg = "SUCCESS"
				}
			}
		}
		data, err := json.Marshal(lLoginRespRec)
		if err != nil {
			fmt.Fprintf(w, "Error taking data"+err.Error())
		} else {

			fmt.Fprint(w, string(data))
			fmt.Printf("%s", data)
		}
		log.Println("Login(-)")
	}
}

func LoginCheckValidation(ploginRec LoginReqStruct) (UserStruct, error) {
	log.Println("LoginCheckValidation(+)")
	var lUserDetailRec UserStruct

	db, err := database.LocalDBConnect()
	if err != nil {
		log.Println(err)
		return lUserDetailRec, err
	} else {
		defer db.Close()
		sqlString := `select userid , role,login_id from medapp_login where userid=? AND password=?`
		rows, err := db.Query(sqlString, ploginRec.UserId, ploginRec.PassWord)
		if err != nil {
			log.Println(err)
			return lUserDetailRec, err
		} else {
			for rows.Next() {
				err := rows.Scan(&lUserDetailRec.UserId, &lUserDetailRec.Role, &lUserDetailRec.LoginId)
				if err != nil {
					log.Println(err)
					return lUserDetailRec, err
				}

			}
			log.Println("LoginCheckValidation(-)")
			return lUserDetailRec, err
		}
	}
}
