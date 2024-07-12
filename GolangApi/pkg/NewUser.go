package method

import (
	"log"
	"medapp/database"
	"medapp/types"
)

func NewUser(newUser types.AddUserReqStruct) types.CommonRespStruct {
	var resp types.CommonRespStruct
	db, err := database.LocalDBConnect()
	resp.Status = "S"
	resp.Msg = true
	if err != nil {
		log.Println("NS01", err)
		resp.ErrMsg = err.Error()
		resp.Status = "E"
		resp.Msg = false
		return resp
	}
	defer db.Close()
	insertQuery2 := `INSERT INTO login
	(UserId, Password, Role, created_by, created_date, updated_by, updated_date) SELECT ?, ?, ?, ?, NOW(), ?, NOW() WHERE NOT EXISTS ( SELECT login_id FROM login ml WHERE ml.UserId = ?)`
	result, err := db.Exec(insertQuery2, newUser.UserId, newUser.Password, newUser.Role, newUser.CreaterUserId, newUser.CreaterUserId, newUser.UserId)
	if err != nil {
		log.Println("NS02", err)
		resp.ErrMsg = err.Error()
		resp.Status = "E"
		resp.Msg = false
		return resp
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("NS03", err)
		resp.ErrMsg = err.Error()
		resp.Status = "E"
		resp.Msg = false
		return resp
	}
	if rowsAffected > 0 {
		return resp
	} else {
		resp.Status = "E"
		resp.Msg = false
		return resp
	}
}
