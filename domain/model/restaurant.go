package model

import "github.com/oklog/ulid/v2"


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

type CheckHistory struct {
	ReceiptID ulid.ULID `json:"receipt_id"`
}