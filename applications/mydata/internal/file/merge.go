package file

import (
	"io"
	"log"
	"os"

	"github.com/pkg/errors"
)

func MergeAndCleanN(dstName string, srcNames []string) error {
	buf := make([]byte, 64*1024)
	dst, err := os.OpenFile(dstName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		return errors.WithStack(err)
	}
	defer dst.Close()

	for _, srcName := range srcNames {
		if err := func() error {
			src, err := os.OpenFile(srcName, os.O_RDONLY, os.ModePerm)
			if err != nil {
				return errors.WithStack(err)
			}
			defer src.Close()

			written, err := io.CopyBuffer(dst, src, buf)
			if err != nil {
				return errors.WithStack(err)
			}
			log.Printf("merge file:%s, byte size:%d", srcName, written)

			return nil
		}(); err != nil {
			return err
		}

		if err := os.Remove(srcName); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}
