package create

type CreateInput struct {
	Userid    string
	Title     string
	Context   string
	Starttime string
	Endtime   string
	Timeup    bool
}

type CreateOutput struct {
	Listid string
}
