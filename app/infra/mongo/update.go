package mongo

import (
	"errors"
	"todolist/app/domain/model"

	"gopkg.in/mgo.v2/bson"
)

func (c *connect) Update(data *model.ListInfo, listid string) error {
	coll := c.session.DB(DB).C(C)

	selector := bson.M{"_id": bson.IsObjectIdHex((listid))}
	err := coll.Update(selector, data)
	if err != nil {
		return errors.New("list update failed")
	}
	return nil
}
