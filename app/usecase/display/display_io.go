package display

import (
	"todolist/app/domain/model"
)

type DisplayInput struct {
	Userid string
}

type DispalyOutput struct {
	Display []model.ListInfo
}
