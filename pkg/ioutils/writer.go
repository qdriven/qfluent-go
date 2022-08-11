package ioutils

import (
	"fmt"
	qlog "github.com/qdriven/qfluent-cli/pkg/log"
	"io/ioutil"
	"os"
	"path/filepath"
)

func WriteFile(destinationBase string, file File, mode os.FileMode) error {
	if file.Discarded {
		qlog.Debugf("File is discarded, not writing: %s", file.RelativePath)
		return nil
	}
	destinationPath := filepath.Join(destinationBase, file.RelativePath)
	qlog.Infof("Writing file %s", destinationPath)
	dir := filepath.Dir(destinationPath)
	err := os.MkdirAll(dir, os.ModeDir|os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating base dir for file: %w", err)
	}
	if err = ioutil.WriteFile(destinationPath, []byte(file.Contents), mode); err != nil {
		return fmt.Errorf("error writing file: %w", err)
	}
	return nil
}
