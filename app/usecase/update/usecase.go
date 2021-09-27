package update

import (
	"todolist/app/domain/repository"
	"todolist/app/domain/service"
)

type updateUsecase struct {
	listService service.ListService
	listRepo    repository.ListRepository
}

func NewUpdateUsecase(listService service.ListService, listRepo repository.ListRepository) UpdateUsecase {
	return &updateUsecase{
		listService: listService,
		listRepo:    listRepo,
	}
}

func (u *updateUsecase) Update(input *UpdateInput) error {
	info, err := u.listService.NewListInfo(input.Userid, input.Title, input.Context, input.Starttime, input.Endtime, input.Timeup)
	if err != nil {
		return err
	}

	list, err := u.listService.NewList(info)
	if err != nil {
		return err
	}

	if err := u.listRepo.Update(list, input.Listid); err != nil {
		return err
	}
	return nil
}
