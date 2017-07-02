package model

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TaskInfo struct {
	Name     string
	Cmd      string
	Exectime string
}

func GetTaskData() []TaskInfo {
	session, err := mgo.Dial("mongodb://127.0.0.1:27017/demo")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("demo").C("task")
	results := make([]TaskInfo, 2)
	//err = c.Find(bson.M{"name": "jh"}).All(&results)
	err = c.Find(bson.M{}).All(&results)
	if err != nil {
		fmt.Println(err)
	}
	return results
}
