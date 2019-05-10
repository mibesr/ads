package service

import (
	"cloud.google.com/go/firestore"
	"context"
	"google.golang.org/api/iterator"
	"time"
)

var GAdPlanService *AdPlanService

type AdPlan struct {
	Id         string `json:"id"`
	UserId     string `json:"user_id"`
	Name       string `json:"name"`
	StartTime  int64  `json:"start_time"`
	EndTime    int64  `json:"end_time"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}

type AdPlanService struct {
	Collection *firestore.CollectionRef
}

func (s *AdPlanService) CreateAdPlan(adPlan *AdPlan) (id string, err error) {

	adPlan.CreateTime = time.Now().Unix()
	adPlan.UpdateTime = time.Now().Unix()
	d, _, err := s.Collection.Add(context.Background(), adPlan)
	adPlan.Id = d.ID
	SaveAdPlanIndex(adPlan)
	return adPlan.Id, err
}

func (s *AdPlanService) GetAdPlans(userId string) (plans []AdPlan, err error) {
	plans = make([]AdPlan, 0)
	it := s.Collection.Where("UserId", "==", userId).Documents(context.Background())
	for {
		doc, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			adPlan := AdPlan{}
			doc.DataTo(&adPlan)
			adPlan.Id = doc.Ref.ID
		}
	}
	return plans, nil
}

func (s *AdPlanService) GetAdPlanById(planId string) (adPlan AdPlan, err error) {
	adPlan = AdPlan{}
	d, err := s.Collection.Doc(planId).Get(context.Background())
	if d != nil {
		d.DataTo(&adPlan)
	}
	return adPlan, err
}

func (s *AdPlanService) DeleteAdPlan(id string) (err error) {
	_, err = s.Collection.Doc(id).Delete(context.Background())
	DeleteAdPlanIndex(id)
	return err
}
