package service

import (
	"ads/pkg/common"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"os"
	"strconv"
	"testing"
	"time"
)

func initMongo() error {
	pwd, _ := os.Getwd()
	conf, err := common.InitConfig(pwd + "/../../config.json")
	if err != nil {
		log.Fatal(err)
	}
	common.GConfig = &conf

	session, err := mgo.DialWithTimeout(common.GConfig.MongoDBUri, time.Duration(common.GConfig.MongoDBTimeout)*time.Millisecond)
	if err != nil {
		return err
	}
	// init services
	GUserService = &UserService{
		Collection: session.DB(common.GConfig.DBName).C("ad_user"),
	}
	GAdPlanService = &AdPlanService{
		Collection: session.DB(common.GConfig.DBName).C("ad_plan"),
	}
	GAdUnitService = &AdUnitService{
		AdUintCollection:           session.DB(common.GConfig.DBName).C("ad_unit"),
		AdUintDistrictCollection:   session.DB(common.GConfig.DBName).C("ad_unit_district"),
		AdUintKeywordCollection:    session.DB(common.GConfig.DBName).C("ad_unit_keyword"),
		AdUintInterestCollection:   session.DB(common.GConfig.DBName).C("ad_unit_interest"),
		AdUnitInnovationCollection: session.DB(common.GConfig.DBName).C("ad_unit_innovation"),
	}
	GAdInnovationService = &AdInnovationService{
		Collection: session.DB(common.GConfig.DBName).C("ad_innovation"),
	}
	return err
}

func TestUserService(t *testing.T) {
	err := initMongo()
	if err != nil {
		log.Fatal(err)
	}
	id, _ := GUserService.CreateUser("test", "123")
	user := AdUser{}
	err = GUserService.Collection.FindId(bson.ObjectIdHex(id)).One(&user)
	if err != nil {
		t.Fatal(err)
	}
	if user.Username != "test" {
		t.Fatal("Except username = test, actual username = " + user.Username)
	} else {
		t.Log("Username: " + user.Username + ", Password: " + user.Password)
	}
}

func TestAdPlanService(t *testing.T) {
	err := initMongo()
	if err != nil {
		log.Fatal(err)
	}
	user, err := GUserService.GetUserByUsername("test")
	if err != nil {
		log.Fatal(err)
	}
	adPlan := &AdPlan{
		UserId:    user.MongoId.Hex(),
		Name:      "test ad plan",
		StartTime: 100000,
		EndTime:   200000,
	}
	id, _ := GAdPlanService.CreateAdPlan(adPlan)
	plans, _ := GAdPlanService.GetAdPlans(user.MongoId.Hex())
	t.Log("got " + strconv.Itoa(len(plans)) + " plans")
	for _, plan := range plans {
		if plan.MongoId.Hex() == id {
			t.Log("plan id = " + id + " got")
			break
		}
	}
	adPlan.Name = "test ad plan 1"
	id, _ = GAdPlanService.UpdateAdPlan(adPlan)
	t.Log("plan id = " + id + " updated")
	GAdPlanService.DeleteAdPlan(id)
	t.Log("plan id = " + id + " deleted")
	plans, _ = GAdPlanService.GetAdPlans(user.MongoId.Hex())
	t.Log("got " + strconv.Itoa(len(plans)) + " plans")
}

func TestAdUnitService(t *testing.T) {
	err := initMongo()
	if err != nil {
		log.Fatal(err)
	}
	user, err := GUserService.GetUserByUsername("test")
	adPlan := &AdPlan{
		UserId:    user.MongoId.Hex(),
		Name:      "test ad plan",
		StartTime: 100000,
		EndTime:   200000,
	}
	id, _ := GAdPlanService.CreateAdPlan(adPlan)
	adUnit := AdUint{
		PlanId:       id,
		Name:         "test ad unit",
		PositionType: common.Video,
		Budget:       1000,
	}
	id, _ = GAdUnitService.CreateAdUnit(&adUnit)

	it1 := AdUnitInterest{
		UnitId: id,
		Tag:    "swimming",
	}
	it2 := AdUnitInterest{
		UnitId: id,
		Tag:    "game",
	}
	it3 := AdUnitInterest{
		UnitId: id,
		Tag:    "guitar",
	}
	GAdUnitService.CreateInterest(&it1)
	GAdUnitService.CreateInterest(&it2)
	GAdUnitService.CreateInterest(&it3)

	kw1 := AdUnitKeyword{
		UnitId:  id,
		Keyword: "sport",
	}
	kw2 := AdUnitKeyword{
		UnitId:  id,
		Keyword: "music",
	}

	GAdUnitService.CreateKeyword(&kw1)
	GAdUnitService.CreateKeyword(&kw2)

	dis1 := AdUnitDistrict{
		UnitId: id,
		State:  "NY",
		City:   "New York",
	}
	dis2 := AdUnitDistrict{
		UnitId: id,
		State:  "CA",
		City:   "LA",
	}

	GAdUnitService.CreateDistrict(&dis1)
	GAdUnitService.CreateDistrict(&dis2)

	inno := AdInnovation{
		UserId:   user.MongoId.Hex(),
		Type:     1,
		Material: 1,
		Height:   100,
		Width:    100,
		Size:     1024,
		Duration: 10,
		Url:      "http://www.google.com",
	}

	GAdInnovationService.CreateAdInnovation(&inno)
	iu := InnovationUnit{
		AdInnovationId: inno.MongoId.Hex(),
		AdUintId:       adUnit.MongoId.Hex(),
	}
	GAdUnitService.CreateUnitInnovation(&iu)
}
