package dbinfo

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/pkg/errors"
)

func GetStep4MySQL(db *sql.DB, query string, conc int) (int64, error) {
	explain := "explain " + query

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

	if len(scanArgs) != 12 {
		return 0, errors.New("not mysql v8.x")
	}

	for rows.Next() {
		if err := rows.Scan(scanArgs...); err != nil {
			return 0, errors.WithStack(err)
		}
		// only scan first row
		break
	}

	// for mysql 8.x
	parseInt, err := strconv.ParseInt(string(scanVals[9]), 10, 64)
	if err != nil {
		return 0, errors.WithStack(err)
	}

	step := parseInt/int64(conc) + 1
	log.Printf("explain want to result size, rows:%s=>int:%d, step:%d", scanVals[9], parseInt, step)

	return step, nil
}
