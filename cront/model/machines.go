package model

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MachineInfo struct {
	Id       bson.ObjectId `bson:"_id"`
	Name     string
	Hostname string
	Ip       string
	Account  string
}

func GetMachineData() []MachineInfo {
	session, err := mgo.Dial("mongodb://127.0.0.1:27017/demo")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("demo").C("machine")
	results := make([]MachineInfo, 2)
	//err = c.Find(bson.M{"name": "jh"}).All(&results)
	err = c.Find(bson.M{}).All(&results)
	if err != nil {
		fmt.Println(err)
	}
	return results
}

func getMachinesMap() map[string]MachineInfo {
	machines := GetMachineData()

	list := make(map[string]MachineInfo)

	for _, result := range machines {
		list[result.Id.Hex()] = result
	}

	return list
}
