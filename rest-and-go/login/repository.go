package login

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Repository struct{}

const (
	mongoUrl   = "localhost:27017"
	collection = "user_details"
	db_name    = "DataX"
)

// Check user_id in database

func (r Repository) checkEmailId(email_id string) Person {
	session, err := mgo.Dial(mongoUrl)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	c := session.DB(db_name).C(collection)
	res := Person{}
	err = c.Find(bson.M{"email_id": email_id}).One(&res)
	return res
}

func (r Repository) InsertCredential(person Person) bool {
	session, err := mgo.Dial(mongoUrl)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	c := session.DB(db_name).C(collection)
	err = c.Insert(person)
	return true
}

func (r Repository) Update(person Person) bool {
	session, err := mgo.Dial(mongoUrl)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	c := session.DB(db_name).C(collection)
	err = c.Update(bson.M{"email_id": person.Email_id}, bson.M{"$set": person})
	return true
}
