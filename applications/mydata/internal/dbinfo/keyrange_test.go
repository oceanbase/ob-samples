package dbinfo

import (
	"log"
	"testing"
)

func TestGetChunkQuery(t *testing.T) {
	dbc := DbConf{Addr: "192.168.111.3:3306", Username: "root", Password: "root", Dbname: "testdb"}
	for _, v := range []int{2, 3, 4, 5, 10} {
		query, err := GetChunkQuery(dbc, "select * from t1", v)
		if err != nil {
			t.Fatal(err)
		}
		for ii, vv := range query {
			log.Printf(">>>v:%d, %d, sql:%s", v, ii, vv)
		}
	}
}
