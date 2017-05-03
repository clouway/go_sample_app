package persistence

import (
	"fmt"
	"time"

	"github.com/mgenov/myproject/domain"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const dbName = "bank"

func NewSessionStore(session *mgo.Session) domain.SessionStore {
	return &mongoSessionStore{session}
}

type mongoSessionStore struct {
	session *mgo.Session
}

func (m *mongoSessionStore) StartSession(u domain.User, time time.Time) (*domain.Session, error) {
	return nil, nil
}

func (m *mongoSessionStore) FindSessionAvailableAt(ID string, instant time.Time) (*domain.Session, bool, error) {
	session := m.session.Clone()
	defer session.Close()

	var res map[string]interface{}
	err := session.DB(dbName).C("sessions").Find(bson.M{"_id": ID}).One(&res)
	if err != nil {
		return nil, false, fmt.Errorf("could not retrieve session due: %v", err)
	}

	exp := res["exp"].(time.Time)

	if !exp.Before(instant) {
		return nil, false, nil
	}

	return &domain.Session{ID: res["_id"].(string), ExpiresAt: exp, UserID: res["userId"].(string)}, true, nil
}
