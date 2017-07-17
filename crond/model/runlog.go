package model

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type RunLog struct {
	Id        bson.ObjectId `bson:"_id"`
	RequestId string
	Flag      int
	Time      int64
	Data      string
}

func AddRunLog(RequestId string, Flag int, Time int64, Data string) bool {
	fmt.Println("RunLog info:", RequestId, Flag, Time, Data)

	session, err := mgo.Dial("mongodb://127.0.0.1:27017/demo")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("demo").C("run_result")
	defer session.Close()
	if err != nil {
		panic(err)
	}

	info := RunLog{bson.NewObjectId(), RequestId, Flag, Time, Data}
	err = c.Insert(&info)
	if err != nil {
		fmt.Println(err)
	}
	return true
}
