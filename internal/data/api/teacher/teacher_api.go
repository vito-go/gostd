package teacher

import (
	"gitea.com/liushihao/gostd/internal/data/api/teacher/info"
)

type API struct {
	infoIface info.Interface
}

func NewApi(api *info.API) *API {
	return &API{infoIface: api}
}
