package main

import (
	"ads/pkg/ad_delivery/controller"
	"ads/pkg/ad_delivery/service"
	"ads/pkg/ad_search"
	"ads/pkg/ad_search/index"
	"ads/pkg/common"
	"github.com/go-redis/redis"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	// init configuration
	pwd, _ := os.Getwd()
	conf, err := common.InitConfig(pwd + "/pkg/config.json")
	if err != nil {
		log.Fatal(err)
	}
	common.GConfig = &conf
	// init MongoDB
	if err := InitMongo(); err != nil {
		log.Fatal(err)
	}
	// init redis
	if err := InitRedis(); err != nil {
		log.Fatal(err)
	}
	// init http server
	if err := InitHttpServer(); err != nil {
		log.Fatal(err)
	}
}

func InitMongo() error {
	session, err := mgo.DialWithTimeout(common.GConfig.MongoDBUri, time.Duration(common.GConfig.MongoDBTimeout)*time.Millisecond)
	if err != nil {
		return err
	}
	// init services
	service.GUserService = &service.UserService{
		Collection: session.DB(common.GConfig.DBName).C("ad_user"),
	}
	service.GAdPlanService = &service.AdPlanService{
		Collection: session.DB(common.GConfig.DBName).C("ad_plan"),
	}
	service.GAdUnitService = &service.AdUnitService{
		AdUintCollection:           session.DB(common.GConfig.DBName).C("ad_unit"),
		AdUintDistrictCollection:   session.DB(common.GConfig.DBName).C("ad_unit_district"),
		AdUintKeywordCollection:    session.DB(common.GConfig.DBName).C("ad_unit_keyword"),
		AdUintInterestCollection:   session.DB(common.GConfig.DBName).C("ad_unit_interest"),
		AdUnitInnovationCollection: session.DB(common.GConfig.DBName).C("ad_unit_innovation"),
	}
	service.GAdInnovationService = &service.AdInnovationService{
		Collection: session.DB(common.GConfig.DBName).C("ad_innovation"),
	}
	return nil
}

func InitRedis() error {
	client := redis.NewClient(&redis.Options{
		Addr:     common.GConfig.RedisUri,
		Password: "",
		DB:       0,
	})
	_, err := client.Ping().Result()
	index.GRedisClient = client
	return err
}

func InitHttpServer() error {
	mux := http.NewServeMux()
	controller.InitUserController(mux)
	controller.InitAdPlanController(mux)
	controller.InitAdUnitController(mux)
	controller.InitAdInnovationController(mux)
	ad_search.InitAdSearchController(mux)
	println("http server is starting on port: " + strconv.Itoa(common.GConfig.HttpPort))
	err := http.ListenAndServe(":"+strconv.Itoa(common.GConfig.HttpPort), mux)
	return err
}
