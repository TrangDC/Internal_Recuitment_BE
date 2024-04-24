package models

type AuditTrailData struct {
	Module string        `json:"module"`
	Create []interface{} `json:"create"`
	Update []interface{} `json:"update"`
	Delete []interface{} `json:"delete"`
}

type AuditTrailUpdate struct {
	Field string      `json:"field"`
	Value ValueChange `json:"value"`
}

type ValueChange struct {
	OldValue interface{} `json:"oldValue"`
	NewValue interface{} `json:"newValue"`
}

type AuditTrailCreateDelete struct {
	Field string      `json:"field"`
	Value interface{} `json:"value"`
}
