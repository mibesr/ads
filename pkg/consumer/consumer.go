package main

import (
	"ads/pkg/ad_delivery/service"
	"ads/pkg/ad_search/index"
	"ads/pkg/common"
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"google.golang.org/api/option"
	"log"
	"os"
)

// the consumer for index updating

var SubClient *pubsub.Client

func main() {
	pwd, _ := os.Getwd()
	conf, err := common.InitConfig(pwd + "/pkg/config.json")
	if err != nil {
		log.Fatal(err)
	}
	common.GConfig = &conf
	// init redis
	if err := InitRedis(); err != nil {
		log.Fatal(err)
	}
	projectID := common.GConfig.ProjectId

	cli, err := pubsub.NewClient(context.Background(), projectID, option.WithCredentialsFile(pwd+"/pkg/auth.json"))
	SubClient = cli

	go ConsumeSaveAdPlanIndex()
	go ConsumeDeleteAdPlanIndex()
	go ConsumeSaveAdUnitIndex()
	go ConsumeSaveKeywordIndex()
	go ConsumeSaveInterestIndex()
	go ConsumeSaveDistrictIndex()
	go ConsumeSaveUnitInnovationIndex()
	go ConsumeSaveAdInnovationIndex()
	select {}
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

func ConsumeSaveAdPlanIndex() {
	sub := SubClient.Subscription("ad_plan_save")
	err := sub.Receive(context.Background(), func(ctx context.Context, msg *pubsub.Message) {
		msg.Ack()
		adPlan := &service.AdPlan{}
		json.Unmarshal(msg.Data, adPlan)
		idx := index.AdPlanIndex{
			Id:        adPlan.Id,
			UserId:    adPlan.UserId,
			StartDate: adPlan.StartTime,
			EndDate:   adPlan.EndTime,
		}
		idx.Save()
		fmt.Printf("Got message: %q\n", string(msg.Data))
	})
	if err != nil {
		log.Fatal(err)
	}
}

func ConsumeDeleteAdPlanIndex() {
	sub := SubClient.Subscription("ad_plan_delete")
	err := sub.Receive(context.Background(), func(ctx context.Context, msg *pubsub.Message) {
		msg.Ack()
		idx := index.AdPlanIndex{
			Id: string(msg.Data),
		}
		idx.Delete()
		fmt.Printf("Got message: %q\n", string(msg.Data))
	})
	if err != nil {
		log.Fatal(err)
	}
}

func ConsumeSaveAdUnitIndex() {
	sub := SubClient.Subscription("ad_unit_save")
	err := sub.Receive(context.Background(), func(ctx context.Context, msg *pubsub.Message) {
		msg.Ack()
		adUnit := &service.AdUint{}
		json.Unmarshal(msg.Data, adUnit)
		planIdx := index.AdPlanIndex{
			Id: adUnit.PlanId,
		}
		planIdx.Get()
		idx := index.AdUnitIndex{
			Id:           adUnit.Id,
			PlanId:       adUnit.PlanId,
			PositionType: adUnit.PositionType,
			AdPlanObj:    planIdx,
		}
		idx.Save()
		fmt.Printf("Got message: %q\n", string(msg.Data))
	})
	if err != nil {
		log.Fatal(err)
	}
}

func ConsumeSaveKeywordIndex() {
	sub := SubClient.Subscription("keyword_save")
	err := sub.Receive(context.Background(), func(ctx context.Context, msg *pubsub.Message) {
		msg.Ack()
		keyword := &service.AdUnitKeyword{}
		json.Unmarshal(msg.Data, keyword)
		index.SaveInvertedIndexById(keyword.UnitId, index.UnitKeywordIndexById, keyword.Keyword)
		fmt.Printf("Got message: %q\n", string(msg.Data))
	})
	if err != nil {
		log.Fatal(err)
	}
}

func ConsumeSaveInterestIndex() {
	sub := SubClient.Subscription("interest_save")
	err := sub.Receive(context.Background(), func(ctx context.Context, msg *pubsub.Message) {
		msg.Ack()
		interest := &service.AdUnitInterest{}
		json.Unmarshal(msg.Data, interest)
		index.SaveInvertedIndexById(interest.UnitId, index.UnitInterestIndexById, interest.Tag)
		fmt.Printf("Got message: %q\n", string(msg.Data))
	})
	if err != nil {
		log.Fatal(err)
	}
}

func ConsumeSaveDistrictIndex() {
	sub := SubClient.Subscription("district_save")
	err := sub.Receive(context.Background(), func(ctx context.Context, msg *pubsub.Message) {
		msg.Ack()
		district := &service.AdUnitDistrict{}
		json.Unmarshal(msg.Data, district)
		index.SaveInvertedIndexById(district.UnitId, index.UnitDistrictIndexById, district.State+"-"+district.City)
		fmt.Printf("Got message: %q\n", string(msg.Data))
	})
	if err != nil {
		log.Fatal(err)
	}
}

func ConsumeSaveUnitInnovationIndex() {
	sub := SubClient.Subscription("unit_innovation_save")
	err := sub.Receive(context.Background(), func(ctx context.Context, msg *pubsub.Message) {
		msg.Ack()
		iu := &service.InnovationUnit{}
		json.Unmarshal(msg.Data, iu)
		index.SaveByInnoId(iu.AdInnovationId, iu.AdUintId)
		fmt.Printf("Got message: %q\n", string(msg.Data))
	})
	if err != nil {
		log.Fatal(err)
	}
}

func ConsumeSaveAdInnovationIndex() {
	sub := SubClient.Subscription("ad_innovation_save")
	err := sub.Receive(context.Background(), func(ctx context.Context, msg *pubsub.Message) {
		msg.Ack()
		inno := &service.AdInnovation{}
		json.Unmarshal(msg.Data, inno)
		idx := index.AdInnovationIndex{
			Id:           inno.Id,
			Name:         inno.Name,
			Type:         inno.Type,
			MaterialType: inno.Material,
			Height:       inno.Height,
			Width:        inno.Width,
			Url:          inno.Url,
		}
		idx.Save()
		fmt.Printf("Got message: %q\n", string(msg.Data))
	})
	if err != nil {
		log.Fatal(err)
	}
}
