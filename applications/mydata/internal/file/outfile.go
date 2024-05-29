package file

import (
	"bufio"
	"os"

	"github.com/pkg/errors"
)

type MyOutfile struct {
	f   *os.File
	w   *bufio.Writer
	Err error
}

func NewMyOutFile(fileName string, bufSize int) (*MyOutfile, error) {
	outFile, err := os.Create(fileName)
	if err != nil {
		return nil, errors.Wrapf(err, "file:%s", fileName)
	}

	return &MyOutfile{
		f: outFile,
		w: bufio.NewWriterSize(outFile, bufSize),
	}, nil
}

func (my *MyOutfile) WriteSingleByte(c byte) {
	if my.Err == nil {
		my.Err = my.w.WriteByte(c)
	}
}

func (my *MyOutfile) Write(p []byte) {
	if my.Err == nil {
		if n, err1 := my.w.Write(p); err1 != nil || n != len(p) {
			my.Err = errors.Wrapf(err1, "wrote %d, want %d", n, len(p))
		}
	}
}

func (my *MyOutfile) Close() {
	if my.Err == nil {
		my.Err = my.w.Flush()
	}

	if my.Err == nil {
		my.Err = my.f.Close()
	}
}
