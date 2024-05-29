package transfor

import (
	"bufio"
	"encoding/hex"
	"log"
	"os"
	"testing"

	"github.com/djimenez/iconv-go"
)

func TestIconvFunc(t *testing.T) {
	dst := make([]byte, 3)
	srcFile, err := os.Open("/Users/abc/workspaces/mydata/gbk.file")
	if err != nil {
		t.Fatal(err)
	}
	defer srcFile.Close()

	dstFile, err := os.Create("utf8")
	if err != nil {
		t.Fatal(err)
	}
	defer dstFile.Close()

	writer := bufio.NewWriter(dstFile)
	defer writer.Flush()

	bufReader := bufio.NewReader(srcFile)

	reader, err := iconv.NewReader(bufReader, "gbk", "utf8")
	if err != nil {
		t.Error(err)
	}

	for {
		//dst := make([]byte, 3)
		n, err := reader.Read(dst)
		if err != nil {
			t.Error(err)
			break
		}
		log.Printf("read:%d, buf:\n%s", n, hex.Dump(dst))
		nn, err := writer.Write(dst)
		if err != nil {
			t.Error(err)
			break
		}
		log.Printf("written:%d", nn)

	}

	//conv, err := myconv.NewConverter("utf-8", "gbk")
	//if err != nil {
	//	t.Error(err)
	//}
	//defer conv.Close()
	//
	//read, written, err := conv.Convert(src, dst)
	//if err != nil {
	//	t.Error(err)
	//}
	//log.Printf("read:%d, written:%d", read, written)
}
