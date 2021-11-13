package studentdao

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"

	"github.com/liushihao/gostd/internal/data/dberr"
	"github.com/liushihao/gostd/pkg/sqlutil"
)

const UserInfoTableName = "user_info"

type userInfoRepo struct {
	db *studentDB
}

type UserInfoModel struct {
	ID       int64  `json:"id,omitempty"`
	Number   int64  `json:"number,omitempty"`
	Name     string `json:"name,omitempty"`
	Province string `json:"province,omitempty"`
	City     string `json:"city,omitempty"`
}

func NewUserInfoRepo(db *studentDB) *userInfoRepo {
	return &userInfoRepo{db: db}
}

func (u *userInfoRepo) GetInfoByID(ctx context.Context, id int64) (*UserInfoModel, error) {
	rows, err := u.db.DB.QueryContext(ctx, fmt.Sprintf("SELECT * FROM %s where id=%d;", UserInfoTableName, id))
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		return rowsToModel(rows)
	}
	return nil, dberr.ErrNotFound

}

func rowsToModel(rows *sql.Rows) (*UserInfoModel, error) {
	b, err := sqlutil.RowsToJsonB(rows)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var m UserInfoModel
	err = json.Unmarshal(b, &m)
	if err != nil {
		return nil, err
	}
	return &m, err
}
