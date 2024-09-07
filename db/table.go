package db

type Table struct {
	Name    string
	Columns []*Column
	Rows    [][]*Row
}

type Column struct {
	Field   string `json:"field" db:"field"`
	Type    string `json:"type" db:"type"`
	Null    bool   `json:"null" db:"null"`
	Key     string `json:"key" db:"key"`
	Default any    `json:"default" db:"default"`
	Extra   string `json:"extra" db:"extra"`
}

type Row struct {
	Name string
	Val  any
}
