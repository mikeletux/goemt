package busemtmad

import (
	"fmt"
	"github.com/mikeletux/goemt"
	"strings"
)

/*
StopsAroundPlaces struct where the list of stops and lines nearby street or places will be stored
*/
type StopsAroundPlaces struct {
	StopID        int    `json:"stopId"`
	StopName      string `json:"stopName"`
	Address       string `json:"address"`
	MetersToPoint int    `json:"metersToPoint"`
	Geometry      struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	} `json:"geometry"`

	Lines []struct {
		Line             string `json:"line"`
		Label            string `json:"label"`
		NameA            string `json:"nameA"`
		NameB            string `json:"nameB"`
		MetersFromHeader int    `json:"metersFromHeader"`
		To               string `json:"to"`
	} `json:"lines"`
}

type auxStopsAroundPlaces struct {
	goemt.Common
	Data []StopsAroundPlaces `json:"data"`
}

/*
GetStopsAroundPlaces func returns the list of stops and lines nearby street or places
Parameters:
	api -> Struct that implements the IAPI interface (i.e APIClient)
	namePlace -> partial name of place or street (use spaces between words, function takes care of adding %20)
	number -> number of street (0 if not need or for getting the first number of street)
	radius -> meters arround de place
Returns:
	stopsAroundPlaces -> slice with all requested data
	err -> if there's any error, err will be set. nil otherwise.
*/
func GetStopsAroundPlaces(api goemt.IAPI, namePlace string, number, radius int) (stopsAroundPlaces []StopsAroundPlaces, err error) {
	if namePlace == "" || radius == 0 {
		return stopsAroundPlaces, fmt.Errorf("namePlace must contain some string and radius must be different from 0")
	}
	finalServiceEnpoint := strings.ReplaceAll(endpointStopsAroundPlaces, "<namePlace>", fmt.Sprintf("%v", strings.ReplaceAll(namePlace, " ", "%20")))
	finalServiceEnpoint = strings.ReplaceAll(finalServiceEnpoint, "<number_or_zero_street_number>", fmt.Sprintf("%v", number))
	finalServiceEnpoint = strings.ReplaceAll(finalServiceEnpoint, "<radius>", fmt.Sprintf("%v", radius))
	var data auxStopsAroundPlaces
	err = getInfoFromPlatform(api, finalServiceEnpoint, &data)
	if err != nil {
		return stopsAroundPlaces, err
	}
	return data.Data, nil
}
