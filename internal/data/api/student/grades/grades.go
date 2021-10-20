package grades

import "gitea.com/liushihao/gostd/internal/data/database"

type API struct {
	db *database.DB
}

func (a *API) GetGradesByNameAndId(id int64, name string) {
	panic("implement me")
}

func (a *API) GetTotalGradesByID(id int64) int64 {
	return 99
}

func NewApi(db *database.DB) *API {
	return &API{db: db}
}
