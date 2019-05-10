package service

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"fmt"
	"log"
)

var PubClient *pubsub.Client

func SaveAdPlanIndex(adPlan *AdPlan) {
	t := PubClient.Topic("ad_plan_save")
	data, _ := json.Marshal(&adPlan)
	result := t.Publish(context.Background(), &pubsub.Message{
		Data: data,
	})
	id, err := result.Get(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Published a message; msg ID: %v\n", id)
}

func DeleteAdPlanIndex(id string) {
	t := PubClient.Topic("ad_plan_delete")
	result := t.Publish(context.Background(), &pubsub.Message{
		Data: []byte(id),
	})
	id, err := result.Get(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Published a message; msg ID: %v\n", id)
}

func SaveAdUnitIndex(adUint *AdUint){
	t := PubClient.Topic("ad_unit_save")
	data, _ := json.Marshal(&adUint)
	result := t.Publish(context.Background(), &pubsub.Message{
		Data: data,
	})
	id, err := result.Get(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Published a message; msg ID: %v\n", id)
}

func SaveKeywordIndex(keyword *AdUnitKeyword){
	t := PubClient.Topic("keyword_save")
	data, _ := json.Marshal(&keyword)
	result := t.Publish(context.Background(), &pubsub.Message{
		Data: data,
	})
	id, err := result.Get(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Published a message; msg ID: %v\n", id)
}

func SaveInterestIndex(interest *AdUnitInterest){
	t := PubClient.Topic("interest_save")
	data, _ := json.Marshal(&interest)
	result := t.Publish(context.Background(), &pubsub.Message{
		Data: data,
	})
	id, err := result.Get(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Published a message; msg ID: %v\n", id)
}

func SaveDistrictIndex(district *AdUnitDistrict){
	t := PubClient.Topic("district_save")
	data, _ := json.Marshal(&district)
	result := t.Publish(context.Background(), &pubsub.Message{
		Data: data,
	})
	id, err := result.Get(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Published a message; msg ID: %v\n", id)
}

func SaveUnitInnovationIndex(iu *InnovationUnit){
	t := PubClient.Topic("unit_innovation_save")
	data, _ := json.Marshal(&iu)
	result := t.Publish(context.Background(), &pubsub.Message{
		Data: data,
	})
	id, err := result.Get(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Published a message; msg ID: %v\n", id)
}

func SaveAdInnovationIndex(inno *AdInnovation){
	t := PubClient.Topic("ad_innovation_save")
	data, _ := json.Marshal(&inno)
	result := t.Publish(context.Background(), &pubsub.Message{
		Data: data,
	})
	id, err := result.Get(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Published a message; msg ID: %v\n", id)
}