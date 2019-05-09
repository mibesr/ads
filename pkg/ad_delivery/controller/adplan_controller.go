package controller

import (
	"ads/pkg/common"
	"ads/pkg/ad_delivery/service"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func InitAdPlanController(mux *http.ServeMux) {
	mux.Handle("/ad_plan/create", http.HandlerFunc(CreateAdPlan))
	mux.Handle("/ad_plans", http.HandlerFunc(GetAdPlans))
	mux.Handle("/ad_plan/update", http.HandlerFunc(UpdateAdPlan))
	mux.Handle("/ad_plan/delete", http.HandlerFunc(DeleteAdPlan))
}

type AdPlanReq struct {
	UserId    string `json:"user_id"`
	Name      string `json:"name"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
}

type AdPlanResp struct {
	Id         string `json:"id"`
	UserId     string `json:"user_id"`
	Name       string `json:"name"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

func CreateAdPlan(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		data, err := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
		r := &AdPlanReq{}
		err = json.Unmarshal(data, r)
		if err != nil {
			common.Return(400, []byte(common.JsonFormatError.Error()), w)
			return
		}
		adPlan := &service.AdPlan{
			UserId:    r.UserId,
			Name:      r.Name,
			StartTime: r.StartTime,
			EndTime:   r.EndTime,
		}
		id, err := service.GAdPlanService.CreateAdPlan(adPlan)
		if err != nil {
			common.Return(400, []byte(err.Error()), w)
		} else {
			common.Return(200, []byte(common.BuildDefaultResponse(id)), w)
		}
	} else {
		common.Return(405, []byte(common.HttpPostOnly.Error()), w)
	}
}

func GetAdPlans(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		ids, ok := req.URL.Query()["user_id"]
		if !ok || len(ids[0]) < 1 {
			common.Return(400, []byte(common.ParamError.Error()), w)
			return
		}
		id := ids[0]
		plans, err := service.GAdPlanService.GetAdPlans(id)
		if err != nil {
			common.Return(400, []byte(err.Error()), w)
		} else {
			resp := make([]AdPlanResp, 0)
			for _, plan := range plans {
				adPlanResp := AdPlanResp{
					Id:         plan.MongoId.Hex(),
					UserId:     plan.UserId,
					Name:       plan.Name,
					StartTime:  common.FormatTime(plan.StartTime),
					EndTime:    common.FormatTime(plan.EndTime),
					CreateTime: common.FormatTime(plan.CreateTime),
					UpdateTime: common.FormatTime(plan.UpdateTime),
				}
				resp = append(resp, adPlanResp)
			}
			common.Return(200, []byte(common.BuildDefaultResponse(resp)), w)
		}
	} else {
		common.Return(405, []byte(common.HttpPostOnly.Error()), w)
	}
}

func UpdateAdPlan(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		data, err := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		r := &AdPlanReq{}
		err = json.Unmarshal(data, r)
		if err != nil {
			common.Return(400, []byte(common.JsonFormatError.Error()), w)
			return
		}
		adPlan := &service.AdPlan{
			UserId:    r.UserId,
			Name:      r.Name,
			StartTime: r.StartTime,
			EndTime:   r.EndTime,
		}
		id, err := service.GAdPlanService.UpdateAdPlan(adPlan)
		if err != nil {
			common.Return(400, []byte(err.Error()), w)
		} else {
			common.Return(200, []byte(common.BuildDefaultResponse(id)), w)
		}
	} else {
		common.Return(405, []byte(common.HttpPostOnly.Error()), w)
	}
}

func DeleteAdPlan(w http.ResponseWriter, req *http.Request) {
	if req.Method == "DELETE" {
		ids, ok := req.URL.Query()["plan_id"]
		if !ok || len(ids[0]) < 1 {
			common.Return(400, []byte(common.ParamError.Error()), w)
			return
		}
		id := ids[0]
		err := service.GAdPlanService.DeleteAdPlan(id)
		if err != nil {
			common.Return(400, []byte(err.Error()), w)
		} else {
			common.Return(200, []byte(common.BuildDefaultResponse(id)), w)
		}
	} else {
		common.Return(405, []byte(common.HttpPostOnly.Error()), w)
	}
}
