package utils

import (
	"io/ioutil"
	"log"
)

func ReadFilePdf(file string) string {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	return string(data)
}
