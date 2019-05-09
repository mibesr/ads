package ad_search

import (
	"ads/pkg/ad_search/protocol"
	"ads/pkg/ad_search/service"
	"ads/pkg/common"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func InitAdSearchController(mux *http.ServeMux) {
	mux.Handle("/ad_search/fetch", http.HandlerFunc(FetchAds))
}

func FetchAds(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		data, err := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
		r := protocol.SearchRequest{}
		err = json.Unmarshal(data, r)
		resp := service.FetchAds(r)
		common.Return(200, []byte(common.BuildDefaultResponse(resp)), w)
	} else {
		common.Return(405, []byte(common.HttpPostOnly.Error()), w)
	}
}
