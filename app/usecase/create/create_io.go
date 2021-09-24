package create

import "time"

type CreateInput struct {
	Title     string
	Userid    string
	Context   string
	Starttime time.Time
	Endtime   time.Time
	Timeup    bool
}

type CreateOutput struct {
	ID string
}
