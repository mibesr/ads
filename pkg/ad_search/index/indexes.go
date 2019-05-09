package index

import (
	"encoding/json"
	"github.com/go-redis/redis"
)

var GRedisClient *redis.Client

type Index interface {
	Get() Index

	Save()

	Delete()

	ToString() string
}

// ad plan index
type AdPlanIndex struct {
	Id        string
	UserId    string
	StartDate int64
	EndDate   int64
}

func (plan *AdPlanIndex) Get() Index {
	str, _ := GRedisClient.Get("AdPlanIndex-" + plan.Id).Result()
	err := json.Unmarshal([]byte(str), plan)
	if err != nil {
		panic(err)
	}
	return plan
}

func (plan *AdPlanIndex) Save() {
	GRedisClient.Set("AdPlanIndex-"+plan.Id, plan.ToString(), 0)
}

func (plan *AdPlanIndex) Delete() {
	GRedisClient.Del(plan.Id)
}

func (plan *AdPlanIndex) ToString() string {
	data, _ := json.Marshal(plan)
	return string(data)
}

type AdUnitIndex struct {
	Id           string
	PlanId       string
	PositionType int
	AdPlanObj    AdPlanIndex
}

func (unit *AdUnitIndex) Get() Index {
	str, _ := GRedisClient.Get("AdUnitIndex-" + unit.Id).Result()
	err := json.Unmarshal([]byte(str), unit)
	if err != nil {
		panic(err)
	}
	return unit
}

func (unit *AdUnitIndex) Save() {
	GRedisClient.Set("AdUnitIndex-"+unit.Id, unit.ToString(), 0)
}

func (unit *AdUnitIndex) Delete() {
	GRedisClient.Del(unit.Id)
}

func (unit *AdUnitIndex) ToString() string {
	data, _ := json.Marshal(unit)
	return string(data)
}

func AdUnitIndexMatch(positionType int) map[string]*AdUnitIndex {
	res := make(map[string]*AdUnitIndex)
	ids, _ := GRedisClient.Keys("AdUnitIndex-*").Result()
	for _, id := range ids {
		data, _ := GRedisClient.Get(id).Result()
		unitIndex := &AdUnitIndex{}
		json.Unmarshal([]byte(data), unitIndex)
		if unitIndex.PositionType == positionType {
			res[unitIndex.Id] = unitIndex
		}
	}
	return res
}

type AdInnovationIndex struct {
	Id           string
	Name         string
	Type         int
	MaterialType int
	Height       int
	Width        int
	Url          string
}

func (inno *AdInnovationIndex) Get() Index {
	str, _ := GRedisClient.Get("AdInnovationIndex-" + inno.Id).Result()
	err := json.Unmarshal([]byte(str), inno)
	if err != nil {
		panic(err)
	}
	return inno
}

func (inno *AdInnovationIndex) Save() {
	GRedisClient.Set("AdInnovationIndex-"+inno.Id, inno.ToString(), 0)
}

func (inno *AdInnovationIndex) Delete() {
	GRedisClient.Del(inno.Id)
}

func (inno *AdInnovationIndex) ToString() string {
	data, _ := json.Marshal(inno)
	return string(data)
}

// ************** ad unit - innovation *********************

func GetByInnoId(innoId string) map[string]bool {
	ids, _ := GRedisClient.SMembers("InnovationUnitIndex-InnovationId-" + innoId).Result()
	res := make(map[string]bool)
	for _, id := range ids {
		res[id] = true
	}
	return res
}

func GetByUnitId(unitId string) map[string]bool {
	ids, _ := GRedisClient.SMembers("InnovationUnitIndex-UnitId-" + unitId).Result()
	res := make(map[string]bool)
	for _, id := range ids {
		res[id] = true
	}
	return res
}

func SelectInnovationFromUnitIndexes(adUnitIndexes map[string]*AdUnitIndex) map[string]*AdInnovationIndex {
	res := make(map[string]*AdInnovationIndex)
	for unitIndexId := range adUnitIndexes {
		adIds := GetByUnitId(unitIndexId)
		for adId := range adIds {
			innoIndex := &AdInnovationIndex{Id: adId}
			innoIndex.Get()
			res[innoIndex.Id] = innoIndex
		}
	}
	return res
}

func SaveByInnoId(innoId string, unitIds ...string) {
	GRedisClient.SAdd("InnovationUnitIndex-InnovationId-"+innoId, unitIds)
	for _, unitId := range unitIds {
		GRedisClient.SAdd("InnovationUnitIndex-UnitId-"+unitId, innoId)
	}
}

