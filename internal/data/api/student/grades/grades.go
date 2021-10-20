package grades

import (
	"gitea.com/liushihao/gostd/internal/data/database"
)

type Table struct {
	db *database.StudentDB
}

func (a *Table) GetGradesByNameAndId(id int64, name string) {
	panic("implement me")
}

func (a *Table) GetTotalGradesByID(id int64) int64 {
	return id * id
}

func NewTable(db *database.StudentDB) *Table {
	return &Table{db: db}
}
