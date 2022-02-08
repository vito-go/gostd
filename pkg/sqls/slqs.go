package sqls

// sql函数，工具， 聚合日志 慢查询分析、告警
import (
	"context"
	"encoding/json"
	"time"

	"github.com/vito-go/mylog"

	"github.com/vito-go/gostd/pkg/database/sql"
)

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
	resultMap := make(map[string]interface{})
	if err != nil {
		return nil, err
	}
	for n, col := range cols {
		resultMap[col] = *scans[n].(*interface{})
	}
	return resultMap, nil
}

func RowsToJsonB(rows *sql.Rows) ([]byte, error) {
	resultMap, err := RowsToMap(rows)
	if err != nil {
		return nil, err
	}
	return json.Marshal(resultMap)
}

type DBQuery interface {
	Db() *sql.DB
}

// slowQuery 慢查询告警.
const slowQuery = time.Millisecond * 300

// QueryRowContext 日志记录、慢查询分析。
func QueryRowContext(ctx context.Context, d DBQuery, query string, dst interface{}) error {
	st := time.Now()
	row := d.Db().QueryRowContext(ctx, query)
	if err := row.Err(); err != nil {
		return err
	}
	err := row.ScanToStruct(dst)
	if err != nil {
		mylog.Ctx(ctx).WithField("sql", query).Error(err.Error())
		return err
	}
	timeElapsed := time.Since(st)
	if timeElapsed > slowQuery {
		mylog.Ctx(ctx).WithFields("timeElapsed", time.Since(st).String(), "result", dst).Warn("SLOW SQL-->", query)
		return nil
	}
	mylog.Ctx(ctx).WithFields("timeElapsed", time.Since(st).String(), "result", dst).Info("SQL==>", query)
	return nil
}
