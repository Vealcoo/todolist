package delete

type DeleteUsecase interface {
	Delete(input *DeleteInput) error
}
