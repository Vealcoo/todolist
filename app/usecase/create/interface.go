package create

type CreateUsecase interface {
	Create(input *CreateInput) (*CreateOutput, error)
}
