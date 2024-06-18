package utils

import (
	"errors"
	"os"
)

func FileExists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}

func FilesExists(names []string) bool {
	for _, name := range names {
		ok, _ := FileExists(name)
		if ok {
			return ok
		}
	}
	return false
}
