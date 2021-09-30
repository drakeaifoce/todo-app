package service

import (
	"github.com/drakeaifoce/todo-app"
	"github.com/drakeaifoce/todo-app/pkg/repository"
)

type TodoListService struct {
	repo repository.ToDoList
}

func NewTodoListService(repo repository.ToDoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list todo.ToDoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAll(userId int) ([]todo.ToDoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetById(userId, listId int) (todo.ToDoList, error) {
	return s.repo.GetById(userId, listId)
}

func (s *TodoListService) Update(userId, listId int, input todo.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, listId, input)
}

func (s *TodoListService) DeleteList(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}
