package service

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

var GAdUnitService *AdUnitService

type AdUint struct {
	MongoId      bson.ObjectId `bson:"_id"`
	PlanId       string        `bson:"plan_id"`
	Name         string        `bson:"name"`
	PositionType int           `bson:"position_type"`
	Budget       int           `bson:"budget"`
	CreateTime   int64         `bson:"create_time"`
	UpdateTime   int64         `bson:"update_time"`
}


type AdUnitKeyword struct {
	MongoId bson.ObjectId `bson:"_id"`
	UnitId  string        `bson:"unit_id"`
	Keyword string        `bson:"keyword"`
}

type AdUnitDistrict struct {
	MongoId bson.ObjectId `bson:"_id"`
	UnitId  string        `bson:"unit_id"`
	State   string        `bson:"state"`
	City    string        `bson:"city"`
}

type AdUnitInterest struct {
	MongoId bson.ObjectId `bson:"_id"`
	UnitId  string        `bson:"unit_id"`
	Tag     string        `bson:"tag"`
}

type InnovationUnit struct {
	MongoId        bson.ObjectId `bson:"_id"`
	AdInnovationId string        `bson:"ad_innovation_id"`
	AdUintId       string        `bson:"ad_uint_id"`
}

type AdUnitService struct {
	AdUintCollection           *mgo.Collection
	AdUintKeywordCollection    *mgo.Collection
	AdUintDistrictCollection   *mgo.Collection
	AdUintInterestCollection   *mgo.Collection
	AdUnitInnovationCollection *mgo.Collection
}

func (s *AdUnitService) CreateAdUnit(adUint *AdUint) (id string, err error) {
	adUint.MongoId = bson.NewObjectId()
	adUint.CreateTime = time.Now().Unix()
	adUint.UpdateTime = time.Now().Unix()
	err = s.AdUintCollection.Insert(adUint)
	return adUint.MongoId.Hex(), err
}

func (s *AdUnitService) CreateKeyword(keyword *AdUnitKeyword) (id string, err error) {
	keyword.MongoId = bson.NewObjectId()
	err = s.AdUintKeywordCollection.Insert(keyword)
	return keyword.MongoId.Hex(), err
}

func (s *AdUnitService) CreateDistrict(district *AdUnitDistrict) (id string, err error) {
	district.MongoId = bson.NewObjectId()
	err = s.AdUintDistrictCollection.Insert(district)
	return district.MongoId.Hex(), err
}

func (s *AdUnitService) CreateInterest(interest *AdUnitInterest) (id string, err error) {
	interest.MongoId = bson.NewObjectId()
	err = s.AdUintInterestCollection.Insert(interest)
	return interest.MongoId.Hex(), err
}

func (s *AdUnitService) CreateUnitInnovation(iu *InnovationUnit) (id string, err error) {
	iu.MongoId = bson.NewObjectId()
	err = s.AdUnitInnovationCollection.Insert(iu)
	return iu.MongoId.Hex(), err
}
