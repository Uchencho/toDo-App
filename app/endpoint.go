package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Uchencho/toDo-App/models"
)

var Db = models.ConnectDatabase()

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

		Db.Create(&b)

		jsonResp, err := json.Marshal(b)
		if err != nil {
			fmt.Println(err)
		}
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, string(jsonResp))

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, `{"Message":"Method not allowed"}`)
	}
}

func ListAPIView(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	switch req.Method {
	case http.MethodGet:
		var b []models.Task

		Db.Find(&b)

		jsonResp, err := json.Marshal(b)
		if err != nil {
			fmt.Println(err)
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(jsonResp))

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, `{"Message":"Method not allowed"}`)
	}
}

func TaskHandler(w http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(strings.TrimPrefix(req.URL.Path, "/tasks/"))
	w.Header().Set("Content-Type", "application/json")

	switch req.Method {
	case http.MethodGet:
		var b models.Task

		Db.Find(&b, id)
		if b.ID == 0 {
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, `{"Message":"No task with that ID"}`)
			return
		}

		jsonResp, err := json.Marshal(b)
		if err != nil {
			fmt.Println(err)
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(jsonResp))
	case http.MethodPut:
		var b models.Task
		var z models.Updatetask

		// Initialize the mode

		Db.Find(&b, id)
		if b.ID == 0 {
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, `{"Message":"No task with that ID"}`)
			return
		}

		// Decode what is sent
		err := json.NewDecoder(req.Body).Decode(&z)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		// Update records that are available
		Db.Model(&b).Updates(models.Task{Name: z.Name,
			Description: z.Description,
			StartTime:   z.StartTime,
			Alarm:       z.Alarm})

		jsonResp, err := json.Marshal(b)
		if err != nil {
			fmt.Printf("Error marshalling json %v", err)
		}
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprint(w, string(jsonResp))
	case http.MethodDelete:
		var b models.Task

		Db.Find(&b, id).Delete(&b) // Delete does not throw error if ID not found
		w.WriteHeader(http.StatusNoContent)
		fmt.Fprint(w, `{"Message":"Successfully deleted"}`)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, `{"Message":"Method not allowed"}`)
	}
}
