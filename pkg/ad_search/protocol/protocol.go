package protocol

import "ads/pkg/ad_search/index"

// requests and responses of searching system

// The request object of media
type SearchRequest struct {
	MediaId     string      `json:"media_id"`
	RequestInfo RequestInfo `json:"request_info"`
	FeatureInfo FeatureInfo `json:"feature_info"`
}

type RequestInfo struct {
	RequestId string   `json:"request_id"`
	AdSlots   []AdSlot `json:"ad_slots"`
	App       App      `json:"app"`
	Geo       Geo      `json:"geo"`
	Device    Device   `json:"device"`
}
type FeatureInfo struct {
	KeyWordFeature  KeyWordFeature  `json:"keyword_feature"`
	InterestFeature InterestFeature `json:"interest_feature"`
	DistrictFeature DistrictFeature `json:"district_feature"`
	IsAnd           bool            `json:"is_and"` // and = true, or = false
}

/******************* Media *************************/

type AdSlot struct {
	Code         string  `json:"code"`
	PositionType int     `json:"position_type"`
	Width        int     `json:"width"`
	Height       int     `json:"height"`
	Types        []int   `json:"types"`
	MinCpm       float32 `json:"min_cpm"` // the lowest cost
}

type App struct {
	Code         string `json:"code"`
	Name         string `json:"name"`
	PackageName  string `json:"package_name"`
	ActivityName string `json:"activity_name"`
}

type Geo struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	State     string  `json:"state"`
	City      string  `json:"city"`
}

type Device struct {
	Code        string `json:"code"`
	Mac         string `json:"mac"`
	Ip          string `json:"ip"`
	Model       string `json:"model"`
	DisplaySize string `json:"display_size"`
	ScreenSize  string `json:"screen_size"`
	SerialName  string `json:"serial_name"`
}

/******************* Feature *************************/

type KeyWordFeature struct {
	Keywords []string `json:"keywords"`
}

type InterestFeature struct {
	Interests []string `json:"interests"`
}

type DistrictFeature struct {
	Districts []district `json:"districts"`
}

type district struct {
	State string `json:"state"`
	City  string `json:"city"`
}

// The response object of media

type SearchResponse struct {
	Slot2Ads map[string][]Innovation `json:"slot_2_ads"` // key -> adSlotCode, value -> innovation set
}

type Innovation struct {
	AdId         string `json:"ad_id"`
	AdUrl        string `json:"ad_url"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	Type         int    `json:"type"`
	MaterialType int    `json:"material_type"`
}

func Convert(index index.AdInnovationIndex) Innovation {
	var inno Innovation
	inno.AdId = index.Id
	inno.AdUrl = index.Url
	inno.Width = index.Width
	inno.Height = index.Height
	inno.Type = index.Type
	inno.MaterialType = index.MaterialType
	return inno
}
