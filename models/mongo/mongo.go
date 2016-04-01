package mongo

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
)

var session *mgo.Session

// Conn returns a mongo session
func Conn() *mgo.Session {
	return session.Copy()
}

func init() {
	url := beego.AppConfig.String("mongodb_url")

	sess, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}

	session = sess
	session.SetMode(mgo.Monotonic, true)
}
