package repository

import "todolist/app/domain/model"

type ListRepository interface {
	Create(data *model.ListInfo) error

	Delete(listid string) error

	Update(data *model.ListInfo, listid string) error

	Display(userid string) ([]model.ListInfo, error)
}
