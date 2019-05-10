package service

import (
	"cloud.google.com/go/firestore"
	"context"
	"time"
)

var GAdUnitService *AdUnitService

type AdUint struct {
	Id           string `json:"id"`
	PlanId       string `json:"plan_id"`
	Name         string `json:"name"`
	PositionType int    `json:"position_type"`
	Budget       int    `json:"budget"`
	CreateTime   int64  `json:"create_time"`
	UpdateTime   int64  `json:"update_time"`
}

type AdUnitKeyword struct {
	Id      string `json:"id"`
	UnitId  string `json:"unit_id"`
	Keyword string `json:"keyword"`
}

type AdUnitDistrict struct {
	Id     string `json:"id"`
	UnitId string `json:"unit_id"`
	State  string `json:"state"`
	City   string `json:"city"`
}

type AdUnitInterest struct {
	Id     string `json:"id"`
	UnitId string `json:"unit_id"`
	Tag    string `json:"tag"`
}

type InnovationUnit struct {
	Id             string `json:"id"`
	AdInnovationId string `json:"ad_innovation_id"`
	AdUintId       string `json:"ad_uint_id"`
}

type AdUnitService struct {
	AdUintCollection           *firestore.CollectionRef
	AdUintKeywordCollection    *firestore.CollectionRef
	AdUintDistrictCollection   *firestore.CollectionRef
	AdUintInterestCollection   *firestore.CollectionRef
	AdUnitInnovationCollection *firestore.CollectionRef
}

func (s *AdUnitService) CreateAdUnit(adUint *AdUint) (id string, err error) {
	adUint.CreateTime = time.Now().Unix()
	adUint.UpdateTime = time.Now().Unix()
	d, _, err := s.AdUintCollection.Add(context.Background(), adUint)
	adUint.Id = d.ID

	SaveAdUnitIndex(adUint)
	return adUint.Id, err
}

func (s *AdUnitService) CreateKeyword(keyword *AdUnitKeyword) (id string, err error) {
	d, _, err := s.AdUintKeywordCollection.Add(context.Background(), keyword)
	keyword.Id = d.ID
	SaveKeywordIndex(keyword)
	return keyword.Id, err
}

func (s *AdUnitService) CreateDistrict(district *AdUnitDistrict) (id string, err error) {
	d, _, err := s.AdUintDistrictCollection.Add(context.Background(), district)
	district.Id = d.ID
	SaveDistrictIndex(district)
	return district.Id, err
}

func (s *AdUnitService) CreateInterest(interest *AdUnitInterest) (id string, err error) {
	d, _, err := s.AdUintInterestCollection.Add(context.Background(), interest)
	interest.Id = d.ID
	SaveInterestIndex(interest)
	return interest.Id, err
}

func (s *AdUnitService) CreateUnitInnovation(iu *InnovationUnit) (id string, err error) {
	d, _, err := s.AdUnitInnovationCollection.Add(context.Background(), iu)
	iu.Id = d.ID
	SaveUnitInnovationIndex(iu)
	return iu.Id, err
}
