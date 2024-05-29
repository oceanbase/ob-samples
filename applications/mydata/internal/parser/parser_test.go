package parser

import (
	"fmt"
	"log"
	"testing"

	"mydata/internal/dbinfo"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func TestJudgeSampleSQL(t *testing.T) {
	cases := []struct {
		in       string
		isSimple bool
	}{
		{"select * from t1", true},
		{"select /*+ query_time(1000000) */ * from a.t1", true},
		{"select /*+ query_time(1000000) */ * from t1 as a", true},
		{"select /*+ query_time(1000000) */ * from t2 where id > 1", true},
		{"select /*+ query_time(1000000) */ * from t1 where id > 1 or age > 20", true},
		{"select /*+ query_time(1000000) */ * from t1 where id > 1 or age > 20 or name = 'yueyt'", true},
		{"select /*+ query_time(1000000) */ * from t1 where id in ('A', 'B')", true},
		{"select * from testdb.t1 where id in ('A', 'B')", true},

		{"select /*+ query_time(1000000) */ * from t1 order by name desc", false},
		{"select /*+ query_time(1000000) */ id,name,age from t1 group by id,name,age ", false},
		{"select id,name,age,sum(1) as count from t1 group by id,name,age having count > 1", false},
		{"select a.*,b.* from t1 left join t2 on t1.id = t2.id where t1.id > 10", false},
		{"select a.*,b.* from t1,t2 where t1.id = t2.id", false},
		{"select a.*,b.* from t1,t2", false},
		{"select a.*,b.* from t1 where id in (select id from t2)", false},
		{"select a.*,b.* from t1 as a where id in (select id from t2)", false},
		{"select a.*,b.* from t1 as a where id in (select id from t2)", false},
		{"select a.*,b.*", false},
		{"with a as (select id from t1) select * from a", false},
		{"with a as (select id from t1) select * from a where a.id > 10", false},

		{"insert into tab values(1,2,'3')", false},
		{"delete from tab where id > 1 or name = 'yueyt'", false},
		{"update t1 set name = 'yueyt', addr = 10", false},
	}

	dbc := dbinfo.DbConf{Dbname: "testdb", Username: "root", Password: "root", Addr: "127.0.0.1:3306"}
	conn, err := dbinfo.NewDB(dbc)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	for i, c := range cases {
		stmt, err := NewSelectStmt(c.in)
		if err != nil {
			t.Error(err)
			continue
		}

		tableOrig, isOneTable := ExtractTable(stmt)
		if isOneTable == false {
			log.Printf("case %02d 复杂:[%s]", i, c.in)
			continue
		}

		if isSimpleSql := JudgeSimpleSQL(stmt); !isSimpleSql {
			log.Printf("case %02d 复杂:[%s]", i, c.in)
			continue
		}

		schema, table, err := GetTabAndSchema(tableOrig, dbc.Dbname)
		if err != nil {
			log.Printf("err:%v", err)
			continue
		}

		keyNameStr, err := dbinfo.GetKeyName(conn, schema, table)
		if err != nil {
			log.Printf("err:%v", err)
			continue
		}

		keyVals, err := dbinfo.GetKeyValueByStep(conn, GetOrigWhere(stmt), schema, table, keyNameStr, 1)
		if err != nil {
			log.Printf("err:%+v", err)
			continue
		}

		keyRanges, err := dbinfo.GetKeyRange(keyVals)
		if err != nil {
			log.Printf("err:%v", err)
			continue
		}

		log.Printf("case %02d 简单:[%s], 单表:%s", i, c.in, table)
		for _, vv := range keyRanges {
			subWhere := fmt.Sprintf("(%s)>=%s and (%s)<%s", keyNameStr, vv.Start, keyNameStr, vv.End)
			newSql := AddSubWhereSql(stmt, subWhere)
			log.Printf("case %02d 改写:[%s]", i, newSql)
		}
	}
}

func TestParse(t *testing.T) {
	for i, v := range []string{
		"select * from t1 where id > 1 and name < 10 order by name",
		"select * from t1",
		"select /*+ query_timeout(1000) */ id from t1",
	} {
		log.Printf(">>> [%s]", v)
		stmt, err := NewSelectStmt(v)
		if err != nil {
			t.Error(err)
		}

		if stmt.TableHints != nil {
			log.Printf("%d->%s", i, v)
		}

	}
}
