package core

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v3"
)

type Todo map[string]bool

type Dayml []DaymlRootObject

func (d Dayml) GetTodoList(isCompleted bool) []string {
	todoList := make([]string, 0)
	for _, rootObject := range d {

		for key, completedStatus := range rootObject.Todo {
			if completedStatus == isCompleted {
				todoList = append(todoList, fmt.Sprintf("%d - %s", rootObject.Date, key))
			}
		}
	}
	return todoList
}

type DaymlRootObject struct {
	Date int
	Todo Todo `yaml:"todo"`
}

func DaymlFromFile(filePath string) (Dayml, error) {
	payload, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	rootObjects, err := sortedDaymlList(payload)
	if err != nil {
		return nil, err
	}
	return rootObjects, nil
}

func sortedDaymlList(yamlPayload []byte) (rootObjects []DaymlRootObject, err error) {
	interfaceMap := make(map[int]map[string]interface{})
	err = yaml.Unmarshal(yamlPayload, interfaceMap)
	if err != nil {
		return rootObjects, err
	}

	sortedKeys := make([]int, 0)
	for key := range interfaceMap {
		sortedKeys = append(sortedKeys, key)
	}

	for i := range sortedKeys {
		newTodo := make(map[string]bool)
		dateKey := sortedKeys[i]
		for k, v := range interfaceMap[dateKey]["todo"].(map[string]interface{}) {
			newTodo[k] = v.(bool)
		}

		daymlRootObject := DaymlRootObject{
			Date: dateKey,
			Todo: newTodo,
		}

		rootObjects = append(rootObjects, daymlRootObject)
	}
	return rootObjects, nil
}
