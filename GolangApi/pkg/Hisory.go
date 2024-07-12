package method

import (
	"database/sql"
	"log"
	"medapp/database"
	"medapp/types"
)

// insert the Login Hisotry
func InsertLogin(db *sql.DB, loginId int, userId string) (int, bool) {
	var resp types.LoginRespStruct
	defer db.Close()
	insertQuery2 := `INSERT into login_history (login_id,Entry,Date,Time,created_by,created_date,updated_by,updated_date) VALUES (?,?,now(),now(),?,now(),?,now());`
	result, err := db.Exec(insertQuery2, loginId, "Login", userId, userId)
	if err != nil {
		log.Println("IL01", err)
		resp.ErrMsg = err.Error()
		resp.Status = "E"
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("IL02", err)
		resp.ErrMsg = err.Error()
		resp.Status = "E"
	}

	if rowsAffected > 0 {
		lh_id, err := result.LastInsertId()
		if err != nil {
			log.Println("IL03", err)
			resp.ErrMsg = err.Error()
			resp.Status = "E"
		}
		return int(lh_id), true
	} else {
		return 0, false
	}
}
func InsertLogout(loginId int, userId string) types.CommonRespStruct {
	var resp types.CommonRespStruct
	db, err := database.LocalDBConnect()
	resp.Status = "S"
	resp.Msg = true
	if err != nil {
		log.Println("IL04", err)
		resp.ErrMsg = err.Error()
		resp.Status = "E"
		resp.Msg = false
		return resp
	}
	defer db.Close()
	insertQuery2 := `INSERT into login_history (login_id,Entry,Date,Time,created_by,created_date,updated_by,updated_date) VALUES (?,?,now(),now(),?,now(),?,now());`
	result, err := db.Exec(insertQuery2, loginId, "Logout", userId, userId)
	if err != nil {
		log.Println("IL05", err)
		resp.ErrMsg = err.Error()
		resp.Status = "E"
		resp.Msg = false
		return resp
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("IL06", err)
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
func GetAllHistory(resp *types.ListHistoryRespStruct) []types.LoginHistory {
	var lhistory types.LoginHistory
	var lhArr []types.LoginHistory
	db, err := database.LocalDBConnect()
	resp.Status = "S"
	resp.Msg = true
	if err != nil {
		log.Println("SR01", err)
		resp.ErrMsg = err.Error()
		resp.Status = "E"
		resp.Msg = false
		return lhArr
	}
	defer db.Close()
	listHistory := `SELECT NVL(lh.Entry,""),NVL( lh.Date,""), NVL(lh.Time,""), NVL(l.UserId,"")
	FROM login_history lh JOIN login l ON lh.login_id = l.login_id;`
	rows, err := db.Query(listHistory)
	if err != nil {
		log.Println("SR02", err)
		resp.ErrMsg = err.Error()
		resp.Status = "E"
		resp.Msg = false
		return lhArr

	} else {
		for rows.Next() {
			err := rows.Scan(&lhistory.Entry, &lhistory.Date, &lhistory.Time, &lhistory.UserId)
			if err != nil {
				log.Println("SR03", err)
				resp.ErrMsg = err.Error()
				resp.Status = "E"
				resp.Msg = false
				return lhArr
			} else {
				lhArr = append(lhArr, lhistory)
			}
		}

		rows.Close()
		return lhArr
	}

}
