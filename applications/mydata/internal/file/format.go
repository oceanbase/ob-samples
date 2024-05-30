package file

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type FileFormat struct {
	FieldsTerminated string
	FieldsEnclosed   string
	EnclosedOptFlag  bool
	FieldsEscaped    string
	LinesTerminated  string
}

// var FieldsTerminated = 0x2c // 列分隔符：, 允许多字节
// var FieldsEnclosed = 0x27   // 包裹符：'  单字节要求
// var FieldsEnclosed = 0x22 // 包裹符：" 单字节要求
// var FieldsEscaped = 0x5c  // 转义符：\ 单字节要求
// var LinesTerminated = 0x0a // 行分隔符：\n 允许多字节
var (
	FieldsTerminated []byte // 多字节
	LinesTerminated  []byte // 多字节
	FieldsEscaped    byte   // 单字节
	FieldsEnclosed   byte   // 单字节
)

func (f FileFormat) String() string {
	return fmt.Sprintf("{FieldsTerminated:%#x LinesTerminated:%#x FieldsEscaped:%v FieldsEnclosed:%v EnclosedOptFlag:%v}",
		string(FieldsTerminated), string(LinesTerminated), string(FieldsEscaped), string(FieldsEnclosed), f.EnclosedOptFlag)
}

func (f *FileFormat) AdjustAndSetFlags() error {
	if strings.HasPrefix(f.FieldsTerminated, "0x") && len(f.FieldsTerminated) > 2 {
		decodeString, err := hex.DecodeString(f.FieldsTerminated[2:])
		if err != nil {
			return errors.Wrapf(err, "%v", f.FieldsTerminated)
		}
		FieldsTerminated = decodeString
	} else {
		FieldsTerminated = []byte(f.FieldsTerminated)
	}
	if len(FieldsTerminated) == 0 {
		return errors.New("FIELDS_TERMINATED is null")
	}

	if strings.HasPrefix(f.LinesTerminated, "0x") && len(f.LinesTerminated) > 2 {
		decodeString, err := hex.DecodeString(f.LinesTerminated[2:])
		if err != nil {
			return errors.Wrapf(err, "%v", f.LinesTerminated)
		}
		LinesTerminated = decodeString
	} else {
		LinesTerminated = []byte(f.LinesTerminated)
	}

	if len(LinesTerminated) == 0 {
		return errors.New("LINES_TERMINATED is null")
	}

	if len(f.FieldsEscaped) > 0 {
		FieldsEscaped = []byte(f.FieldsEscaped)[0]
	}

	if len(f.FieldsEnclosed) > 0 {
		FieldsEnclosed = []byte(f.FieldsEnclosed)[0]
	}

	return nil
}
