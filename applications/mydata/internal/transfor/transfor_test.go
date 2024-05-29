package transfor

import (
	"bytes"
	"log"
	"testing"
)

func TestConvGBK2Str(t *testing.T) {
	srcByte := []byte("中国人")
	gbkByte, err := ConvStr2GBK(srcByte)
	if err != nil {
		t.Error(err)
	}
	log.Printf(">>>%x", gbkByte)

	utf8Byte, err := ConvGBK2Str(gbkByte)
	if err != nil {
		t.Error(err)
	}

	if bytes.Compare(srcByte, utf8Byte) != 0 {
		log.Printf("srcByte:%X, utf8Byte:%X", srcByte, utf8Byte)
	}

	log.Printf("srcByte:%X, utf8Byte:%X", srcByte, utf8Byte)
}
