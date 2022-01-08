package fileUtils

import (
	"fmt"
	"os"

	l "digger/logUtils"
)

func OpenFile() {

}

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
	fmt.Println(file.Mode())
	return file.Mode().IsRegular()
}

func GetPermissions(fileName string) string {
	file, err := os.Lstat(fileName)
	if l.LogFatalError(err) {
		return ""
	}
	return file.Mode().Perm().String()
}
