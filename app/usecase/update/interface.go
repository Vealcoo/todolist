package update

type UpdateUsecase interface {
	Update(input *UpdateInput) error
}
