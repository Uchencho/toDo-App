package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type task struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	StartTime   string `json:"start-time"`
	Alarm       bool   `json:"alarm"`
}

func createEntry(alarm bool, name, description, startTime string) task {

	u := task{
		Name:        name,
		Description: description,
		StartTime:   startTime,
		Alarm:       alarm,
	}

	// var tasks []task
	// tasks = append(tasks, u)
	// fmt.Println(tasks)
	return u

}

func CreateEntryEndpoint(w http.ResponseWriter, req *http.Request) {
	z := createEntry(false, "Uche", "First To DO entry", "01-08-2020")
	u := &z
	jsonResp, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprint(w, string(jsonResp))

}

func ListAPIView(w http.ResponseWriter, req *http.Request) {
	tasks := []task{
		task{
			Name:        "Nils",
			Description: "test",
			StartTime:   "startTime",
			Alarm:       true,
		},
		task{
			Name:        "Uche",
			Description: "List of all tasks created",
			StartTime:   "08-08-2020",
			Alarm:       false,
		},
	}
	taskadd := &tasks
	jsonResp, err := json.Marshal(taskadd)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprint(w, string(jsonResp))
}
