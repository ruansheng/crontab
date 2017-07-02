package models

import (
	"fmt"
	"web/libs"

	"gopkg.in/mgo.v2/bson"
)

type MachineInfo struct {
	Id       bson.ObjectId `bson:"_id"`
	Name     string
	Hostname string
	Ip       string
	Account  string
}

func GetMachines() []MachineInfo {
	c, session, err := libs.GetMongoClient("machine")
	defer session.Close()
	if err != nil {
		panic(err)
	}
	results := make([]MachineInfo, 2)
	err = c.Find(bson.M{}).All(&results)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(results)
	return results
}

func AddMachine(name string, hostname string, ip string, account string) bool {
	fmt.Println("AddMachine info:", name, hostname, ip, account)

	c, session, err := libs.GetMongoClient("machine")
	defer session.Close()
	if err != nil {
		panic(err)
	}

	info := MachineInfo{bson.NewObjectId(), name, hostname, ip, account}
	err = c.Insert(&info)
	if err != nil {
		fmt.Println(err)
	}
	return true
}

func EditMachine(dataid string, name string, hostname string, ip string, account string) bool {
	fmt.Println("EditMachine info:", dataid, name, hostname, ip, account)

	c, session, err := libs.GetMongoClient("machine")
	defer session.Close()
	if err != nil {
		panic(err)
	}

	query := bson.M{"_id": bson.ObjectIdHex(dataid)}
	data := bson.M{"$set": bson.M{"name": name, "hostname": hostname, "ip": ip, "account": account}}
	err = c.Update(query, data)
	if err != nil {
		fmt.Println(err)
	}
	return true
}

func RemoveMachine(dataid string) bool {
	fmt.Println("RemoveTask info:", dataid)

	c, session, err := libs.GetMongoClient("machine")
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
