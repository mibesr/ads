package controller

import (
	"ads/pkg/ad_delivery/service"
	"ads/pkg/common"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func InitAdUnitController(mux *http.ServeMux) {
	mux.Handle("/ad_unit/create", http.HandlerFunc(CreateAdUnit))
	mux.Handle("/ad_unit/keyword", http.HandlerFunc(CreateAdUnitKeyword))
	mux.Handle("/ad_unit/district", http.HandlerFunc(CreateAdUnitDistrict))
	mux.Handle("/ad_unit/interest", http.HandlerFunc(CreateAdUnitInterest))
}

type AdUnitReq struct {
	PlanId       string `json:"plan_id"`
	Name         string `json:"name"`
	PositionType int    `json:"position_type"`
	Budget       int    `json:"budget"`
}

type KeywordReq struct {
	UnitId  string `json:"unit_id"`
	Keyword string `json:"keyword"`
}

type DistrictReq struct {
	UnitId string `json:"unit_id"`
	State  string `json:"state"`
	City   string `json:"city"`
}

type InterestReq struct {
	UnitId string `json:"unit_id"`
	Tag    string `json:"tag"`
}

func CreateAdUnit(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		data, err := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
		r := &AdUnitReq{}
		err = json.Unmarshal(data, r)
		if err != nil {
			common.Return(400, []byte(common.JsonFormatError.Error()), w)
			return
		}
		unit := &service.AdUint{
			PlanId:       r.PlanId,
			Name:         r.Name,
			PositionType: r.PositionType,
			Budget:       r.Budget,
		}
		id, err := service.GAdUnitService.CreateAdUnit(unit)
		if err != nil {
			common.Return(400, []byte(err.Error()), w)
		} else {
			common.Return(200, []byte(common.BuildDefaultResponse(id)), w)
		}
	} else {
		common.Return(405, []byte(common.HttpPostOnly.Error()), w)
	}
}

func CreateAdUnitKeyword(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		data, err := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
		r := &KeywordReq{}
		err = json.Unmarshal(data, r)
		if err != nil {
			common.Return(400, []byte(common.JsonFormatError.Error()), w)
			return
		}
		keyword := &service.AdUnitKeyword{
			UnitId:  r.UnitId,
			Keyword: r.Keyword,
		}
		id, err := service.GAdUnitService.CreateKeyword(keyword)
		if err != nil {
			common.Return(400, []byte(err.Error()), w)
		} else {
			common.Return(200, []byte(common.BuildDefaultResponse(id)), w)
		}
	} else {
		common.Return(405, []byte(common.HttpPostOnly.Error()), w)
	}
}

func CreateAdUnitDistrict(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		data, err := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
		r := &DistrictReq{}
		err = json.Unmarshal(data, r)
		if err != nil {
			common.Return(400, []byte(common.JsonFormatError.Error()), w)
			return
		}
		district := &service.AdUnitDistrict{
			UnitId: r.UnitId,
			State:  r.State,
			City:   r.City,
		}
		id, err := service.GAdUnitService.CreateDistrict(district)
		if err != nil {
			common.Return(400, []byte(err.Error()), w)
		} else {
			common.Return(200, []byte(common.BuildDefaultResponse(id)), w)
		}
	} else {
		common.Return(405, []byte(common.HttpPostOnly.Error()), w)
	}
}

func CreateAdUnitInterest(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		data, err := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
		r := &InterestReq{}
		err = json.Unmarshal(data, r)
		if err != nil {
			common.Return(400, []byte(common.JsonFormatError.Error()), w)
			return
		}
		it := &service.AdUnitInterest{
			UnitId: r.UnitId,
			Tag:    r.Tag,
		}
		id, err := service.GAdUnitService.CreateInterest(it)
		if err != nil {
			common.Return(400, []byte(err.Error()), w)
		} else {
			common.Return(200, []byte(common.BuildDefaultResponse(id)), w)
		}
	} else {
		common.Return(405, []byte(common.HttpPostOnly.Error()), w)
	}
}
