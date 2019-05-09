package service

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

var GAdPlanService *AdPlanService

type AdPlan struct {
	MongoId    bson.ObjectId `bson:"_id"`
	UserId     string        `bson:"user_id"`
	Name       string        `bson:"name"`
	StartTime  int64         `bson:"start_time"`
	EndTime    int64         `bson:"end_time"`
	CreateTime int64         `bson:"create_time"`
	UpdateTime int64         `bson:"update_time"`
}

type AdPlanService struct {
	Collection *mgo.Collection
}

func (s *AdPlanService) CreateAdPlan(adPlan *AdPlan) (id string, err error) {
	adPlan.MongoId = bson.NewObjectId()
	adPlan.CreateTime = time.Now().Unix()
	adPlan.UpdateTime = time.Now().Unix()
	err = s.Collection.Insert(adPlan)
	return adPlan.MongoId.Hex(), err
}

func (s *AdPlanService) GetAdPlans(userId string) (plans []AdPlan, err error) {
	plans = make([]AdPlan, 0)
	err = s.Collection.Find(bson.M{"user_id": userId}).All(&plans)
	return
}

func (s *AdPlanService) UpdateAdPlan(adPlan *AdPlan) (id string, err error) {
	err = s.Collection.UpdateId(adPlan.MongoId, bson.M{"$set": bson.M{"name": adPlan.Name, "start_time": adPlan.StartTime, "end_time": adPlan.EndTime}})
	return adPlan.MongoId.Hex(), err
}

func (s *AdPlanService) DeleteAdPlan(id string) (err error) {
	err = s.Collection.RemoveId(bson.ObjectIdHex(id))
	return err
}
