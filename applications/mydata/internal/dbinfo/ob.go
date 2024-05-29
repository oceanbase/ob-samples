package dbinfo

import (
	"database/sql"
	"encoding/json"
	"log"
	"strings"

	"github.com/pkg/errors"
)

func IsOB(conn *sql.DB) bool {
	rows, err := conn.Query("show variables like 'version'")
	if err != nil {
		log.Printf("err:%v", err)
		return false
	}
	defer rows.Close()

	scanArgs, scanVals, err := MakeScanBuf(rows)
	if err != nil {
		return false
	}

	for rows.Next() {
		if err := rows.Scan(scanArgs...); err != nil {
			log.Printf("err:%v", err)
			return false
		}
		// only scan first row
		break
	}

	return strings.Contains(string(scanVals[1]), "OceanBase")
}

func GetStep4OB(db *sql.DB, query string, conc int) (int64, error) {
	explain := "explain format=json " + query

	log.Printf("get step by conc:[%s]", explain)
	rows, err := db.Query(explain)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	defer rows.Close()

	scanArgs, scanVals, err := MakeScanBuf(rows)
	if err != nil {
		return 0, err
	}

	for rows.Next() {
		if err := rows.Scan(scanArgs...); err != nil {
			return 0, errors.WithStack(err)
		}
		// only scan first row
		break
	}

	var obExplainRst struct {
		Rows int64 `json:"EST.ROWS"`
	}
	if err := json.Unmarshal(scanVals[0], &obExplainRst); err != nil {
		return 0, errors.WithStack(err)
	}

	step := obExplainRst.Rows/int64(conc) + 1
	log.Printf("explain want to result size, rows:%d, step:%d", obExplainRst.Rows, step)

	return step, nil
}
