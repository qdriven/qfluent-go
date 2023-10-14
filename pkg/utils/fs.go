package utils

import (
	"errors"
	"os"
)

func PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("File existing, but not a Directory")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
