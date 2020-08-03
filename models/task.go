package models

type Task struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	StartTime   string `json:"start-time"`
	Alarm       bool   `json:"alarm"`
}

type Createtask struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	StartTime   string `json:"start-time" binding:"required"`
	Alarm       bool   `json:"alarm" binding:"required"`
}

type Updatetask struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	StartTime   string `json:"start-time"`
	Alarm       bool   `json:"alarm"`
}
