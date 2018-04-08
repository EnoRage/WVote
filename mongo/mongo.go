package mongo

import (
	"log"
	"time"

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

// Voters Структура голосующих
type Voters struct {
	Num     int    `bson:"num"`
	Address string `bson:"address"`
	Vote    bool   `bson:"vote"`
}

// Votes Структура голосующих
type Votes struct {
	Num               int       `bson:"num"`
	Description       string    `bson:"description"`
	ApprovedAddresses []string  `bson:"approvedAddresses"`
	StartTime         time.Time `bson:"startTime"`
	EndTime           time.Time `bson:"endTime"`
	End               bool      `bson:"end"`
	ValidatorsAddress string    `bson:"validatorsAddress"`
}

// ConnectToMongo mongo connection
func ConnectToMongo() *mgo.Session {
	// user := "erage"
	// password := "doBH8993nnjdoBH8993nnj"
	// uri :=

	session, err := mgo.Dial("mongodb://erage:doBH8993nnjdoBH8993nnj@51.144.89.99:27017")
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

	c := session.DB("unblock").C("users")

	err := c.Insert(&Users{UserID: userID, Name: name, EncryptedSeed: encryptedSeed, Address: address})

	if err != nil {
		log.Fatal(err)
	}
}

// FindAllUsers Поиск всех users
func FindAllUsers(openSession *mgo.Session) []Users {
	session := openSession.Copy()
	defer CloseMongoConnection(session)

	c := session.DB("unblock").C("users")

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

	c := session.DB("unblock").C("users")

	var results Users
	c.Find(bson.M{"userID": userid}).One(&results)

	// if err != nil {
	// 	log.Fatal(err)
	// }
	return results
}

// FindUserByAddress Поиск конкретного пользователя по адресу
func FindUserByAddress(openSession *mgo.Session, address string) Users {
	session := openSession.Copy()
	defer CloseMongoConnection(session)

	c := session.DB("unblock").C("users")

	var results Users
	c.Find(bson.M{"address": address}).One(&results)

	// if err != nil {
	// 	log.Fatal(err)
	// }
	return results
}

// FindAllVotes Поиск всех голосований
func FindAllVotes(openSession *mgo.Session) []Votes {
	session := openSession.Copy()
	defer CloseMongoConnection(session)

	c := session.DB("unblock").C("votes")

	var results []Votes
	err := c.Find(bson.M{}).All(&results)

	if err != nil {
		log.Fatal(err)
	}

	return results
}

// FindVoteByNum находит конкретное голосование
func FindVoteByNum(openSession *mgo.Session, num int) Votes {
	session := openSession.Copy()
	defer CloseMongoConnection(session)

	c := session.DB("unblock").C("votes")

	var results Votes
	err := c.Find(bson.M{"num": num}).One(&results)

	if err != nil {
		log.Fatal(err)
	}

	return results
}

// FindAllVoters Поиск всех users
func FindAllVoters(openSession *mgo.Session) []Voters {
	session := openSession.Copy()
	defer CloseMongoConnection(session)

	c := session.DB("unblock").C("voters")

	var results []Voters
	err := c.Find(bson.M{}).All(&results)

	if err != nil {
		log.Fatal(err)
	}

	return results
}

// FindAllVotersByNum Поиск всех голосующих по номеру
func FindAllVotersByNum(openSession *mgo.Session, num int) []Voters {
	session := openSession.Copy()
	defer CloseMongoConnection(session)

	c := session.DB("unblock").C("voters")

	var results []Voters
	err := c.Find(bson.M{"num": num}).All(&results)

	if err != nil {
		log.Fatal(err)
	}

	return results
}
