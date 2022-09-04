package db

import (
	"time"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type Todo struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Description string `json:"description"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func RetrieveTodo(id string) (*Todo, error) {
	todo := Todo{ID: uuid.MustParse(id)}
	result := DB.First(&todo)

	if result.Error != nil {
		log.WithFields(log.Fields{
			"resource": "todos",
			"todo_id": todo.ID,
		}).Error(result.Error)
		return nil, result.Error
	}

	log.WithFields(log.Fields{
		"resource": "todos",
		"todo_id": todo.ID,
	}).Infof("Todo with ID: %v retrieved", todo.ID)

	return &todo, nil 
}

func ListTodos(filter *Todo, skip int, limit int) ([]Todo, error) {
	var todos []Todo

	result := DB.Where(filter).Limit(limit).Offset(skip).Find(&todos)

	if result.Error != nil {
		log.WithFields(log.Fields{
			"resource": "todos",
			"filter": filter,
		}).Error(result.Error)
		return nil, result.Error
	}

	log.WithFields(log.Fields{
		"resource": "todos",
		"filter": filter,
	}).Infof("Todos retrieved")

	return todos, nil
}