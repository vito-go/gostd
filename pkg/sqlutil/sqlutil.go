// Package sqlutil  存放一些sql语句通用的工具类函数.
package sqlutil

import (
	"database/sql"
	"encoding/json"
)

// RowsToJsonB 将sql查询结果rows转换为[]byte.
// 可以将结果unmarshal到表结构体.
func RowsToJsonB(rows *sql.Rows) ([]byte, error) {
	resultMap, err := RowsToMap(rows)
	if err != nil {
		return nil, err
	}
	return json.Marshal(resultMap)
}

// RowsToMap 将sql查询结果rows转换为map.
func RowsToMap(rows *sql.Rows) (map[string]interface{}, error) {
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	scans := make([]interface{}, len(cols))
	for n := range scans {
		scans[n] = new(interface{})
	}
	err = rows.Scan(scans...)
	if err != nil {
		return nil, err
	}
	var resultMap = make(map[string]interface{})
	for n, col := range cols {
		resultMap[col] = *scans[n].(*interface{})
	}
	return resultMap, nil
}
