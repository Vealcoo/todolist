package display

import (
	"todolist/app/domain/repository"
)

type displayUsecase struct {
	listRepo repository.ListRepository
}

func NewDisplayUsecase(listRepo repository.ListRepository) DisplayUsecase {
	return &displayUsecase{
		listRepo: listRepo,
	}
}

func (u *displayUsecase) Display(input *DisplayInput) (*DispalyOutput, error) {
	output := new(DispalyOutput)
	output.Display, _ = u.listRepo.Display(input.Userid)

	return output, nil
}
