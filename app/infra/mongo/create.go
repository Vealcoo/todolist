package mongo

import (
	"errors"
	"todolist/app/domain/model"

	"gopkg.in/mgo.v2/bson"
)

func (c *connect) Create(data *model.ListInfo) (string, error) {
	coll := c.session.DB(DB).C(C)
	info, err := coll.Upsert(nil, data)
	if err != nil {
		return "", errors.New("create list failed")
	}
	if info.UpsertedId != nil {
		x := info.UpsertedId.(bson.ObjectId)
		data.ID = string(x)
	}
	return data.ID, nil
}
