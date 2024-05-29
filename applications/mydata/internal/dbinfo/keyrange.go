package dbinfo

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
	"mydata/internal/parser"
)

type RangeVals struct {
	Start string
	End   string
}

func GetKeyRange(keyVals []string) ([]RangeVals, error) {
	var lenKeyVals = len(keyVals)
	if lenKeyVals < 2 {
		return nil, errors.New("key values only small")
	}

	tmpRange := make([]RangeVals, lenKeyVals-1)
	var start = 0
	var end = 1
	for end < lenKeyVals {
		tmpRange[start].Start = keyVals[start]
		tmpRange[start].End = keyVals[end]
		start++
		end++
	}

	return tmpRange, nil
}

func GetChunkQuery(dbc DbConf, queryOrig string, concNum int) ([]string, error) {
	if concNum <= 1 {
		log.Printf("concurrency num set too small:%d", concNum)
		return []string{queryOrig}, nil
	}

	conn, err := NewDB(dbc)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = conn.Close()
	}()

	stmt, err := parser.NewSelectStmt(queryOrig)
	if err != nil {
		return nil, err
	}

	tableOrig, isOneTable := parser.ExtractTable(stmt)
	if isOneTable == false {
		return nil, err
	}

	if isSimpleSql := parser.JudgeSimpleSQL(stmt); !isSimpleSql {
		return nil, errors.Errorf("not one simple query sql")
	}

	schema, table, err := parser.GetTabAndSchema(tableOrig, dbc.Dbname)
	if err != nil {
		return nil, err
	}

	keyNameStr, err := GetKeyName(conn, schema, table)
	if err != nil {
		return nil, err
	}

	// 根据explain和并发读计算step大小
	var step int64
	if IsOB(conn) {
		step, err = GetStep4OB(conn, queryOrig, concNum)
		if err != nil {
			return nil, err
		}
	} else {
		step, err = GetStep4MySQL(conn, queryOrig, concNum)
		if err != nil {
			return nil, err
		}
	}

	keyValues, err := GetKeyValueByStep(conn, parser.GetOrigWhere(stmt), schema, table, keyNameStr, step)
	if err != nil {
		return nil, err
	}

	keyRanges, err := GetKeyRange(keyValues)
	if err != nil {
		return nil, err
	}

	rewriterSql := make([]string, 0)
	maxChunk := len(keyRanges)
	for i, vv := range keyRanges {
		subWhere := ""
		if i == maxChunk-1 {
			subWhere = fmt.Sprintf("(%s)>=%s", keyNameStr, vv.Start)
		} else {
			subWhere = fmt.Sprintf("(%s)>=%s and (%s)<%s", keyNameStr, vv.Start, keyNameStr, vv.End)
		}

		newSql := parser.AddSubWhereSql(stmt, subWhere)
		rewriterSql = append(rewriterSql, newSql)
	}

	return rewriterSql, nil
}
