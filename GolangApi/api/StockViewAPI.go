package api

import (
	"encoding/json"
	"fmt"
	"log"
	method "medapp/pkg"
	"medapp/types"
	"net/http"
)

// Find the Sales in Particular Date API
func StockViewAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET,OPTIONS")
	w.Header().Set("Access-Control-Allow-headers", "USER,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
	log.Println("stock insert(+)")
	if r.Method == "GET" {
		var resp types.StockViewRespStruct
		stockArr := method.GetAllStock(&resp)
		resp.StockArr = stockArr
		data, err := json.Marshal(resp)
		if err != nil {
			fmt.Fprintf(w, "Error taking data"+err.Error())
		} else {
			w.Write(data)
		}
		log.Println("stock insert(-)")
	}
}
