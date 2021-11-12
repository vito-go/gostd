package studentdb

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"

	"gitea.com/liushihao/gostd/internal/data/dberr"
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
	b, err := rowsToJsonB(rows)
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
func rowsToJsonB(rows *sql.Rows) ([]byte, error) {
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	scans := make([]interface{}, len(cols))
	for n := range scans {
		scans[n] = new(interface{})
	}
	err = rows.Scan(scans...)
	var resultMap = make(map[string]interface{})
	if err != nil {
		return nil, err
	}
	for n, col := range cols {
		resultMap[col] = *scans[n].(*interface{})
	}
	return json.Marshal(resultMap)
}
