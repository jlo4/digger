package fileUtils

import (
	"io"
	"os"
	"os/exec"

	l "digger/logUtils"
)

func IsFile(fileName string) bool {
	file, err := os.Lstat(fileName)
	if l.LogFatalError(err) {
		return false
	}
	return file.Mode().IsRegular()
}

func canOpenFile(fileName string) bool {
	file, err := os.Lstat(fileName)
	if l.LogFatalError(err) {
		return false
	}
	return file.Mode().Perm().IsRegular()
}

func GetPermissions(fileName string) string {
	file, err := os.Lstat(fileName)
	if l.LogFatalError(err) {
		return ""
	}
	return file.Mode().Perm().String()
}

func OpenFile(fileName string) string {

	file, err := os.Open(fileName)
	if l.LogFatalError(err) {
		return ""
	}

	const chunk = 4
	b := make([]byte, chunk)

	var readFile int = 0
	for {
		readFile, err = file.Read(b)
		if err != nil {
			if err != io.EOF {
				l.LogFatalError(err)
				return ""
			}
			break
		}

	}
	return string(b[:readFile])
}

func OpenFileWithProgram(programName string, fileName string) {
	if canOpenFile(fileName) {
		cmd := exec.Command(programName, fileName)
		err := cmd.Run()
		if l.LogFatalError(err) {
			return
		}
	}
}
