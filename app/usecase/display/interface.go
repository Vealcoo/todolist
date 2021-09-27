package display

type DisplayUsecase interface {
	Display(input *DisplayInput) (*DisplayOutput, error)
}
