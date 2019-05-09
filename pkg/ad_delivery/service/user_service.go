package service

import (
	"ads/pkg/common"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

var GUserService *UserService

type AdUser struct {
	MongoId  bson.ObjectId `bson:"_id"`
	Username string        `bson:"username"`
	Password string        `bson:"password"`
	//Token      string        `json:"token" bson:"token"`
	Status     int   `bson:"status"`
	CreateTime int64 `bson:"create_time"`
	UpdateTime int64 `bson:"update_time"`
}

type UserService struct {
	Collection *mgo.Collection
}

// user status
const (
	UserValid   = 1
	UserInvalid = 0
)

func (s *UserService) CreateUser(username string, password string) (id string, err error) {
	adUser := &AdUser{}
	adUser.Username = username
	adUser.Password = common.Md5(password)
	adUser.MongoId = bson.NewObjectId()
	adUser.CreateTime = time.Now().Unix()
	adUser.UpdateTime = time.Now().Unix()
	adUser.Status = UserValid
	err = s.Collection.Insert(adUser)
	return adUser.MongoId.Hex(), err
}
