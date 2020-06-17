package i18n

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"reflect"
	"strings"
)

func T(language, targetKey string) (interface{}, error) {
	configMap := map[interface{}]interface{}{}
	buffer, err := ioutil.ReadFile(fmt.Sprintf("./config/locales/%v.yml", language))
	err = yaml.Unmarshal(buffer, &configMap)
	if err != nil {
		log.Fatalf(err.Error())
	}

	if targetKey == "" {
		return "", errors.New("targetKey should not be an empty string")
	}

	keys := strings.Split(targetKey, ".")
	for index, key := range keys {
		config, ok := configMap[key]
		if !ok {
			return "", errors.New(fmt.Sprintf("targetKey: %v does Not exist", targetKey))
		}
		configType := reflect.TypeOf(config)
		//log.Printf("config: %v, configType.Kind(): %v", config, configType.Kind())
		if configType.Kind() == reflect.String {
			if index == len(keys)-1 {
				return reflect.ValueOf(config), nil
			}
		} else if configType.Kind() == reflect.Map {
			valueOfConfig := reflect.ValueOf(config)
			configMap = valueOfConfig.Interface().(map[interface{}]interface{})
		}else {
			return "", errors.New("configMap type error")
		}
	}

	return "", errors.New(fmt.Sprintf("targetKey: %v does Not exist", targetKey))
}
