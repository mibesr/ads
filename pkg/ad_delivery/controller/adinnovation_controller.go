package controller

import (
	"ads/pkg/common"
	"ads/pkg/ad_delivery/service"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func InitAdInnovationController(mux *http.ServeMux) {
	mux.Handle("/ad_innovation/create", http.HandlerFunc(CreateInnovation))
}

type AdInnovationReq struct {
	UserId   string `json:"user_id"`
	Name     string `json:"name"`
	Type     int    `json:"type"`
	Material int    `json:"material"`
	Height   int  `json:"height"`
	Width    int  `json:"width"`
	Size     int    `json:"size"`
	Duration int    `json:"duration"`
	Url      string `json:"url"`
}

func CreateInnovation(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		data, err := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
		r := &AdInnovationReq{}
		err = json.Unmarshal(data, r)
		if err != nil {
			common.Return(400, []byte(common.JsonFormatError.Error()), w)
			return
		}
		inno := &service.AdInnovation{
			UserId:   r.UserId,
			Name:     r.Name,
			Type:     r.Type,
			Material: r.Material,
			Height:   r.Height,
			Width:    r.Width,
			Size:     r.Size,
			Duration: r.Duration,
			Url:      r.Url,
		}
		id, err := service.GAdInnovationService.CreateAdInnovation(inno)
		if err != nil {
			common.Return(400, []byte(err.Error()), w)
		} else {
			common.Return(200, []byte(common.BuildDefaultResponse(id)), w)
		}
	} else {
		common.Return(405, []byte(common.HttpPostOnly.Error()), w)
	}
}
