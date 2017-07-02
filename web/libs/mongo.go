package libs

import (
	"fmt"

	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
)

func GetMongoClient(collection string) (*mgo.Collection, *mgo.Session, error) {
	host := beego.AppConfig.String("mongohost")
	port := beego.AppConfig.String("monogport")
	db := beego.AppConfig.String("mongodb")
	url := fmt.Sprintf("mongodb://%s:%s/%s", host, port, db)
	fmt.Println(url)
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, nil, err
	}
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(db).C(collection)
	return c, session, nil
}
