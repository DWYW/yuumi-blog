package conf

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
)

func Read(file string) ConfigSource {
	var filename string

	if ok, _ := regexp.MatchString("^/", file); ok {
		filename = file
	} else {
		path, _ := filepath.Abs(filepath.Dir("."))
		filename = filepath.Join(path, file)
	}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Panicln(err)
	}

	fileExt := filepath.Ext(filepath.Base(file))
	fileExtLen := len(fileExt)

	return ConfigSource{
		Type:   fileExt[1:fileExtLen],
		Source: &bytes,
	}
}
