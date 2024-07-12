package billentry

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"medapp/database"
	"net/http"
)

type BillStockMedicineStruct struct {
	Medicine_master_id int     `json:"medicine_id"`
	Medicine_name      string  `json:"medicinename"`
	Quantity           int     `json:"quantity"`
	Unit_price         float64 `json:"unitprice"`
	Amount             float64 `json:"amount"`
	Brand              string  `json:"brand"`
}

type BillStockReqStruct struct {
	User_id             string                    `json:"userid"`
	Bill_no             int                       `json:"billno"`
	Bill_date           string                    `json:"billdate"`
	Bill_amount         int                       `json:"billamount"`
	Net_price           float64                   `json:"netprice"`
	Bill_gst            float64                   `json:"gst"`
	Login_id            int                       `json:"login_id"`
	BillMedicineDetails []BillStockMedicineStruct `json:"billmeddetails"`
}

type CommonSaveRespStruct struct {
	ErrMsg string `json:"errmsg"`
	Status string `json:"status"`
	Msg    int    `json:"msg"`
}

func SaveBill(w http.ResponseWriter, r *http.Request) {
	log.Println("Update Stock started(+)")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "NAME,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")

	if r.Method == "PUT" {
		var lBillStockRec BillStockReqStruct
		var lDetailsrespRec CommonSaveRespStruct
		body, err := ioutil.ReadAll(r.Body)
		log.Println(string(body))
		if err != nil {
			log.Println("Login verify", err)
			lDetailsrespRec.ErrMsg = err.Error()
			lDetailsrespRec.Status = "error"

		} else {

			err = json.Unmarshal(body, &lBillStockRec)

			if err != nil {
				log.Print("unmarshall verify err ", err)
				lDetailsrespRec.ErrMsg = err.Error()
				lDetailsrespRec.Status = "error"

			} else {

				Details, err := SaveBillEntryStock(lBillStockRec)
				log.Println(err)
				if Details.Status != "S" {
					lDetailsrespRec.ErrMsg = err.Error()
					lDetailsrespRec.Status = "error"

				} else {
					lDetailsrespRec = Details
					lDetailsrespRec.Status = "S"
				}
			}
		}
		data, err := json.Marshal(lDetailsrespRec)
		if err != nil {
			log.Println(err)
		} else {
			fmt.Fprintln(w, string(data))
		}
		log.Println("Update Stock Verification(-)")
	}
}

func SaveBillEntryStock(collect BillStockReqStruct) (CommonSaveRespStruct, error) {
	db, err := database.LocalDBConnect()

	var responsedetails CommonSaveRespStruct
	user_id := collect.User_id
	billno := collect.Bill_no
	billdate := collect.Bill_date
	billamount := collect.Bill_amount
	gst := collect.Bill_gst
	net := collect.Net_price
	loginid := collect.Login_id
	details := collect.BillMedicineDetails
	if err != nil {
		log.Println(err)
		responsedetails.ErrMsg = err.Error()
		responsedetails.Status = "E"

		return responsedetails, err
	} else {
		defer db.Close()
		sqlString := `insert into medapp_bill_master 
		(billNo,billDate,billAmount,billGst,netPrice,login_id,created_by,
		created_date,updated_by,updated_date) values (?,?,?,?,?,?,?,now(),?,now());`
		rows, err := db.Exec(sqlString, billno, billdate, billamount, gst, net, loginid, user_id, user_id)
		if err != nil {
			log.Println(err)
			responsedetails.ErrMsg = err.Error()
			responsedetails.Status = "E"
			aff, err := rows.RowsAffected()
			log.Println(aff, err)
			responsedetails.Msg = int(aff)
		} else {

			responsedetails.Status = "S"
			aff, err := rows.RowsAffected()
			log.Println(aff, err)
			responsedetails.Msg = int(aff)
		}
		updateString := `update medapp_stock 
		set quantity =quantity - ?,updated_by =?,updated_date =now()
		where medicine_master_id = (select medicine_master_id 
		from medapp_medicine_master mmm 
		where medicineName =? and brand=?);`
		for _, val := range details {
			row, err := db.Exec(updateString, val.Quantity, user_id, val.Medicine_name, val.Brand)
			if err != nil {
				log.Println(err)
				responsedetails.ErrMsg = "Error" + err.Error()
				responsedetails.Status = "E"
				aff, err := row.RowsAffected()
				log.Println(aff, err)
				responsedetails.Msg = int(aff)
			} else {
				responsedetails.Status = "S"
				aff, err := row.RowsAffected()
				log.Println(aff, err)
				responsedetails.Msg = int(aff)
			}
		}
		updateString2 := `insert into medapp_bill_details
		(billNo,medicine_master_id,quantity,unitPrice,amount,created_by,created_date,
		updated_by,updated_date)values(?,?,?,?,?,?,now(),?,now());`
		for _, val1 := range details {
			row, err := db.Exec(updateString2, billno, val1.Medicine_master_id, val1.Quantity,
				val1.Unit_price, val1.Amount, user_id, user_id)
			if err != nil {
				log.Println(err)
				responsedetails.Status = "E"
				aff, err := row.RowsAffected()
				log.Println(aff, err)
				responsedetails.Msg = int(aff)
			} else {
				responsedetails.Status = "S"
				aff, err := row.RowsAffected()
				log.Println(aff, err)
				responsedetails.Msg = int(aff)
			}
		}
	}

	return responsedetails, err
}