func SaveByUnitId(unitId string, innoIds ...string) {
	GRedisClient.SAdd("InnovationUnitIndex-UnitId-"+unitId, innoIds)
	for _, innoId := range innoIds {
		GRedisClient.SAdd("InnovationUnitIndex-InnovationId-"+innoId, unitId)
	}
}

func DeleteByInnoId(innoId string, unitIds ...string) {
	GRedisClient.SRem("InnovationUnitIndex-InnovationId-"+innoId, unitIds)
	for _, unitId := range unitIds {
		GRedisClient.SRem("InnovationUnitIndex-UnitId-"+unitId, innoId)
	}
}

func DeleteByUnitId(unitId string, innoIds ...string) {
	GRedisClient.SRem("InnovationUnitIndex-UnitId-"+unitId, innoIds)
	for _, innoId := range innoIds {
		GRedisClient.SRem("InnovationUnitIndex-InnovationId-"+innoId, unitId)
	}
}

// *************************************************************************
// Inverted index is used for unit keyword, unit interest, unit district obj
// id = unit id
// for unit keyword, keyword = {keyword}
// for unit interest, keyword = {interest}
// for unit district, keyword = {state}-{city}

type InvertedIndexKeywordType string

const (
	UnitKeywordIndexByKey  InvertedIndexKeywordType = "UnitKeywordIndex-Keyword-"
	UnitInterestIndexByKey InvertedIndexKeywordType = "UnitInterestIndex-Keyword-"
	UnitDistrictIndexByKey InvertedIndexKeywordType = "UnitDistrictIndex-Keyword-"
)

func (t InvertedIndexKeywordType) toIdType() InvertedIndexIdType {
	switch t {
	case UnitKeywordIndexByKey:
		return UnitKeywordIndexById
	case UnitInterestIndexByKey:
		return UnitInterestIndexById
	case UnitDistrictIndexByKey:
		return UnitDistrictIndexById
	}
	return ""
}

func GetInvertedIndexByKey(keyword string, t InvertedIndexKeywordType) map[string]bool {
	ids, _ := GRedisClient.SMembers(string(t) + keyword).Result()
	res := make(map[string]bool)
	for _, id := range ids {
		res[id] = true
	}
	return res
}

func SaveInvertedIndexByKey(keyword string, t InvertedIndexKeywordType, ids ...string) {
	GRedisClient.SAdd(string(t)+keyword, ids)
	for _, id := range ids {
		GRedisClient.SAdd(string(t.toIdType())+id, keyword)
	}
}

func DeleteInvertedIndexByKey(keyword string, t InvertedIndexKeywordType, ids ...string) {
	GRedisClient.SRem(string(t)+keyword, ids)
	for _, id := range ids {
		GRedisClient.SRem(string(t.toIdType())+id, keyword)
	}
}

type InvertedIndexIdType string

const (
	UnitKeywordIndexById  InvertedIndexIdType = "UnitKeywordIndex-Id-"
	UnitInterestIndexById InvertedIndexIdType = "UnitInterestIndex-Id-"
	UnitDistrictIndexById InvertedIndexIdType = "UnitDistrictIndex-Id-"
)

func (t InvertedIndexIdType) toKeywordType() InvertedIndexKeywordType {
	switch t {
	case UnitKeywordIndexById:
		return UnitKeywordIndexByKey
	case UnitInterestIndexById:
		return UnitInterestIndexByKey
	case UnitDistrictIndexById:
		return UnitDistrictIndexByKey
	}
	return ""
}

func GetInvertedIndexById(id string, t InvertedIndexIdType) map[string]bool {
	ids, _ := GRedisClient.SMembers(string(t) + id).Result()
	res := make(map[string]bool)
	for _, id := range ids {
		res[id] = true
	}
	return res
}

func SaveInvertedIndexById(id string, t InvertedIndexIdType, keywords ...string) {
	GRedisClient.SAdd(string(t)+id, keywords)
	for _, keyword := range keywords {
		GRedisClient.SAdd(string(t.toKeywordType())+keyword, id)
	}
}

func DeleteInvertedIndexById(id string, t InvertedIndexIdType, keywords ...string) {
	GRedisClient.SRem(string(t)+id, keywords)
	for _, keyword := range keywords {
		GRedisClient.SRem(string(t.toKeywordType())+keyword, id)
	}
}

func MarchInvertedIndex(id string, t InvertedIndexIdType, keywords ...string) bool {
	kset := GetInvertedIndexById(id, t)
	for _, key := range keywords {
		if _, ok := kset[key]; ok == false {
			return false
		}
	}
	return true
}
