package mongo

import (
	"todolist/app/domain/repository"

	"gopkg.in/mgo.v2"
)

const (
	DB = "tyr-project"
	C  = "list"
)

type connect struct {
	session *mgo.Session
}

func NewRepository(m *mgo.Session) repository.ListRepository {
	return &connect{
		session: m,
	}
}
