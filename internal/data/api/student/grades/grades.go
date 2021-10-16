package grades

import "local/gostd/internal/data/database"

type Api struct {
	db *database.Db
}

func (a *Api) GetGradesByNameAndId(id int64, name string) {
	panic("implement me")
}

func (a *Api) GetTotalGradesById(id int64) int64 {
	return 99
}

func NewApi(db *database.Db) *Api {
	return &Api{db: db}
}
