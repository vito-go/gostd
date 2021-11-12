package teacher

import (
	"gitea.com/liushihao/gostd/internal/data/api/teacher/info"
)

type API struct {
	infoCliAPI info.Interface
}

func NewApi(api *info.Cli) *API {
	return &API{infoCliAPI: api}
}
