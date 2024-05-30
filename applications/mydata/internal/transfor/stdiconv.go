package transfor

import (
	"bytes"
	"log"

	"github.com/djimenez/iconv-go"
	"github.com/pkg/errors"
)

func StdIconv(inByte, outByte []byte, fromCode, toCode string) error {
	converter, err := iconv.NewConverter(fromCode, toCode)
	if err != nil {
		return errors.WithStack(err)
	}
	defer converter.Close()

	read, written, err := converter.Convert(inByte, outByte)
	if err != nil {
		return errors.WithStack(err)
	}

	log.Printf("read:%d, written:%d", read, written)
	return nil
}

func StdIconv2() {
	reader, err := iconv.NewReader(bytes.NewReader([]byte("我是中国人")), "utf-8", "gbk")
	if err != nil {
		panic(err)
	}

	buf := make([]byte, 10)
	n, err := reader.Read(buf)
	if err != nil {
		panic(err)
	}

	log.Printf("read:%d, buf:%X", n, buf)
}
