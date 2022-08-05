package core

import (
	"fmt"
	"testing"
)

var payload []byte

func init() {
	payload = []byte(`
20220202:
  activities:
    - activity1
    - activity2
  meetings:
    - meeting1
  todo:
    feb tasks: false
20220128:
  activities:
    - activity1
    - activity2
  meetings:
    - meeting1
  todo:
    stuff: false
20220127:
  activities:
    - activity1
    - activity2
  meetings:
    - meeting1
  todo:
    things: false
20220120:
  activities:
    - activity1
    - activity2
  meetings:
    - meeting1
  todo:
    do the thing: false
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

	dayml, err := sortedDaymlList(payload)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	fmt.Print(dayml)
}

func TestDaymlObjectsAreSortedInChronOrder(t *testing.T) {
	dayml, err := sortedDaymlList(payload)
	if err != nil {
		t.Errorf("error: %v", err)
	}

	expectedDates := []int{
		20220101,
		20220120,
		20220127,
		20220128,
		20220202,
	}
	for i, day := range dayml {
		if day.Date != expectedDates[i] {
			t.Errorf("error: expected %d got %d", expectedDates[i], day.Date)
		}
	}
}
