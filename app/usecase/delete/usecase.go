package delete

import (
	"todolist/app/domain/repository"
)

type deleteUsecase struct {
	listRepo repository.ListRepository
}

func NewDeleteUsecase(listRepo repository.ListRepository) DeleteUsecase {
	return &deleteUsecase{
		listRepo: listRepo,
	}
}
