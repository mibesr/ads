package controller

import (
	"ads/pkg/ad_delivery/service"
	"ads/pkg/common"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type CreateUserReq struct {
	Username  string `json:"username"`
	Password1 string `json:"password1"`
	Password2 string `json:"password2"`
}

func InitUserController(mux *http.ServeMux) {
	mux.Handle("/ad_user/create", http.HandlerFunc(CreateUser))
}

// post
func CreateUser(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		data, err := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
		r := &CreateUserReq{}
		err = json.Unmarshal(data, r)
		id, err := service.GUserService.CreateUser(r.Username, r.Password1)
		if err != nil {
			common.Return(400, []byte(err.Error()), w)
		} else {
			common.Return(200, []byte(common.BuildDefaultResponse(id)), w)
		}
	} else {
		common.Return(405, []byte(common.HttpPostOnly.Error()), w)
	}
}
