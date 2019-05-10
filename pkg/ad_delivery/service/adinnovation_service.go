package service

import (
	"cloud.google.com/go/firestore"
	"context"
	"time"
)

var GAdInnovationService *AdInnovationService

type AdInnovation struct {
	Id         string
	UserId     string
	Name       string
	Type       int
	Material   int
	Height     int
	Width      int
	Size       int
	Duration   int
	Url        string
	CreateTime int64
	UpdateTime int64
}

type AdInnovationService struct {
	Collection *firestore.CollectionRef
}

func (s *AdInnovationService) CreateAdInnovation(inno *AdInnovation) (id string, err error) {
	inno.CreateTime = time.Now().Unix()
	inno.UpdateTime = time.Now().Unix()
	d, _, err := s.Collection.Add(context.Background(), inno)
	inno.Id = d.ID

	SaveAdInnovationIndex(inno)
	return inno.Id, err
}
