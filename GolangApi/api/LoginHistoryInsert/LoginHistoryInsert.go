package logininsertpkg

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"medapp/database"
	"net/http"
)

type UserStruct struct {
	LoginId int    `json:"login_id"`
	UserId  string `json:"userid"`
}
type LoginRespStruct struct {
	ErrMsg string `json:"errmsg"`
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

func Logininsert(w http.ResponseWriter, r *http.Request) {
	log.Println("SaveBill(+)")
	{
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
		w.Header().Set("Access-Control-Allow-headers", "USER,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
		log.Println("LoginHistory(+)")
		if r.Method == "PUT" {
			var Userrec UserStruct
			var resp LoginRespStruct
			body, err := io.ReadAll(r.Body)
			log.Println(string(body))
			resp.Status = "S"
			if err != nil {
				log.Println(err)
				resp.Status = "E"
				resp.ErrMsg = "Error" + err.Error()
			} else {
				err := json.Unmarshal(body, &Userrec)
				log.Println(Userrec)
				if err != nil {
					log.Println(err)
					resp.Status = "E"
					resp.ErrMsg = "Error" + err.Error()
				} else {
					resp, err = InsertLogin(Userrec)
					if err != nil {
						log.Println(err)
						resp.Status = "E"
						resp.ErrMsg = "Error" + err.Error()
					} else {
						resp.Status = "S"
						resp.Msg = "Success"
					}
				}
			}
			data, err := json.Marshal(resp)
			if err != nil {
				fmt.Fprintf(w, "Error taking data"+err.Error())
			} else {
				w.Write(data)
			}
			log.Println("LoginHistory insert(-)")
		}
	}
}

func InsertLogin(pUserrec UserStruct) (LoginRespStruct, error) {
	log.Println("insert")
	log.Println(pUserrec)
	var resp LoginRespStruct
	db, err := database.LocalDBConnect()
	resp.Status = "S"
	resp.Msg = "Success"
	if err != nil {
		log.Println("SB01", err)
		resp.ErrMsg = err.Error()
		resp.Status = "E"
		resp.Msg = "Error"
		return resp, err
	}
	defer db.Close()
	insertQuery2 := `INSERT into medapp_login_history (login_id,Entry,Date,Time,created_by,created_date,updated_by,updated_date)
					 VALUES (?,?,now(),now(),?,now(),?,now());`
	result, err := db.Exec(insertQuery2, pUserrec.LoginId, "Login", pUserrec.UserId, pUserrec.UserId)
	if err != nil {
		log.Println("IL01", err)
		resp.ErrMsg = err.Error()
		resp.Status = "E"
		return resp, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("IL02", err)
		resp.ErrMsg = err.Error()
		resp.Status = "E"
		return resp, err
	}

	if rowsAffected > 0 {
		log.Println("Inserted Successfully")
		return resp, err
	} else {
		resp.Status = "E"
		resp.Msg = "Error"
		return resp, err
	}
}
