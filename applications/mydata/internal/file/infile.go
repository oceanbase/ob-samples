package file

import (
	"bufio"
	"bytes"
	"io"
	"os"

	"github.com/pkg/errors"
)

type MyInfile struct {
	f       *os.File
	scanner *bufio.Scanner
	Err     error
}

func NewMyInFile(fileName string, bufSize int) (*MyInfile, error) {
	inFile, err := os.Open(fileName)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &MyInfile{
		f:       inFile,
		scanner: NewScannerDelim(inFile, bufSize),
	}, nil
}

var LineDelim []byte

func NewScannerDelim(r io.Reader, bufSize int) *bufio.Scanner {
	scanner := bufio.NewScanner(r)
	LineDelim = LinesTerminated // 设置行分割符
	scanner.Split(ScanLines)    // 设置行分割符函数

	return scanner
}

func ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.Index(data, LineDelim); i >= 0 {
		return i + len(LineDelim), dropCR(data[0:i]), nil
	}
	if atEOF {
		return len(data), dropCR(data), nil
	}

	return 0, nil, nil
}

func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}
