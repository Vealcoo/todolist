package display

import (
	"todolist/app/domain/model"
)

type DisplayInput struct {
	Userid string
}

type DisplayOutput struct {
	Display []model.ListInfo
}
