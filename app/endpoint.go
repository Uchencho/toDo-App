package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/Uchencho/toDo-App/models"
	"github.com/go-playground/validator"
)

type User struct {
	Username        string    `json:"user_name" validate:"required"`
	Email           string    `json:"email" validate:"required,email"`
	Password        string    `json:"password" validate:"required"`
	ConfirmPassword string    `json:"confirm_password" validate:"eqfield=Password"`
	FirstName       string    `json:"first_name,omitempty"`
	LastName        string    `json:"last_name,omitempty"`
	DateJoined      time.Time `json:"date_joined,omitempty"`
	RegistrationID  string    `json:"registration_id,omitempty"`
}

var (
	Db       = models.ConnectDatabase()
	validate = validator.New()
)

// Validates a struct
func validateInput(object interface{}) (error, bool) {

	err := validate.Struct(object)
	if err != nil {

		//Validation syntax is invalid
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Println(err)
			return err, false
		}

		if len(err.(validator.ValidationErrors)) > 1 {
			log.Println("Error is more than one")
			return err, true
		}

		for _, err := range err.(validator.ValidationErrors) {

			// Retrieve json field
			reflectedValue := reflect.ValueOf(object)
			field, _ := reflectedValue.Type().FieldByName(err.StructField())

			var name string
			if name = field.Tag.Get("json"); name == "" {
				name = strings.ToLower(err.StructField())
			}

			switch err.Tag() {
			case "required":
				return fmt.Errorf("%s is required", name), false
			case "email":
				return fmt.Errorf("%s should be a valid email", name), false
			case "eqfield":
				return fmt.Errorf("%s should be the same as %s", name, err.Param()), false
			default:
				return fmt.Errorf("%s is Invalid", name), false
			}
		}
		return err, false
	}
	return nil, false
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
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, `{"Message":"No task with that ID"}`)
			return
		}

		jsonResp, err := json.Marshal(b)
		if err != nil {
			fmt.Println(err)
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(jsonResp))
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
	case http.MethodPut:
		var b models.Task
		var z models.Updatetask

		// Initialize the mode

		Db.Find(&b, id)
		if b.ID == 0 {
			w.WriteHeader(http.StatusBadRequest)
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

		Db.Find(&b, id) // Delete does not throw error if ID not found
		if b.ID == 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, `{"Message":"No task with that ID"}`)
			return
		}
		Db.Delete(&b)
		w.WriteHeader(http.StatusNoContent)
		fmt.Fprint(w, `{"Message":"Successfully deleted"}`)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, `{"Message":"Method not allowed"}`)
	}
}

// Mock register endpoint to test validation package
func Register(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	switch req.Method {
	case http.MethodPost:
		var user User
		err := json.NewDecoder(req.Body).Decode(&user)
		if err != nil {
			log.Println("Could not decode json with error: ", err)
			return
		}

		err, aboveOneField := validateInput(user)
		if aboveOneField {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, `{"error" : "Invalid Payload"}`)
			return
		}
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, `{"error" : "%s"}`, err.Error())
			return
		}

		jsonResp, err := json.Marshal(user)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprint(w, string(jsonResp))
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, `{"error" : "Method not allowed, try post"}`)
	}
}
