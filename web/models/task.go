package models

import (
	"fmt"
	"web/libs"

	"gopkg.in/mgo.v2/bson"
)

type TaskInfo struct {
	Id        bson.ObjectId `bson:"_id"`
	Name      string
	Exectime  string
	Cmd       string
	Machineid string
}

func GetTasks() []TaskInfo {
	c, session, err := libs.GetMongoClient("task")
	defer session.Close()
	if err != nil {
		panic(err)
	}

	results := make([]TaskInfo, 2)
	err = c.Find(bson.M{}).All(&results)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(results)
	return results
}

func AddTask(name string, exectime string, cmd string, machineid string) bool {
	fmt.Println("AddTask info:", name, cmd, exectime, machineid)

	c, session, err := libs.GetMongoClient("task")
	defer session.Close()
	if err != nil {
		panic(err)
	}

	info := TaskInfo{bson.NewObjectId(), name, exectime, cmd, machineid}
	err = c.Insert(&info)
	if err != nil {
		fmt.Println(err)
	}
	return true
}

func EditTask(dataid string, name string, exectime string, cmd string, machineid string) bool {
	fmt.Println("EditTask info:", dataid, name, cmd, exectime, machineid)

	c, session, err := libs.GetMongoClient("task")
	defer session.Close()
	if err != nil {
		panic(err)
	}

	query := bson.M{"_id": bson.ObjectIdHex(dataid)}
	data := bson.M{"$set": bson.M{"name": name, "exectime": exectime, "cmd": cmd, "machineid": machineid}}
	err = c.Update(query, data)
	if err != nil {
		fmt.Println(err)
	}
	return true
}

func RemoveTask(dataid string) bool {
	fmt.Println("RemoveTask info:", dataid)

	c, session, err := libs.GetMongoClient("task")
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
