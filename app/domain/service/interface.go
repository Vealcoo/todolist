package service

import (
	"todolist/app/domain/model"

	"github.com/gin-gonic/gin"
)

// ListService - List Domain Service
type ListService interface {
	// NewList - 建立一筆List
	NewList()
	// DeleteList - 刪除一筆List
	DeleteList(ctx *gin.Context, list *model.ListInfo) error
	// UpdateList() - 新增一筆List
	UpdateList()
	// DisplayList() - 更新一筆List
	DisplayList()
}
