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
func SaveBillMasterAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")
	w.Header().Set("Access-Control-Allow-headers", "USER,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
	log.Println("saveBill master(+)")
	if r.Method == "POST" {
		var lbm types.BillMasterReqStruct

		var resp types.CommonRespStruct
		body, err := io.ReadAll(r.Body)
		resp.Status = "S"
		if err != nil {
			log.Println(err)
			resp.Status = "E"
			resp.ErrMsg = "Error" + err.Error()
		} else {
			err := json.Unmarshal(body, &lbm)
			if err != nil {
				log.Println(err)
				resp.Status = "E"
				resp.ErrMsg = "Error" + err.Error()
			} else {
				resp = method.SaveBillMaster(lbm)

			}
		}
		data, err := json.Marshal(resp)
		if err != nil {
			fmt.Fprintf(w, "Error taking data"+err.Error())
		} else {
			w.Write(data)
		}
		log.Println("saveBill master(+)")
	}
}
