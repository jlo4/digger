package directoryUtils

import (
	"os"
	"path/filepath"

	l "digger/logUtils"
)

func GetWorkingDirectory() (dir string) {
	currentDirectory, err := os.Getwd()
	if l.LogFatalError(err) {
		return
	}
	return currentDirectory
}

func GetDirectoryList(dir string) []os.DirEntry {
	dirEntry, err := os.ReadDir(dir)
	if l.LogFatalError(err) {
		return nil
	}
	return dirEntry
}

func GetFullPath(dir string) string {
	fullpath, err := filepath.Abs(dir)
	if l.LogFatalError(err) {
		return ""
	}

	fullName, err := filepath.Abs(fullpath)
	if l.LogFatalError(err) {
		return ""
	}

	return fullName
}

func IsDirectory(dirEntry string) bool {
	path, err := filepath.Abs(dirEntry)
	if l.LogFatalError(err) {
		return false
	}

	fileInfo, err := os.Stat(path)
	if l.LogFatalError(err) {
		return false
	}

	return fileInfo.IsDir()
}
