package mongo

import (
	"errors"
	"todolist/app/domain/model"

	"gopkg.in/mgo.v2/bson"
)

func (c *connect) Display(userid string) ([]model.ListInfo, error) {
	coll := c.session.DB(DB).C(C)
	var result []model.ListInfo
	err := coll.Find(bson.M{"userid": userid}).All(&result)
	if err != nil {
		return nil, errors.New("can not find list")
	}
	return result, nil
}
