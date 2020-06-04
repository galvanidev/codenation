package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func SaveFile(name string, data interface{}) {
	file, err := json.MarshalIndent(data, "", "")
	Check(err)

	err = ioutil.WriteFile(name, file, 0644)
	Check(err)
}

func ReadFile(name string) []byte {
	file, err := ioutil.ReadFile(name)
	Check(err)
	return file
}

func Check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}