package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"

	"mydata/app/export"
	"mydata/internal/dbinfo"
	"mydata/internal/file"
)

var defaultFormat = file.FileFormat{
	FieldsTerminated: ",",
	FieldsEnclosed:   "",
	EnclosedOptFlag:  false,
	FieldsEscaped:    "\\",
	LinesTerminated:  "\n",
}

var (
	ff  = defaultFormat
	dbc = dbinfo.DbConf{}

	fileName string
	querySql string
	bufSize  int
	concNum  int
	notMerge bool
)

var expCmd = &cobra.Command{
	Use:     "export",
	Aliases: []string{"exp"},
	Short:   "Export data from remote database server",
	RunE:    RunExpE,
}

func init() {
	expCmd.Flags().StringVarP(&dbc.Addr, "addr", "a", "", "mysql database addr, format: ip:port")
	expCmd.Flags().StringVarP(&dbc.Username, "username", "u", "", "username for connect database")
	expCmd.Flags().StringVarP(&dbc.Password, "password", "p", "", "password for connect database")
	expCmd.Flags().StringVarP(&dbc.Dbname, "dbname", "D", "", "default database name")
	expCmd.Flags().StringVarP(&querySql, "query-sql", "e", "", "select sql")
	expCmd.Flags().StringVarP(&fileName, "output", "o", "", "output filename")

	expCmd.Flags().StringVar(&ff.FieldsTerminated, "fields-terminated", defaultFormat.FieldsTerminated, "fields terminated")
	expCmd.Flags().StringVar(&ff.FieldsEnclosed, "fields-enclosed", defaultFormat.FieldsEnclosed, "fields enclosed")
	expCmd.Flags().StringVar(&ff.FieldsEscaped, "fields-escaped", defaultFormat.FieldsEscaped, "fields escaped")
	expCmd.Flags().StringVar(&ff.LinesTerminated, "lines-terminated", defaultFormat.LinesTerminated, "lines terminated")
	expCmd.Flags().BoolVar(&ff.EnclosedOptFlag, "enclosed-optionally", defaultFormat.EnclosedOptFlag, "fields enclosed optionally")

	expCmd.Flags().StringVar(&dbc.Params, "params", "timeout=3s", "connection Params")
	expCmd.Flags().IntVar(&bufSize, "buf-size", 1024*32, "buf size for write outfile")
	expCmd.Flags().IntVar(&concNum, "concurrency", 5, "concurrency number")
	expCmd.Flags().BoolVar(&notMerge, "not-merge", false, "merge chunks to one file")

	_ = expCmd.MarkFlagRequired("output")
	_ = expCmd.MarkFlagRequired("query-sql")
	_ = expCmd.MarkFlagRequired("username")
	_ = expCmd.MarkFlagRequired("password")
	_ = expCmd.MarkFlagRequired("addr")
	_ = expCmd.MarkFlagRequired("dbname")

	expCmd.Flags().SortFlags = false // 禁止flag排序

	rootCmd.AddCommand(expCmd)
}

func RunExpE(*cobra.Command, []string) error {
	defer func(startAt time.Time) {
		log.Printf("execute cmd elapse time(s):%.3f", time.Since(startAt).Seconds())
	}(time.Now())

	// 处理输入参数
	if err := ff.AdjustAndSetFlags(); err != nil {
		return err
	}

	// 判断输入fileName必须为文件
	stat, err := os.Stat(fileName)
	if err == nil && stat.IsDir() {
		return errors.Errorf("input:%s is directory, must include filename", fileName)
	}

	// 默认分块并发，获取query range，
	queryWithRanges, err := dbinfo.GetChunkQuery(dbc, querySql, concNum)
	if err != nil {
		// 获取chunk失败就用原始sql单并发执行
		if debug {
			log.Printf("%v", err)
		}
		queryWithRanges = []string{querySql}
	}

	// 获取输入sql的range信息
	wg := errgroup.Group{}
	pid := os.Getpid()

	outFiles := make([]string, len(queryWithRanges))
	for i, q := range queryWithRanges {
		rangeSeq := i
		// filename = input prefix+rangeSeq+pid
		outFiles[rangeSeq] = fmt.Sprintf("%s.%d.%d", fileName, rangeSeq, pid)
		query := q
		wg.Go(func() error {
			out, err := file.NewMyOutFile(outFiles[rangeSeq], bufSize)
			if err != nil {
				return err
			}
			defer out.Close()

			exp := export.Exporter{ChunkSeq: rangeSeq, DbConf: dbc, Query: query, EnclosedOps: ff.EnclosedOptFlag, OutFile: out}

			return exp.ExpData()
		})
	}
	log.Printf("chunk:%d, outfiles:%v", len(queryWithRanges), outFiles)

	// 等待所有并发任务结束
	if err := wg.Wait(); err != nil {
		return errors.WithStack(err)
	}

	if notMerge { // 不合并文件,直接退出
		return nil
	}

	// 合并临时文件
	if err := file.MergeAndCleanN(outFiles[0], outFiles[1:]); err != nil {
		return err
	}

	// 重命名目标文件
	if err := os.Rename(outFiles[0], fileName); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
