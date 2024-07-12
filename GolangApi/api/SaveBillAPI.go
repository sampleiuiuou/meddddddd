package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	method "medapp/pkg"
	"medapp/types"
	"net/http"
)

// Add new Medicine in this Api
func SaveBillAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")
	w.Header().Set("Access-Control-Allow-headers", "USER,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
	log.Println("savebill")
	if r.Method == "POST" {
		var lbill types.BillDetailReqStruct
		var resp types.CommonRespStruct
		body, err := io.ReadAll(r.Body)
		resp.Status = "S"
		if err != nil {
			log.Println(err)
			resp.Status = "E"
			resp.ErrMsg = "Error" + err.Error()
		} else {
			err := json.Unmarshal(body, &lbill)
			if err != nil {
				log.Println(err)
				resp.Status = "E"
				resp.ErrMsg = "Error" + err.Error()
			} else {
				resp = method.SaveBill(lbill)

			}
		}
		data, err := json.Marshal(resp)
		if err != nil {
			fmt.Fprintf(w, "Error taking data"+err.Error())
		} else {
			w.Write(data)
		}
		log.Println("savebill(-)")
	}
}
