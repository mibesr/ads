package service

import (
	"ads/pkg/common"
	"cloud.google.com/go/firestore"
	"context"
	"time"
)

var GUserService *UserService

type AdUser struct {
	Id         string
	Username   string
	Password   string
	CreateTime int64
	UpdateTime int64
}

type UserService struct {
	Collection *firestore.CollectionRef
}

func (s *UserService) CreateUser(username string, password string) (id string, err error) {
	adUser := &AdUser{}
	adUser.Username = username
	adUser.Password = common.Md5(password)
	adUser.CreateTime = time.Now().Unix()
	adUser.UpdateTime = time.Now().Unix()
	d, _, err := s.Collection.Add(context.Background(), adUser)
	adUser.Id = d.ID
	return adUser.Id, err
}

func (s *UserService) GetUserByUsername(username string) (user AdUser, err error) {
	adUser := AdUser{}
	it := s.Collection.Where("Username", "==", username).Documents(context.Background())
	doc, err := it.Next()
	if doc != nil {
		doc.DataTo(&adUser)
		adUser.Id = doc.Ref.ID
		return adUser, nil
	}
	return adUser, err
}
