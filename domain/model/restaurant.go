package model

import "time"

type InitialTable struct {
	TableNo uint8 `json:"table_no"`
}

type TableDetail struct {
	TableNo uint8 `json:"table_no"`
	Status  string `json:"status"`
	Time time.Time `json:"time"`
}

type GiveTable struct {
	TableNo uint8 `json:"table_no"`
}