package class

import (
	"gitea.com/liushihao/gostd/internal/data/database"
)

type Table struct {
	db *database.StudentDB
}

func (A Table) GetNameByID(id int64) string {
	panic("implement me")
}

func NewAPI(db *database.StudentDB) *Table {
	return &Table{db: db}
}
