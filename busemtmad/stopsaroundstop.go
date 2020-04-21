package busemtmad

import (
	//"encoding/json"
	"fmt"
	"github.com/mikeletux/goemt"
	"strings"
)

/*
BusStop struct holds all information regarding the bus stops sourrounding one, given a radius
*/
type BusStop struct {
	Geometry struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	} `json:"geometry"`

	StopID        int64  `json:"stopId"`
	MetersToPoint int64  `json:"metersToPoint"`
	StopName      string `json:"stopName"`
	Lines         []struct {
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
	Data []BusStop `json:"data"`
}

/*
GetStopsAroundStop returns the stops that sourround a BUS stop, given an specific distance (radius)
Parameters:
	api -> Struct that implements the IAPI interface (i.e APIClient)
	stopID -> Stop number.
	radius -> Radius in meters
Returns:
	busStops -> slice with BusStop structs with the queried data
	err -> if there's any error, err will be set. nil otherwise.
*/
func GetStopsAroundStop(api goemt.IAPI, stopID int, radius int) (busStops []BusStop, err error) {
	var data busStopsAux
	if stopID == 0 || radius == 0 {
		return busStops, fmt.Errorf("stopID and radius must be different from 0")
	}
	finalServiceEnpoint := strings.ReplaceAll(endpointStopsAroundStop, "<stopId>", fmt.Sprintf("%d", stopID))
	finalServiceEnpoint = strings.ReplaceAll(finalServiceEnpoint, "<radius>", fmt.Sprintf("%d", radius))
	err = getInfoFromPlatform(api, finalServiceEnpoint, &data)
	if err != nil {
		return busStops, err
	}
	return data.Data, nil
}

/*
GetStopsAroundGeographicalPoint returns the stops that sourround a given geographical point (longitude and latitude), given an specific distance (radius)
Parameters:
	api -> Struct that implements the IAPI interface (i.e APIClient)
	longitude and latitude -> Geographical point
	radius -> Radius in meters
Returns:
	busStops -> slice with BusStop structs with the queried data
	err -> if there's any error, err will be set. nil otherwise.
*/
func GetStopsAroundGeographicalPoint(api goemt.IAPI, longitude, latitude float64, radius int) (busStops []BusStop, err error) {
	var data busStopsAux
	if longitude == 0.0 || latitude == 0.0 || radius == 0 { //It is true that lat or long 0.0 exist, but since the API is for Madrid EMT bus system, it doesn't make sense
		return busStops, fmt.Errorf("longitude, latitude and radius must be different from 0")
	}
	finalServiceEnpoint := strings.ReplaceAll(endpointStopsAroundGeographicalPoint, "<longitude>", fmt.Sprintf("%v", longitude))
	finalServiceEnpoint = strings.ReplaceAll(finalServiceEnpoint, "<latitude>", fmt.Sprintf("%v", latitude))
	finalServiceEnpoint = strings.ReplaceAll(finalServiceEnpoint, "<radius>", fmt.Sprintf("%d", radius))
	err = getInfoFromPlatform(api, finalServiceEnpoint, &data)
	if err != nil {
		return busStops, err
	}
	return data.Data, nil
}
