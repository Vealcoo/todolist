package mongo

import (
	"errors"

	"gopkg.in/mgo.v2/bson"
)

func (c *connect) Delete(listid string) error {
	coll := c.session.DB(DB).C(C)

	selector := bson.M{"id:": bson.ObjectIdHex(listid)}
	err := coll.Remove(selector)
	if err != nil {
		return errors.New("delete list failed")
	}
	return nil
}
