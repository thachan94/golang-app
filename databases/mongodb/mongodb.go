package mongodb

import (
	"time"

	"gopkg.in/mgo.v2"
)

type DBConn struct {
	session *mgo.Session
}

func NewConn(host string) (conn *DBConn) {
	session, err := mgo.Dial(host)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	session.SetSocketTimeout(1 * time.Hour)
	conn = &DBConn{session}
	return conn
}

func (conn *DBConn) Use(dbName, collectionName string) (collection *mgo.Collection) {
	return conn.session.DB(dbName).C(collectionName)
}

func (conn *DBConn) Close() {
	conn.session.Close()
	return
}
