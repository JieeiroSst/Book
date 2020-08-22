package utils

import (
	"io/ioutil"
)

func ReadFilePdf(file string) string {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err.Error())
	}
	return string(data)
}
