package create

import (
	"todolist/app/domain/repository"
	"todolist/app/domain/service"
)

type createUsecase struct {
	listService service.ListService
	listRepo    repository.ListRepository
}

func NewCreateUsecase(listService service.ListService, listRepo repository.ListRepository) CreateUsecase {
	return &createUsecase{
		listService: listService,
		listRepo:    listRepo,
	}
}

func (u *createUsecase) Create(input *CreateInput) (*CreateOutput, error) {
	output := new(CreateOutput)

	info, err := u.listService.NewListInfo(input.Userid, input.Title, input.Context, input.Starttime, input.Endtime, input.Timeup)
	if err != nil {
		return nil, err
	}

	list, err := u.listService.NewList(info)
	if err != nil {
		return nil, err
	}

	if err := u.listRepo.Create(list); err != nil {
		return nil, err
	}

	// if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
	// 	return c.JSON(http.StatusCreated, map[string]interface{}{
	// 		"id": oid.Hex(),
	// 	})
	// }
	return output, nil
}
