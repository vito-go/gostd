package info

import "gitea.com/liushihao/gostd/internal/data/database"

type API struct {
	db *database.TeacherDB
}

func (A *API) GetInfoByID(id int64) {
	panic("implement me")
}

func NewTable(db *database.TeacherDB) *API {
	return &API{db: db}
}
