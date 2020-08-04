package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Uchencho/toDo-App/models"
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

func CreateEntryEndpoint(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	switch req.Method {
	case http.MethodPost:
		var b models.Task

		err := json.NewDecoder(req.Body).Decode(&b)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		db := models.SetupModels()
		defer db.Close()
		db.Create(&b)

		jsonResp, err := json.Marshal(b)
		if err != nil {
			fmt.Println(err)
		}
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, string(jsonResp))

	default:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"Message":"Method not allowed"}`)
	}
}

func ListAPIView(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	switch req.Method {
	case http.MethodGet:
		var b []models.Task

		db := models.SetupModels()
		defer db.Close()
		db.Find(&b)

		jsonResp, err := json.Marshal(b)
		if err != nil {
			fmt.Println(err)
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(jsonResp))

	default:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"Message":"Method not allowed"}`)
	}
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
