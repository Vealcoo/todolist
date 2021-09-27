package service

import (
	"todolist/app/domain/model"
)

// ListService - List Domain Service
type ListService interface {
	// NewListInfo - 建立新的ListInfo or 更新的ListInfo
	NewListInfo(id, title, context, starttime, endtime string, timeup bool) (*model.ListInfo, error)
	// NewList - 建立一筆List
	NewList(list *model.ListInfo) (*model.ListInfo, error)
}
