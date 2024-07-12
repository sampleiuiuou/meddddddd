package main

import (
	"log"
	"medapp/api"
	addmed "medapp/api/AddMed"
	adduser "medapp/api/AddUser"
	billdropdown "medapp/api/BillDropDown"
	billerdashboard "medapp/api/BillerDashBoard"
	logout "medapp/api/LogOut"
	login "medapp/api/Login"
	loginhistory "medapp/api/LoginHistory"
	logininsertpkg "medapp/api/LoginHistoryInsert"
	managerdashboard "medapp/api/ManagerDashBoard"
	salesreport "medapp/api/SalesReport"
	billentry "medapp/api/SaveBillDetails"
	stockentrydropdown "medapp/api/StockEntryDropDown"
	stockview "medapp/api/StockView"
	updatemed "medapp/api/UpdateMed"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// fmt.Println("Hello World")
	log.Println("Server Started")
	http.HandleFunc("/login", login.Login)                                        //1
	http.HandleFunc("/loginhistoryinsert", logininsertpkg.Logininsert)            //2
	http.HandleFunc("/adduser", adduser.AddUser)                                  //3
	http.HandleFunc("/addmed", addmed.AddMed)                                     //4
	http.HandleFunc("/updatemed", updatemed.UpdateMed)                            //5
	http.HandleFunc("/savebilldetails", billentry.SaveBill)                       //6
	http.HandleFunc("/salesreport", salesreport.SalesReport)                      //7
	http.HandleFunc("/loginhistory", loginhistory.LoginHistory)                   //8
	http.HandleFunc("/logout", logout.Logout)                                     //9
	http.HandleFunc("/stockview", stockview.StockView)                            //10
	http.HandleFunc("/billdropdown", billdropdown.BillDropDown)                   //11
	http.HandleFunc("/stockentrydropdown", stockentrydropdown.StockEntryDropDown) //12
	http.HandleFunc("/billerdashboard", billerdashboard.BillerDashBoard)          //13
	http.HandleFunc("/managerdashboard", managerdashboard.ManagerDashBoard)       //14
	http.HandleFunc("/view", api.StockViewAPI)
	http.HandleFunc("/savebill", api.SaveBillAPI)
	http.HandleFunc("/billmaster", api.SaveBillMasterAPI)
	http.ListenAndServe(":1094", nil)
	log.Println("Server Ended")
}
