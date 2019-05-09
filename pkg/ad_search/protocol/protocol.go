package protocol

import "ads/pkg/ad_search/index"

// requests and responses of searching system

// The request object of media
type SearchRequest struct {
	MediaId string
	RequestInfo
	FeatureInfo
}

type RequestInfo struct {
	RequestId string
	AdSlots   []AdSlot
	App
	Geo
	Device
}
type FeatureInfo struct {
	KeyWordFeature
	InterestFeature
	DistrictFeature
	IsAnd bool // and = true, or = false
}

/******************* Media *************************/

type AdSlot struct {
	Code         string
	PositionType int
	Width        int
	Height       int
	Types        []int
	MinCpm       float32 // the lowest cost
}

type App struct {
	Code         string
	Name         string
	PackageName  string
	ActivityName string
}

type Geo struct {
	Latitude  float32
	Longitude float32
	State     string
	City      string
}

type Device struct {
	Code        string
	Mac         string
	Ip          string
	Model       string
	DisplaySize string
	ScreenSize  string
	SerialName  string
}

/******************* Feature *************************/

type KeyWordFeature struct {
	Keywords []string
}

type InterestFeature struct {
	Interests []string
}

type DistrictFeature struct {
	Districts []district
}

type district struct {
	State string
	City  string
}

// The response object of media

type SearchResponse struct {
	Slot2Ads map[string][]Innovation // key -> adSlotCode, value -> innovation set
}

type Innovation struct {
	AdId         string
	AdUrl        string
	Width        int
	Height       int
	Type         int
	MaterialType int
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
