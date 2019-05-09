package service

import (
	"ads/pkg/ad_search/index"
	"ads/pkg/ad_search/protocol"
)

func FetchAds(request protocol.SearchRequest) protocol.SearchResponse {
	adSlots := request.RequestInfo.AdSlots
	keywordFeature := request.FeatureInfo.KeyWordFeature
	interestFeature := request.FeatureInfo.InterestFeature
	districtFeature := request.FeatureInfo.DistrictFeature
	isAnd := request.FeatureInfo.IsAnd

	resp := protocol.SearchResponse{Slot2Ads: make(map[string][]protocol.Innovation)}

	for _, slot := range adSlots {
		// 1 level filtering
		adUnitIndexes := index.AdUnitIndexMatch(slot.PositionType)
		// 2 level filtering
		if isAnd {
			filterKeyword(adUnitIndexes, keywordFeature)
			filterDistrict(adUnitIndexes, districtFeature)
			filterInterest(adUnitIndexes, interestFeature)
		} else {
			//Todo: or filtering implementation
		}
		adInnovationIndexes := index.SelectInnovationFromUnitIndexes(adUnitIndexes)
		// 3 level filtering
		filterInnovationBySlot(adInnovationIndexes, slot.Width, slot.Height, slot.Types)
		innos := buildInnovation(adInnovationIndexes)
		resp.Slot2Ads[slot.Code] = innos
	}
	return resp
}

func filterKeyword(adUnitIndexes map[string]*index.AdUnitIndex, feature protocol.KeyWordFeature) {
	for id := range adUnitIndexes {
		if !index.MarchInvertedIndex(id, index.UnitKeywordIndexById, feature.Keywords...) {
			delete(adUnitIndexes, id)
		}
	}
}
func filterDistrict(adUnitIndexes map[string]*index.AdUnitIndex, feature protocol.DistrictFeature) {
	districts := make([]string, 0)
	for _, dis := range feature.Districts {
		districts = append(districts, dis.State+"-"+dis.City)
	}
	for id := range adUnitIndexes {
		if !index.MarchInvertedIndex(id, index.UnitDistrictIndexById, districts...) {
			delete(adUnitIndexes, id)
		}
	}
}

func filterInterest(adUnitIndexes map[string]*index.AdUnitIndex, feature protocol.InterestFeature) {
	for id := range adUnitIndexes {
		if !index.MarchInvertedIndex(id, index.UnitInterestIndexById, feature.Interests...) {
			delete(adUnitIndexes, id)
		}
	}
}

func filterInnovationBySlot(innoIndexes map[string]*index.AdInnovationIndex, width int, height int, types []int) {
	for id, index := range innoIndexes {
		if index.Width != width {
			delete(innoIndexes, id)
		} else if index.Height != height {
			delete(innoIndexes, id)
		} else {
			res := false
			for _, tp := range types {
				if tp == index.Type {
					res = true
				}
			}
			if !res {
				delete(innoIndexes, id)
			}
		}
	}
}

func buildInnovation(innoIndexes map[string]*index.AdInnovationIndex) []protocol.Innovation {
	res := make([]protocol.Innovation, 0)
	for _, innoIdex := range innoIndexes {
		inno := protocol.Convert(*innoIdex)
		res = append(res, inno)
	}
	return res
}
