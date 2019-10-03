package userModel

import (
	"app/databases/mongodb"

	"gopkg.in/mgo.v2/bson"
)

const (
	Host           = "localhost:27017"
	DBMgo          = "v4update"
	DBUser         = `v4`
	DBPassword     = `s3JR#~7g#5YNM_o`
	UserCollection = "user"
)

type User struct {
	ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	UserName string        `json:"username" bson:"username"`
	Password string        `json:"password" bson:"password"`
	Token    string        `json:"token" bson"token"`
}
type UserModel struct{}

var dbConn = mongodb.NewConn("mongodb://" + DBUser + ":" + DBPassword + "@" + Host + "/" + DBMgo)

func (m *UserModel) Login(user *User) error {
	collection := dbConn.Use(DBMgo, "user")
	err := collection.Find(bson.M{"username": user.UserName}).One(&user)
	return err
}

func (m *UserModel) Signup(user *User) error {
	collection := dbConn.Use(DBMgo, "user")
	err := collection.Insert(bson.M{"username": user.UserName, "password": user.Password})
	return err
}

func (m *UserModel) CheckUserExist(user *User) (int, error) {
	collection := dbConn.Use(DBMgo, "user")
	count, err := collection.Find(bson.M{"username": user.UserName}).Count()
	return count, err
}

func (m *UserModel) GetUserInfo(user *User) (User, error) {
	collection := dbConn.Use(DBMgo, "user")
	err := collection.Find(bson.M{"_id": user.ID}).One(&user)
	return *user, err
}
