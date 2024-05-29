package dbinfo

import (
	"database/sql"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type DbConf struct {
	Addr     string
	Username string
	Password string
	Dbname   string
	Params   string
}

var reg = regexp.MustCompile("[!-~]")

func (c DbConf) String() string {
	return fmt.Sprintf("{Addr:%s, Username:%s, Passwd:%s, DbName:%s, Params:%s}",
		c.Addr, c.Username, reg.ReplaceAllString(c.Password, "*"), c.Dbname, c.Params)
}

func NewDB(conf DbConf) (*sql.DB, error) {
	//dsn := "root:root@tcp(192.168.56.128:3306)/testdb?param=value"
	if strings.Count(conf.Addr, ":") != 1 {
		return nil, errors.Errorf("database addr invalid, format ip:port, but got [%s]", conf.Addr)
	}

	dbc := mysql.NewConfig()
	dbc.Addr = conf.Addr
	dbc.User = conf.Username
	dbc.Passwd = conf.Password
	dbc.DBName = conf.Dbname

	// 处理输入连接参数
	if err := parseDSNParams(dbc, conf.Params); err != nil {
		return nil, err
	}

	connector, err := mysql.NewConnector(dbc)
	if err != nil {
		return nil, errors.Wrapf(err, "create connector faild:[%s:%s@tcp(%s)/%s]",
			conf.Username, reg.ReplaceAllString(conf.Password, "*"), conf.Addr, conf.Dbname)
	}

	conn := sql.OpenDB(connector)
	if err := conn.Ping(); err != nil {
		return nil, errors.Wrapf(err, "ping db faild:[%s:%s@tcp(%s)/%s]", conf.Username,
			reg.ReplaceAllString(conf.Password, "*"), conf.Addr, conf.Dbname)
	}

	return conn, nil
}

func MakeScanBuf(rows *sql.Rows) ([]interface{}, []sql.RawBytes, error) {
	cols, err := rows.Columns()
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	colSize := len(cols)
	scanArgs := make([]interface{}, colSize)
	scanValues := make([]sql.RawBytes, colSize)

	for i := range scanValues {
		scanArgs[i] = &scanValues[i]
	}

	return scanArgs, scanValues, nil
}

func MakeColFlags(rows *sql.Rows, enclosedFlag bool) ([]bool, []bool, error) {
	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	size := len(columnTypes)

	isCharFlag := make([]bool, size)
	isEnclosed := make([]bool, size)
	for i, ct := range columnTypes {
		// 字符类型
		isCharFlag[i] = isCharFunc(ct.DatabaseTypeName())
		// 当opt=true时，只有字符类型添加enclosed
		// 当opt=false时，所有字符类型都添加enclosed

		if enclosedFlag { // 设置opt的话，只有字符类型被包裹
			if isCharFlag[i] {
				isEnclosed[i] = true
			}
		} else { // 默认不设置opt，所有字段都被包裹
			isEnclosed[i] = true
		}
	}

	return isCharFlag, isEnclosed, nil
}

func isCharFunc(s string) bool {
	return s == "CHAR" || s == "VARCHAR" || s == "TIMESTAMP" || s == "TEXT" || s == "LONGTEXT" || s == "TINYTEXT" ||
		s == "MEDIUMTEXT" || s == "LONGBLOB" || s == "BLOB" || s == "TINYBLOB"
}

func GetKeyName(db *sql.DB, schema, table string) (string, error) {
	keys := make([]string, 0)

	// TODO:适配ob的weak模式查询
	sqlKeyName := fmt.Sprintf("select /*+READ_CONSISTENCY(WEAK) */ COLUMN_NAME from information_schema.columns where TABLE_SCHEMA='%s' and TABLE_NAME='%s' and COLUMN_KEY ='PRI' order by ORDINAL_POSITION asc",
		schema, table)
	log.Printf("get key name:[%s]", sqlKeyName)
	rows, err := db.Query(sqlKeyName)
	if err != nil {
		return "", errors.WithStack(err)
	}
	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var s string
		if err := rows.Scan(&s); err != nil {
			return "", errors.WithStack(err)
		}
		keys = append(keys, s)
	}

	if err := rows.Err(); err != nil {
		return "", errors.WithStack(err)
	}

	if len(keys) == 0 {
		return "", errors.Errorf("not found keys")
	}

	return strings.Join(keys, ","), nil
}

func GetKeyValueByStep(db *sql.DB, origWhere string, schema, table string, keyNames string, stepSize int64) ([]string, error) {
	if stepSize <= 0 {
		return nil, errors.New("stepSize is zero")
	}

	// TODO:根据输入条件，计算主键的范围
	var sqlKeyValue string
	if len(origWhere) > 0 {
		sqlKeyValue = fmt.Sprintf("select /*+READ_CONSISTENCY(WEAK) */ %s from %s.%s where %s order by %s",
			keyNames, schema, table, origWhere, keyNames)
	} else {
		sqlKeyValue = fmt.Sprintf("select /*+READ_CONSISTENCY(WEAK) */ %s from %s.%s order by %s",
			keyNames, schema, table, keyNames)
	}
	log.Printf("get key value by step:[%s]", sqlKeyValue)

	rows, err := db.Query(sqlKeyValue)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer rows.Close()

	colTypes, err := rows.ColumnTypes()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	isChars := make([]bool, 0)
	for _, v := range colTypes {
		isChars = append(isChars, isCharFunc(v.DatabaseTypeName()))
	}

	scanArgs, scanVals, err := MakeScanBuf(rows)
	if err != nil {
		return nil, err
	}

	var count int64 = 0
	keyVals := make([]string, 0)
	for rows.Next() {
		if err := rows.Scan(scanArgs...); err != nil {
			return nil, errors.WithStack(err)
		}

		if count%stepSize == 0 {
			keyVals = append(keyVals, getKeyValueCond(scanVals, isChars))
		}

		count++
	}

	if count%stepSize != 0 { // 处理尾部数据
		keyVals = append(keyVals, getKeyValueCond(scanVals, isChars))
	}

	if err := rows.Err(); err != nil {
		return nil, errors.WithStack(err)
	}

	return keyVals, nil
}

func getKeyValueCond(scanVals []sql.RawBytes, isChars []bool) string {
	var buf strings.Builder
	buf.WriteString("(")
	for i, v := range scanVals {
		if i > 0 {
			buf.WriteString(",")
		}
		if isChars[i] {
			buf.WriteString("'")
			buf.Write(v)
			buf.WriteString("'")
		} else {
			buf.Write(v)
		}
	}

	buf.WriteString(")")

	return buf.String()
}
