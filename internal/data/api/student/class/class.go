package class

import "gitea.com/liushihao/gostd/internal/data/database"

type API struct {
	db *database.DB
}

func (A API) GetNameByID(id int64) string {
	panic("implement me")
}

func NewAPI(db *database.DB) *API {
	return &API{db: db}
}
