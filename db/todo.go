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


func InsertTodo(todo *Todo) (*Todo, error) {
	result := DB.Create(todo)

	if result.Error != nil {
		log.WithFields(log.Fields{
			"resource": "todos",
		}).Error(result.Error)
		return nil, result.Error
	}

	log.WithFields(log.Fields{
		"resource": "todos",
		"todo_id": todo.ID,
	}).Infof("Todo with ID: %v retrieved", todo.ID)

	return todo, nil
}

func UpdateTodo(id string, updates *Todo) (*Todo, error) {
  todo := Todo{ID: uuid.MustParse(id)}
	result := DB.Model(&todo).Updates(updates)

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
	}).Infof("Todo with ID: %v updated", todo.ID)
	return &todo, nil
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