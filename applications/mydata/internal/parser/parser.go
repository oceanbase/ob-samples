package parser

import (
	"fmt"
	"log"
	"strings"

	"github.com/pkg/errors"

	"github.com/pingcap/tidb/pkg/parser"
	"github.com/pingcap/tidb/pkg/parser/ast"
	"github.com/pingcap/tidb/pkg/parser/format"
	_ "github.com/pingcap/tidb/pkg/parser/test_driver"
)

func init() {
	log.SetFlags(log.LstdFlags)
}

type tabX struct {
	tabNames []string
}

func (v *tabX) Enter(in ast.Node) (ast.Node, bool) {
	if name, ok := in.(*ast.TableName); ok {
		if len(name.Schema.O) > 0 {
			v.tabNames = append(v.tabNames, fmt.Sprintf("%s.%s", name.Schema.O, name.Name.O))
		} else {
			v.tabNames = append(v.tabNames, name.Name.O)
		}

	}
	return in, false
}

func (v *tabX) Leave(in ast.Node) (ast.Node, bool) {
	return in, true
}

func ExtractTable(rootNode *ast.SelectStmt) (string, bool) {
	v := &tabX{}
	(*rootNode).Accept(v)

	if len(v.tabNames) != 1 {
		return "", false
	}

	return v.tabNames[0], true
}

// NewSelectStmt 创建select stmt
func NewSelectStmt(query string) (*ast.SelectStmt, error) {
	p := parser.New()

	stmt, err := p.ParseOneStmt(query, "", "")
	if err != nil {
		return nil, errors.Errorf("[%s] parse failed", query)
	}

	s, ok := stmt.(*ast.SelectStmt)
	if !ok {
		return nil, errors.Errorf("[%s] is not select sql", query)
	}

	return s, nil
}

// JudgeSimpleSQL 判断query-sql中是否是否符合并发分块要求，
// 当符合要求时，需要按照分块chunk处理where条件拼接chunk range范围
func JudgeSimpleSQL(s *ast.SelectStmt) bool {
	if s.StraightJoin {
		return false
	}

	if s.GroupBy != nil {
		return false
	}

	if s.Having != nil {
		return false
	}

	if s.OrderBy != nil {
		return false
	}

	if s.Limit != nil {
		return false
	}

	return true
}

func GetOrigWhere(s *ast.SelectStmt) string {
	var sb strings.Builder
	var ff = format.RestoreStringWithoutCharset | format.RestoreStringSingleQuotes | format.RestoreSpacesAroundBinaryOperation

	origTxt := s.OriginalText()
	if s.Where != nil {
		//sb.WriteString(origTxt[:s.Where.OriginTextPosition()])
		//sb.WriteString(" and (")
		if err := s.Where.Restore(format.NewRestoreCtx(ff, &sb)); err != nil {
			return origTxt
		}
		//sb.WriteString(") ")
	}

	return sb.String()
}

// 修改where条件
func AddSubWhereSql(s *ast.SelectStmt, cond string) string {
	var sb strings.Builder
	var ff = format.RestoreStringWithoutCharset | format.RestoreStringSingleQuotes | format.RestoreSpacesAroundBinaryOperation

	origTxt := s.OriginalText()
	var isWhere = 0
	var notOnlyWhere = false
	if s.Where != nil {
		sb.WriteString(origTxt[:s.Where.OriginTextPosition()])
		sb.WriteString(cond)
		sb.WriteString(" and (")
		if err := s.Where.Restore(format.NewRestoreCtx(ff, &sb)); err != nil {
			return origTxt
		}
		sb.WriteString(") ")
		isWhere++
		notOnlyWhere = true
	}

	if s.GroupBy != nil {
		if isWhere == 0 {
			sb.WriteString(" where ")
			sb.WriteString(cond)
			isWhere++
		}
		sb.WriteString(" ")
		if err := s.GroupBy.Restore(format.NewRestoreCtx(ff, &sb)); err != nil {
			return origTxt
		}

		notOnlyWhere = true
	}

	if s.OrderBy != nil {
		if isWhere == 0 {
			sb.WriteString(" where ")
			sb.WriteString(cond)
			isWhere++
		}
		sb.WriteString(" ")
		if err := s.OrderBy.Restore(format.NewRestoreCtx(ff, &sb)); err != nil {
			return origTxt
		}

		notOnlyWhere = true
	}

	if s.Limit != nil {
		if isWhere == 0 {
			sb.WriteString(" where ")
			sb.WriteString(cond)
			isWhere++
		}
		sb.WriteString(" ")
		if err := s.Limit.Restore(format.NewRestoreCtx(ff, &sb)); err != nil {
			return origTxt
		}

		notOnlyWhere = true
	}

	if !notOnlyWhere {
		if isWhere == 0 {
			sb.WriteString(origTxt)
			sb.WriteString(" where ")
			sb.WriteString(cond)
			isWhere++
		}
	}

	return sb.String()
}

func GetTabAndSchema(tab string, dbname string) (string, string, error) {
	if strings.Contains(tab, ".") {
		s := strings.Split(tab, ".")
		if len(s) != 2 {
			return "", "", errors.Errorf("tabname format failed:%s", tab)
		}
		return s[0], s[1], nil
	}

	if len(dbname) == 0 {
		return "", "", errors.Errorf("not dbname and no schema")
	}

	return dbname, tab, nil
}
