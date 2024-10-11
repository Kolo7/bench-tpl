package db

type Table struct {
	Name    string
	Columns []*Column
	Rows    []Row
}

type Column struct {
	Field   string `json:"field" db:"field"`
	Type    string `json:"type" db:"type"`
	Null    string `json:"null" db:"null"`
	Key     string `json:"key" db:"key"`
	Default any    `json:"default" db:"default"`
	Extra   string `json:"extra" db:"extra"`
	GoType  string `json:"go_type" db:"go_type"`
	Lower   string `json:"lower" db:"lower"`
	Upper   string `json:"upper" db:"upper"`
}

type Row []*Ele

type Ele struct {
	Name   string
	Val    any
	Column *Column
}
