package service

import (
	"github.com/drakeaifoce/todo-app"
	"github.com/drakeaifoce/todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type ToDoList interface {
	Create(userId int, list todo.ToDoList) (int, error)
	GetAll(userId int) ([]todo.ToDoList, error)
	GetById(userId, listId int) (todo.ToDoList, error)
	Update(userId, listId int, input todo.UpdateListInput) error
	DeleteList(userId, listId int) error
}

type TodoItem interface {
	Create(userId, listId int, item todo.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todo.TodoItem, error)
	GetById(userId, itemId int) (todo.TodoItem, error)
	Update(userId, itemId int, input todo.UpdateItemInput) error
	Delete(userId, itemId int) error
}

type Service struct {
	Authorization
	ToDoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		ToDoList:      NewTodoListService(repos.ToDoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.ToDoList),
	}
}
