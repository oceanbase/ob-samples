package export

import (
	"context"
	"log"

	"github.com/pkg/errors"

	"mydata/internal/dbinfo"
	"mydata/internal/file"
)

type Exporter struct {
	ChunkSeq    int             // 分块序号
	DbConf      dbinfo.DbConf   // 数据库配置信息
	Query       string          // 输入查询sql
	EnclosedOps bool            // 包裹符标记
	OutFile     *file.MyOutfile // 输出文件
}

func (e *Exporter) ExpData() error {
	conn, err := dbinfo.NewDB(e.DbConf)
	if err != nil {
		return err
	}
	defer func() {
		_ = conn.Close()
	}()

	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	log.Printf("[chunk:%02d]execute query sql:[%s]", e.ChunkSeq, e.Query)
	rows, err := conn.QueryContext(ctx, e.Query)
	if err != nil {
		return errors.Wrapf(err, "query sql:[%s]", e.Query)
	}
	defer func() {
		_ = rows.Close()
	}()

	isChar, isEnclosed, err := dbinfo.MakeColFlags(rows, e.EnclosedOps)
	if err != nil {
		return err
	}

	scanArgs, scanValues, err := dbinfo.MakeScanBuf(rows)
	if err != nil {
		return err
	}

	log.Printf("[chunk:%02d]scan rows into buffer ...", e.ChunkSeq)
	count := 0

	for rows.Next() {
		if err := rows.Scan(scanArgs...); err != nil {
			return errors.WithStack(err)
		}

		// scan rows
		for colIndex, col := range scanValues {
			if colIndex > 0 {
				// 写入列分隔符
				e.OutFile.Write(file.FieldsTerminated)
			}

			if col == nil { // 字段为null的转义
				e.OutFile.WriteSingleByte(file.FieldsEscaped)
				e.OutFile.WriteSingleByte('N')
				continue
			}

			// 写入前包裹符号
			if isEnclosed[colIndex] && file.FieldsEnclosed > 0 {
				e.OutFile.WriteSingleByte(file.FieldsEnclosed)
			}
			// 写入列值
			// 针对该列的字段，只能针对字符类型
			if isChar[colIndex] { // 是字符类型
				for ii := range col {
					if col[ii] == file.FieldsEscaped || col[ii] == file.FieldsEnclosed ||
						col[ii] == file.LinesTerminated[0] || col[ii] == file.FieldsTerminated[0] {
						e.OutFile.WriteSingleByte(file.FieldsEscaped)
					}

					e.OutFile.Write([]byte{col[ii]})
				}
			} else {
				e.OutFile.Write(col)
			}

			// 写入后包裹符号
			if isEnclosed[colIndex] && file.FieldsEnclosed > 0 {
				e.OutFile.WriteSingleByte(file.FieldsEnclosed)
			}
		}

		// 写入行分割符
		e.OutFile.Write(file.LinesTerminated)

		// 统一处理错误异常
		if e.OutFile.Err != nil {
			return errors.WithStack(e.OutFile.Err)
		}

		count++
	}

	if err := rows.Err(); err != nil {
		return errors.WithStack(err)
	}

	log.Printf("[chunk:%02d]finished successfully, records:%d", e.ChunkSeq, count)

	return nil
}
