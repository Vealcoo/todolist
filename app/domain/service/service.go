package service

import (
	"time"
	"todolist/app/domain/model"
)

type listService struct {
}

func (l *listService) NewListInfo(id, title, context, starttime, endtime, timeup string) (*model.ListInfo, error) {
	newInfo := new(model.ListInfo)

	newInfo.ListTitle = func() string {
		if title == "" {
			return "[草稿]"
		}
		return title
	}()
	start, _ := time.Parse(time.RFC3339, starttime)
	end, _ := time.Parse(time.RFC3339, endtime)
	newInfo.UserID = id
	newInfo.ListContext = context
	newInfo.StartTime = start
	newInfo.EndTime = end

	return newInfo, nil
}

func (l *listService) NewList(list *model.ListInfo) (*model.ListInfo, error) {
	newList := list
	return newList, nil
}
