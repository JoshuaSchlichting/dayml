package core

import (
	"fmt"
	"testing"
)

var payload []byte

func init() {
	payload = []byte(`
20220102:
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

	if dayml[1].Date != 20220101 || dayml[1].Todo["task1"] != false {
		t.Errorf("error: %v", err)
	}

}
