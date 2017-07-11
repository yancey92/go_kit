package dbkit

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

var sslUrl string = "mongodb://gumpmongodev:y04eTIOHkDTXsTMuXWNtlDtiWSasByuIQ5owemMbwdxJCNcUDurJMDhXO6zXzzYDjILolp3yLT31Dk9ETuKSJQ==@gumpmongodev.documents.azure.cn:10250/"

type Person struct {
	Name  string
	Phone string
}

func TestMongoUpsertDoc(t *testing.T) {
	//InitMongoDB("mongodb://localhost:27017", "gumpcome")
	InitMongoDBWithSSL(sslUrl, "gumpcome")
	_, isOk, err := MongoUpsertDoc(&MongoSearch{
		Collection: "people",
		Key:        "phone",
		Value:      "123456789",
	}, &Person{"guanyu", "123456789"})
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(isOk)
}

func TestMongoFindDoc(t *testing.T) {
	//InitMongoDB("mongodb://localhost:27017", "gumpcome")
	InitMongoDBWithSSL(sslUrl, "gumpcome")
	result := Person{}
	MongoFindDoc("people", func(c *mgo.Collection) {
		c.Find(bson.M{"phone": "123456789"}).One(&result)
	})
	fmt.Println(result)
}

func TestMongoRemoveAllDoc(t *testing.T) {
	InitMongoDB("mongodb://localhost:27017", "gumpcome")
	isOk, err := MongoRemoveAllDoc(&MongoSearch{
		Collection: "people",
		Key:        "phone",
		Value:      "123456789",
	})
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(isOk)
}
