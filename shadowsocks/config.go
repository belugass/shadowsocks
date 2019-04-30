package shadowsocks

import (
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"os"
	"reflect"
)

type Config struct {
	Server     string `json:"server"`
	ServerPort int    `json:"server_port"`
	Local      string `json:"local"`
	LocalPort  int    `json:"local_port"`
	Method     string `json:"method"`
	Password   string `json:"password"`

	Timeout int `json:"timeout"`
}

func ParseConfig(filepath string) (config *Config, err error) {

	file, err := os.OpenFile(filepath, os.O_RDONLY, 0)
	if err != nil {
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}

	if err := jsoniter.Unmarshal(data, config); err != nil {
		return nil, err
	}
	return
}

func UpdateConfig(config *Config, annotherConfig *Config) {

	newValue := reflect.ValueOf(config).Elem()
	oldValue := reflect.ValueOf(annotherConfig).Elem()
	for i := 0; i < newValue.NumField(); i++ {

		newField := newValue.Field(i)
		oldField := oldValue.Field(i)

		switch newField.Kind() {
		case reflect.Int:
			if v := oldField.Int(); v != 0 {
				newField.SetInt(v)
			}
		case reflect.String:
			if v := oldField.String(); v != "" {
				newField.SetString(v)
			}
		}
	}
}
