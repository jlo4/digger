package directoryUtils

import (
	"os"
	"path/filepath"

	l "digger/logUtils"
)

func WalkDownDirectory(dir string) {
	if IsDirectory(dir) {
		err := os.Chdir(dir)
		if l.LogFatalError(err) {
			return
		}
	}
}

func WalkUpDirectory() {
	str, err := os.Getwd()
	if l.LogFatalError(err) {
		return
	}
	err = os.Chdir(filepath.Dir(str))
	if l.LogFatalError(err) {
		return
	}
}
