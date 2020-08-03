package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type task struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	StartTime   string `json:"start-time"`
	Alarm       bool   `json:"alarm"`
}

func createTasks() []task {
	tasks := []task{
		task{
			Name:        "Nils",
			Description: "Create issues on github to solve",
			StartTime:   "02-08-2020",
			Alarm:       true,
		},
		task{
			Name:        "Uche",
			Description: "Create list view endpoint",
			StartTime:   "08-08-2020",
			Alarm:       false,
		},
		task{
			Name:        "Uche",
			Description: "Complete OkraGo app",
			StartTime:   "09-08-2020",
			Alarm:       true,
		},
		task{
			Name:        "Uche",
			Description: "Start learning Goroutines",
			StartTime:   "19-08-2020",
			Alarm:       false,
		},
	}
	return tasks
}

func getTask(id int) task {
	return createTasks()[id]
}

func createEntry(alarm bool, name, description, startTime string) task {

	u := task{
		Name:        name,
		Description: description,
		StartTime:   startTime,
		Alarm:       alarm,
	}
	return u
}

func CreateEntryEndpoint(w http.ResponseWriter, req *http.Request) {
	z := createEntry(false, "Uche", "First To DO entry", "01-08-2020")
	jsonResp, err := json.Marshal(z)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprint(w, string(jsonResp))
}

func ListAPIView(w http.ResponseWriter, req *http.Request) {
	tasks := createTasks()
	jsonResp, err := json.Marshal(tasks)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprint(w, string(jsonResp))
}

func TaskHandler(w http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(strings.TrimPrefix(req.URL.Path, "/tasks/"))
	w.Header().Set("Content-Type", "application/json")

	switch req.Method {
	case http.MethodGet:
		jsonResp, err := json.Marshal(getTask(id))
		if err != nil {
			fmt.Printf("Error marshalling json %v", err)
		}
		w.WriteHeader(200)
		fmt.Fprint(w, string(jsonResp))
	case http.MethodPut:
		u := getTask(id)
		u.Alarm = true

		jsonResp, err := json.Marshal(getTask(id))
		if err != nil {
			fmt.Printf("Error marshalling json %v", err)
		}
		w.WriteHeader(201)
		fmt.Fprint(w, string(jsonResp))
	case http.MethodDelete:
		// tasks := createTasks()
		w.WriteHeader(204)
		// tasks = append(tasks[:id], tasks[id+1:]...)
		fmt.Fprint(w, "Item with ID "+strconv.Itoa(id)+" has been successully deleted")
	default:
		w.WriteHeader(400)
		fmt.Fprint(w, "Method is not allowed")
	}
}
