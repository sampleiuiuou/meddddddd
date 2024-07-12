package method

import (
	"log"
	"medapp/database"
	"medapp/types"
)

func LoginCheckValidation(userId, password string) (types.UserStruct, types.LoginRespStruct) {
	var user types.UserStruct
	var resp types.LoginRespStruct
	resp.Status = "S"
	resp.Msg = true
	db, err := database.LocalDBConnect()
	if err != nil {
		log.Println("LV01", err)
		resp.ErrMsg = err.Error()
		resp.Status = "E"
		resp.Msg = false
		return user, resp
	} else {
		defer db.Close()
		resp.Status = "S"
		SelectQuery := `SELECT login_id, NVL(Role,""), NVL(UserId,"") FROM login WHERE UserId=? and Password =?`
		rows, err := db.Query(SelectQuery, userId, password)
		if err != nil {
			log.Println("LV02", err)
			resp.ErrMsg = err.Error()
			resp.Status = "E"
			resp.Msg = false
			return user, resp

		} else {
			if !rows.Next() {
				resp.Status = "E"
				resp.ErrMsg = "user not Found"
				resp.Msg = false
				return user, resp
			} else {
				err := rows.Scan(&user.LoginId, &user.Role, &user.UserId)
				if err != nil {
					log.Println("LV03", err)
					resp.ErrMsg = err.Error()
					resp.Status = "E"
					resp.Msg = false
					return user, resp

				} else {
					lh_id, msg := InsertLogin(db, user.LoginId, user.UserId)
					if msg {
						user.LHId = lh_id
						resp.Msg = msg
						return user, resp

					} else {
						resp.Msg = msg
						return user, resp
					}
				}
			}

		}

	}
}
