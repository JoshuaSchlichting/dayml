package main

import (
	"fmt"
	"testing"
)

var payload []byte

func init() {
	payload = []byte(`
20220101:
  activities:
    - activity1
    - activity2
  meetings:
    - meeting1
  todo:
    task1: false
    task2: true
20220101:
  activities:
    - activity1
	- activity2
  meetings:
    - meeting1
  todo:
    task1: false
    task2: true
`)
}

func TestCreateDaymlFromPayload(t *testing.T) {

	dayml, err := NewDaymlList(payload)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	fmt.Print(dayml)

	if dayml[0].Date != 20220101 {
		t.Errorf("error: %v", err)
	}

}
