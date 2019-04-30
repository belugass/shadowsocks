package shadowsocks

import (
	"errors"
	"os"
)

func IsExist(filepath string) (bool, error) {

	fileInfo, err := os.Stat(filepath)
	if err == nil {
		if fileInfo.Mode() & os.ModeType == 0 {
			return true, nil
		}
		return false, errors.New(filepath + "is exist, but is not regular file")
	}

	if os.IsNotExist(err) {
		return false, err
	}

	return false, nil
}
