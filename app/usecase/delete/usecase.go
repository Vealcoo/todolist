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

func (u *deleteUsecase) Delete(input *DeleteInput) error {
	if err := u.listRepo.Delete(input.Listid); err != nil {
		return err
	}
	return nil
}
