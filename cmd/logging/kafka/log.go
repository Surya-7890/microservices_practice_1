package kafka

import (
	"os"
	"reflect"
	"strings"

	"github.com/Surya-7890/book_store/logging/config"
)

func CreateLogFiles(cfg config.KafkaConfig, fileMap map[string]*os.File) {
	_value := reflect.ValueOf(cfg.Readers)
	if err := os.Mkdir("logs", os.ModeDir); err != nil {
		panic(err)
	}
	for i := 0; i < _value.NumField(); i++ {
		filename := strings.Split(_value.Field(i).String(), "-")[0]
		if _, ok := fileMap[filename]; ok {
			continue
		}

		file, err := os.OpenFile("./logs/"+filename+".log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			continue
		}
		fileMap[filename] = file
	}
}
