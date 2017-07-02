package models

import (
	"fmt"
	"web/libs"

	"gopkg.in/mgo.v2/bson"
)

type AccountInfo struct {
	Id          bson.ObjectId `bson:"_id"`
	Name        string
	Password    string
	Phonenumber string
	Email       string
}

func GetAccounts() []AccountInfo {
	c, session, err := libs.GetMongoClient("account")
	defer session.Close()
	if err != nil {
		panic(err)
	}

	results := make([]AccountInfo, 2)
	err = c.Find(bson.M{}).All(&results)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(results)
	return results
}

func AddAccount(name string, password string, phonenumber string, email string) bool {
	fmt.Println("AddAccount info:", name, phonenumber, email)

	c, session, err := libs.GetMongoClient("account")
	defer session.Close()
	if err != nil {
		panic(err)
	}

	info := AccountInfo{bson.NewObjectId(), name, password, phonenumber, email}
	err = c.Insert(&info)
	if err != nil {
		fmt.Println(err)
	}
	return true
}

func EditAccount(dataid string, name string, password string, phonenumber string, email string) bool {
	fmt.Println("EditAccount info:", dataid, name, password, phonenumber, email)

	c, session, err := libs.GetMongoClient("account")
	defer session.Close()
	if err != nil {
		panic(err)
	}

	query := bson.M{"_id": bson.ObjectIdHex(dataid)}
	if password != "" {
		data := bson.M{"$set": bson.M{"name": name, "password": password, "phonenumber": phonenumber, "email": email}}
		err = c.Update(query, data)
	} else {
		data := bson.M{"$set": bson.M{"name": name, "phonenumber": phonenumber, "email": email}}
		err = c.Update(query, data)
	}
	if err != nil {
		fmt.Println(err)
	}
	return true
}

func RemoveAccount(dataid string) bool {
	fmt.Println("RemoveAccount info:", dataid)

	c, session, err := libs.GetMongoClient("account")
	defer session.Close()
	if err != nil {
		panic(err)
	}

	id := bson.ObjectIdHex(dataid)
	_, err = c.RemoveAll(bson.M{"_id": id})
	if err != nil {
		fmt.Println(err)
	}
	return true
}
