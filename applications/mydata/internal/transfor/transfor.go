package transfor

import (
	"github.com/pkg/errors"
	"golang.org/x/text/encoding/simplifiedchinese"
)

func ConvStr2GBK(b []byte) ([]byte, error) {
	bs, err := simplifiedchinese.GBK.NewEncoder().Bytes(b)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return bs, nil
}

func ConvGBK2Str(gb []byte) ([]byte, error) {
	bs, err := simplifiedchinese.GBK.NewDecoder().Bytes(gb)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return bs, nil
}
