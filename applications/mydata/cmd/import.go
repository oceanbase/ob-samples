package cmd

import (
	"bytes"
	"log"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"mydata/internal/file"
)

var impCmd = &cobra.Command{
	Use:     "import",
	Aliases: []string{"imp"},
	Short:   "Import data to remote database server",
	RunE:    RunImpE,
}

func init() {
	impCmd.Flags().StringVarP(&dbc.Addr, "addr", "a", "", "mysql database addr, format: ip:port")
	impCmd.Flags().StringVarP(&dbc.Username, "username", "u", "", "username for connect database")
	impCmd.Flags().StringVarP(&dbc.Password, "password", "p", "", "password for connect database")
	impCmd.Flags().StringVarP(&dbc.Dbname, "dbname", "D", "", "default database name")
	impCmd.Flags().StringVarP(&querySql, "query-sql", "e", "", "select sql")
	impCmd.Flags().StringVarP(&fileName, "file-name", "o", "", "output filename")

	impCmd.Flags().StringVar(&ff.FieldsTerminated, "fields-terminated", defaultFormat.FieldsTerminated, "fields terminated")
	impCmd.Flags().StringVar(&ff.FieldsEnclosed, "fields-enclosed", defaultFormat.FieldsEnclosed, "fields enclosed")
	impCmd.Flags().StringVar(&ff.FieldsEscaped, "fields-escaped", defaultFormat.FieldsEscaped, "fields escaped")
	impCmd.Flags().StringVar(&ff.LinesTerminated, "lines-terminated", defaultFormat.LinesTerminated, "lines terminated")
	impCmd.Flags().BoolVar(&ff.EnclosedOptFlag, "enclosed-optionally", defaultFormat.EnclosedOptFlag, "fields enclosed optionally")

	impCmd.Flags().StringVar(&dbc.Params, "params", "", "connection Params")
	impCmd.Flags().IntVar(&bufSize, "buf-size", 1024*32, "buf size for write outfile")
	impCmd.Flags().IntVar(&concNum, "concurrency", 5, "concurrency number")
	impCmd.Flags().BoolVar(&notMerge, "not-merge", false, "merge chunks to one file")

	_ = impCmd.MarkFlagRequired("file-name")
	_ = impCmd.MarkFlagRequired("query-sql")
	_ = impCmd.MarkFlagRequired("username")
	_ = impCmd.MarkFlagRequired("password")
	_ = impCmd.MarkFlagRequired("addr")
	_ = impCmd.MarkFlagRequired("dbname")

	impCmd.Flags().SortFlags = false // 禁止flag排序

	// TODO: import
	//rootCmd.AddCommand(impCmd)
}

func RunImpE(*cobra.Command, []string) error {
	log.Printf("filename:%s", fileName)
	// 根据行、列分隔符，解析文件
	srcFile, err := os.Open(fileName)
	if err != nil {
		return errors.WithStack(err)
	}
	defer srcFile.Close()

	scanner := file.NewScannerDelim(srcFile, bufSize)
	for scanner.Scan() {
		row := bytes.Split(scanner.Bytes(), file.FieldsTerminated)
		log.Printf(">>> row:%v", row)
	}

	if scanner.Err() != nil {
		return errors.WithStack(scanner.Err())
	}

	// 目标库执行
	return nil
}
