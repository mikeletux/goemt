package busemtmad

import (
	//"encoding/json"
	//"fmt"
	"github.com/mikeletux/goemt"
	//"strings"
)

const (
	endpointStopsAroundStop = "/transport/busemtmad/stops/arroundstop/<stopId>/<radius>/"
)

/*
BusStops struct holds all information regarding the bus stops sourrounding one, given a radius
*/
type BusStops struct {
	Geometry struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	} `json:"geometry"`

	StopID        string `json:"stopId"`
	MetersToPoint int64  `json:"metersToPoint"`
	StopName      string `json:"stopName"`
	Lines         struct {
		NameA            string `json:"nameA"`
		NameB            string `json:"nameB"`
		MetersFromHeader int64  `json:"metersFromHeader"`
		Label            string `json:"label"`
		To               string `json:"to"`
		Line             string `json:"line"`
	} `json:"lines"`
}

/*
Aux struct that will be used to get the info from the Rest API
*/
type busStopsAux struct {
	goemt.Common
	BusStops
}

/*
GetStopsAroundStop returns the stops that sourround a BUS stop, given an specific distance (radius)
Parameters:
	api -> Struct that implements the IAPI interface (i.e APIClient)
	stopID -> Stop number.
	radius -> Radius in meters
*/
/*func GetStopsAroundStop(api goemt.IAPI, stopID string, radius int64) (busStops BusStops, err error) {
	if api == nil {

	}
}*/
