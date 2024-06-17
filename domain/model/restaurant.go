package model

type InitialTable struct {
	TableNo uint8 `json:"table_no"`
}

type TableDetail struct {
	TableNo uint8 `json:"table_no"`
	Status  string `json:"status"`
}

type GiveTable struct {
	TableNo uint8 `json:"table_no"`
}