package mongo

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Users Структура пользователя
type Users struct {
	UserID        string `bson:"userID"`
	Name          string `bson:"name"`
	EncryptedSeed string `bson:"encryptedSeed"`
	Address       string `bson:"address"`
}

// ConnectToMongo mongo connection
func ConnectToMongo() *mgo.Session {
	user := "erage"
	password := "doBH8993nnjdoBH8993nnj"
	uri := "mongodb://" + user + ":" + password + "@51.144.89.99:27017"

	session, err := mgo.Dial(uri)
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	return session
}

// CloseMongoConnection mongo close connection
func CloseMongoConnection(session *mgo.Session) {
	session.Close()
}

// AddUser Добавление пользователя
func AddUser(openSession *mgo.Session, userID string, name string, encryptedSeed string, address string) {
	session := openSession.Copy()
	defer CloseMongoConnection(session)

	c := session.DB("admin").C("users")

	//var foundationNullArray []investInFoundation

	err := c.Insert(&Users{UserID: userID, Name: name, EncryptedSeed: encryptedSeed, Address: address})

	if err != nil {
		log.Fatal(err)
	}
}

// FindAllUsers Поиск всех users
func FindAllUsers(openSession *mgo.Session) []Users {
	session := openSession.Copy()
	defer CloseMongoConnection(session)

	c := session.DB("admin").C("users")

	var results []Users
	err := c.Find(bson.M{}).All(&results)

	if err != nil {
		log.Fatal(err)
	}

	return results
}

// FindUser Поиск конкретного пользователя
func FindUser(openSession *mgo.Session, userid string) Users {
	session := openSession.Copy()
	defer CloseMongoConnection(session)

	c := session.DB("admin").C("users")

	var results Users
	err := c.Find(bson.M{"userID": userid}).One(&results)

	if err != nil {
		log.Fatal(err)
	}

	return results
}
