package mongo

import (
	"errors"
	"todolist/app/domain/model"
)

func (c *connect) Create(data *model.ListInfo) error {
	coll := c.session.DB(DB).C(C)
	err := coll.Insert(data)
	if err != nil {
		return errors.New("create list failed")
	}
	return nil
}
