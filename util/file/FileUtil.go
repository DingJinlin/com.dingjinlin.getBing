package file

import (
	"os"
	"log"
)

func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func CreateDir(path string)(directPath string, err error)  {
	absPath, err := os.Getwd()
	if err != nil {
		log.Panicln(err)
	}

	direct := absPath + "/" + path + "/"

	haveDirect := CheckFileIsExist(direct)
	if !haveDirect {
		err = os.Mkdir(direct, os.ModePerm)
	}

	return direct, err
}
