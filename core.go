package main

import (
	"log"
	"reflect"

	yaml "gopkg.in/yaml.v3"
)

type Todo map[string]bool

type DaymlRootObject struct {
	Date int
	Todo Todo `yaml:"todo"`
}

func NewDaymlList(yamlPayload []byte) (rootObject []DaymlRootObject, err error) {
	interfaceMap := make(map[int]map[string]interface{})
	err = yaml.Unmarshal(yamlPayload, interfaceMap)
	if err != nil {
		return rootObject, err
	}

	for date := range interfaceMap {
		log.Print(reflect.TypeOf(interfaceMap[date]["todo"]))
		newTodo := make(map[string]bool)

		for i, v := range interfaceMap[date]["todo"].(map[string]interface{}) {
			newTodo[i] = v.(bool)
		}

		daymlRootObject := DaymlRootObject{
			Date: date,
			Todo: newTodo,
		}

		rootObject = append(rootObject, daymlRootObject)
	}
	return rootObject, nil
}
