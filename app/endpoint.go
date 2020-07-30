package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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
	jsonResp, err := json.Marshal(z)
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
			StartTime:   "02-08-2020",
			Alarm:       true,
		},
		task{
			Name:        "Uche",
			Description: "List of all tasks created",
			StartTime:   "08-08-2020",
			Alarm:       false,
		},
	}
	jsonResp, err := json.Marshal(tasks)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprint(w, string(jsonResp))
}

func RetrieveAPIView(w http.ResponseWriter, req *http.Request) {
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

	jsonResp, err := json.Marshal(tasks[0])
	if err != nil {
		fmt.Printf("Error marshalling json %v", err)
	}
	fmt.Fprint(w, string(jsonResp))
}

func DeleteTaskAPIView(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodDelete:
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
		tasks = append(tasks[:2], tasks[3:]...)
		fmt.Fprint(w, "Item with ID "+strconv.Itoa(len(tasks))+" has been successully deleted")
	default:
		fmt.Fprint(w, "Method Not allowed")
	}

}
