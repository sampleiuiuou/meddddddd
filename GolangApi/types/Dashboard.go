package types

// get Today Sales Report Response Structure
type TodaySalesRespStruct struct {
	TodaySales float64 `json:"today_sales"`
	ErrMsg     string  `json:"errmsg"`
	Status     string  `json:"status"`
	Msg        bool    `json:"msg"`
}

// get Today sales and inventry in response Structure
type InventryRespStruct struct {
	Inventry   float64 `json:"inventry"`
	TodaySales float64 `json:"today_sales"`
	ErrMsg     string  `json:"errmsg"`
	Status     string  `json:"status"`
	Msg        bool    `json:"msg"`
}
