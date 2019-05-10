package main

import (
	"ads/pkg/ad_delivery/controller"
	"ads/pkg/ad_delivery/service"
	"ads/pkg/ad_search"
	"ads/pkg/ad_search/index"
	"ads/pkg/common"
	"cloud.google.com/go/firestore"
	"cloud.google.com/go/pubsub"
	"context"
	"github.com/go-redis/redis"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	// init configuration
	pwd, _ := os.Getwd()
	conf, err := common.InitConfig(pwd + "/pkg/config.json")
	if err != nil {
		log.Fatal(err)
	}
	common.GConfig = &conf
	// init database
	if err := InitDB(); err != nil {
		log.Fatal(err)
	}
	// init redis
	if err := InitRedis(); err != nil {
		log.Fatal(err)
	}
	// init message queue
	if err := InitPub(); err != nil {
		log.Fatal(err)
	}
	// init http server
	if err := InitHttpServer(); err != nil {
		log.Fatal(err)
	}
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

func InitDB() error {
	// Sets your Google Cloud Platform project ID.
	projectID := common.GConfig.ProjectId
	ctx := context.Background()
	// Get a Firestore client.
	pwd, _ := os.Getwd()
	client, err := firestore.NewClient(ctx, projectID, option.WithCredentialsFile(pwd+"/pkg/auth.json"))
	if err != nil {
		return err
	}
	service.GUserService = &service.UserService{
		Collection: client.Collection("ad_user"),
	}
	service.GAdPlanService = &service.AdPlanService{
		Collection: client.Collection("ad_plan"),
	}
	service.GAdUnitService = &service.AdUnitService{
		AdUintCollection:           client.Collection("ad_unit"),
		AdUintDistrictCollection:   client.Collection("ad_unit_district"),
		AdUintKeywordCollection:    client.Collection("ad_unit_keyword"),
		AdUintInterestCollection:   client.Collection("ad_unit_interest"),
		AdUnitInnovationCollection: client.Collection("ad_unit_innovation"),
	}
	service.GAdInnovationService = &service.AdInnovationService{
		Collection: client.Collection("ad_innovation"),
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

func InitPub() error {
	pwd, _ := os.Getwd()
	projectID := common.GConfig.ProjectId
	cli, err := pubsub.NewClient(context.Background(), projectID, option.WithCredentialsFile(pwd+"/pkg/auth.json"))
	if err != nil {
		return err
	}
	service.PubClient = cli
	return nil
}
