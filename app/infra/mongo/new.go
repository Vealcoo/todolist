package mongo

import (
	"gopkg.in/mgo.v2"
)

type connect struct {
	session *mgo.Session
}

const (
	DB = "tyr-project"
	C  = "list"
)
