package display

import "todolist/app/domain/model"

type DisplayUsecase interface {
	Display(input *DisplayInput) ([]model.ListInfo, error)
}
