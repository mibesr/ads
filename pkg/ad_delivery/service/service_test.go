package service

import (
	"ads/pkg/common"
	"cloud.google.com/go/firestore"
	"cloud.google.com/go/pubsub"
	"context"
	"google.golang.org/api/option"
	"log"
	"os"
	"strconv"
	"testing"
)


func initMongo() error {
	pwd, _ := os.Getwd()
	conf, err := common.InitConfig(pwd + "/../../config.json")
	if err != nil {
		log.Fatal(err)
	}
	common.GConfig = &conf
	projectID := common.GConfig.ProjectId
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID, option.WithCredentialsFile(pwd+"/../../auth.json"))
	if err != nil {
		return err
	}

	GUserService = &UserService{
		Collection: client.Collection("ad_user"),
	}

	GAdPlanService = &AdPlanService{
		Collection: client.Collection("ad_plan"),
	}
	GAdUnitService = &AdUnitService{
		AdUintCollection:           client.Collection("ad_unit"),
		AdUintDistrictCollection:   client.Collection("ad_unit_district"),
		AdUintKeywordCollection:    client.Collection("ad_unit_keyword"),
		AdUintInterestCollection:   client.Collection("ad_unit_interest"),
		AdUnitInnovationCollection: client.Collection("ad_unit_innovation"),
	}
	GAdInnovationService = &AdInnovationService{
		Collection: client.Collection("ad_innovation"),
	}

	return nil
}

func initPub() error {
	pwd, _ := os.Getwd()
	projectID := common.GConfig.ProjectId
	cli, err := pubsub.NewClient(context.Background(), projectID, option.WithCredentialsFile(pwd+"/../../auth.json"))
	if err != nil {
		return err
	}
	PubClient = cli
	return nil
}

func TestUserService(t *testing.T) {
	err := initMongo()
	if err != nil {
		log.Fatal(err)
	}
	GUserService.CreateUser("test", "123")
	user, err := GUserService.GetUserByUsername("test")
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
	err = initPub()
	if err != nil {
		log.Fatal(err)
	}
	GUserService.CreateUser("test", "123")
	user, err := GUserService.GetUserByUsername("test")
	if err != nil {
		log.Fatal(err)
	}
	adPlan := &AdPlan{
		UserId:    user.Id,
		Name:      "test ad plan",
		StartTime: 100000,
		EndTime:   200000,
	}
	id, _ := GAdPlanService.CreateAdPlan(adPlan)
	plans, _ := GAdPlanService.GetAdPlans(user.Id)
	t.Log("got " + strconv.Itoa(len(plans)) + " plans")
	for _, plan := range plans {
		if plan.Id == id {
			t.Log("plan id = " + id + " got")
			break
		}
	}
	GAdPlanService.DeleteAdPlan(id)
	t.Log("plan id = " + id + " deleted")
	plans, _ = GAdPlanService.GetAdPlans(user.Id)
	t.Log("got " + strconv.Itoa(len(plans)) + " plans")
}

func TestAdUnitService(t *testing.T) {
	err := initMongo()
	if err != nil {
		log.Fatal(err)
	}
	err = initPub()
	if err != nil {
		log.Fatal(err)
	}
	GUserService.CreateUser("test", "123")

	user, err := GUserService.GetUserByUsername("test")
	adPlan := &AdPlan{
		UserId:    user.Id,
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
		UserId:   user.Id,
		Name:     "test ad innovation",
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
		AdInnovationId: inno.Id,
		AdUintId:       adUnit.Id,
	}
	GAdUnitService.CreateUnitInnovation(&iu)
}
