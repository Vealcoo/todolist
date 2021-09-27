package model

import "time"

type ListInfo struct {
	ID          interface{} `bson:"_id"`
	UserID      string      `bson:"userid"`
	ListTitle   string      `bson:"listtitle"`
	ListContext string      `bson:"listcontext"`
	StartTime   time.Time   `bson:"starttime"`
	EndTime     time.Time   `bson:"endtime"`
	TimeUp      bool        `bson:"timeupmark"`
}
