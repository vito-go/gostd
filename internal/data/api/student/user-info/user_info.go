package userinfo

import (
	"context"
	"encoding/json"
	"fmt"

	"gitea.com/liushihao/gostd/internal/data/database/studentdb"
)

type Cli struct {
	dao *studentdb.Dao
}

func (c *Cli) Hello() string {
	return "hello world"
}

func NewCli(dao *studentdb.Dao) *Cli {
	return &Cli{dao: dao}
}

func (c *Cli) GetUserInfoMapByID(id int64) (map[string]string, error) {
	m, err := c.dao.UserInfoRepo.GetInfoById(context.Background(), id)
	if err != nil {
		return nil, err
	}
	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	var resultMap map[string]interface{}
	err = json.Unmarshal(b, &resultMap)
	if err != nil {
		return nil, err
	}
	var result = make(map[string]string, len(resultMap))
	for k, v := range resultMap {
		result[k] = fmt.Sprint(v)
	}
	return result, err

}
func (c *Cli) GetNameById(id int64) (string, error) {
	m, err := c.dao.UserInfoRepo.GetInfoById(context.Background(), id)
	if err != nil {
		return "", err
	}
	return m.Name, nil
}
