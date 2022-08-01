package core

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v3"
)

type Todo map[string]bool

type DaymlRootObject struct {
	Date int
	Todo Todo `yaml:"todo"`
}

func GetTodoListFromFile(filePath string) (map[string]bool, error) {
	payload, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	rootObjects, err := NewDaymlList(payload)
	if err != nil {
		return nil, err
	}
	return GetToDoList(rootObjects), nil
}

func ReadYamlFile(filePath string) ([]byte, error) {
	fileContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return fileContents, nil
}

func NewDaymlList(yamlPayload []byte) (rootObjects []DaymlRootObject, err error) {
	interfaceMap := make(map[int]map[string]interface{})
	err = yaml.Unmarshal(yamlPayload, interfaceMap)
	if err != nil {
		return rootObjects, err
	}

	for date := range interfaceMap {
		newTodo := make(map[string]bool)

		for i, v := range interfaceMap[date]["todo"].(map[string]interface{}) {
			newTodo[i] = v.(bool)
		}

		daymlRootObject := DaymlRootObject{
			Date: date,
			Todo: newTodo,
		}

		rootObjects = append(rootObjects, daymlRootObject)
	}
	return rootObjects, nil
}

func GetToDoList(rootObjects []DaymlRootObject) map[string]bool {
	todoList := make(Todo)
	for _, rootObject := range rootObjects {
		for key, value := range rootObject.Todo {
			todoList[fmt.Sprintf("%v - %s", rootObject.Date, key)] = value
		}
	}
	return todoList
}
