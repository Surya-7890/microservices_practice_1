package kafka

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

func CreateLogFiles(fileMap map[string]*os.File) {
	_value := reflect.ValueOf(kafkaConfig.Readers)
	if err := os.Mkdir("logs", os.ModeDir); err != nil {
		panic(err)
	}
	for i := 0; i < _value.NumField(); i++ {
		filename := strings.Split(_value.Field(i).String(), "-")[0]
		if _, ok := fileMap[filename]; ok {
			continue
		}
		fmt.Println(filename, _value.Field(i).String())
		file, err := os.OpenFile("./logs/"+filename+".log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			continue
		}
		fileMap[filename] = file
	}
}
